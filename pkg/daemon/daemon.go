package daemon

import (
	"fmt"
	"os"
	"path"
	"strings"

	ssql "database/sql"

	_ "github.com/mattn/go-sqlite3"

	"nimona.io/internal/api"
	"nimona.io/internal/store/sql"
	"nimona.io/internal/version"
	"nimona.io/pkg/context"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/daemon/config"
	"nimona.io/pkg/discovery"
	"nimona.io/pkg/discovery/hyperspace"
	"nimona.io/pkg/exchange"
	"nimona.io/pkg/log"
	"nimona.io/pkg/middleware/handshake"
	"nimona.io/pkg/net"
	"nimona.io/pkg/orchestrator"
	"nimona.io/pkg/peer"
)

func New() *Daemon {
	return &Daemon{}
}

type Daemon struct {
}

func (d *Daemon) Start() {
	nodeAlias := os.Getenv("NIMONA_ALIAS")
	if nodeAlias != "" {
		log.DefaultLogger = log.DefaultLogger.With(
			log.String("$alias", nodeAlias),
		)
	}

	ctx := context.New(
		context.WithCorrelationID("nimona"),
	)

	logger := log.FromContext(ctx).With(
		log.String("build.version", version.Version),
		log.String("build.commit", version.Commit),
		log.String("build.timestamp", version.Date),
	)

	// load config
	logger.Info("loading config file")
	config := config.New()
	if err := config.Load(); err != nil {
		logger.Fatal("could not load config file", log.Error(err))
	}

	// create peer key pair if it does not exist
	if config.Peer.PeerKey == "" {
		logger.Info("creating new peer key pair")
		peerKey, err := crypto.GenerateEd25519PrivateKey()
		if err != nil {
			logger.Fatal("could not generate peer key", log.Error(err))
		}
		config.Peer.PeerKey = peerKey
	}

	// create identity key pair if it does not exist
	// TODO this is temporary
	if config.Peer.IdentityKey == "" {
		logger.Info("creating new identity key pair")
		identityKey, err := crypto.GenerateEd25519PrivateKey()
		if err != nil {
			logger.Fatal("could not generate identity key", log.Error(err))
		}
		config.Peer.IdentityKey = identityKey
	}

	// make sure relays are valid
	for i, addr := range config.Peer.RelayAddresses {
		if !strings.HasPrefix(addr, "relay:") {
			config.Peer.RelayAddresses[i] = "relay:" + addr
		}
	}

	// update config
	if err := config.Update(); err != nil {
		logger.Fatal("could not update config", log.Error(err))
	}

	logger.Info("loaded config", log.Any("config", config))

	// start daemon

	// construct discoverer
	discoverer := discovery.NewDiscoverer()

	// construct local info
	localInfo, err := peer.NewLocalPeer(
		config.Peer.AnnounceHostname,
		config.Peer.PeerKey,
	)
	if err != nil {
		logger.Fatal("could not create local info", log.Error(err))
	}

	// add content types
	localInfo.AddContentTypes(config.Peer.ContentTypes...)

	// add identity key to local info
	if config.Peer.IdentityKey != "" {
		if err := localInfo.AddIdentityKey(config.Peer.IdentityKey); err != nil {
			logger.Fatal("could not register identity key", log.Error(err))
		}
	}

	// add relay addresses to local info
	localInfo.AddAddress("relay", config.Peer.RelayAddresses)

	network, err := net.New(discoverer, localInfo)
	if err != nil {
		logger.Fatal("could not create network", log.Error(err))
	}

	// construct tcp transport
	tcpTransport := net.NewTCPTransport(
		localInfo,
		fmt.Sprintf("0.0.0.0:%d", config.Peer.TCPPort),
	)

	// add transports to network
	network.AddTransport("tcps", tcpTransport)

	// construct handshake
	handshakeMiddleware := handshake.New(
		localInfo,
		discoverer,
	)

	// add middleware to network
	network.AddMiddleware(handshakeMiddleware.Handle())

	// construct graph store
	db, err := ssql.Open("sqlite3", path.Join(config.Path, "sqlite3.db"))
	if err != nil {
		logger.Fatal("could not open sql file", log.Error(err))
	}

	store, err := sql.New(db)
	if err != nil {
		logger.Fatal("could not start sql store", log.Error(err))
	}

	// construct exchange
	exchange, err := exchange.New(
		ctx,
		config.Peer.PeerKey,
		network,
		store,
		discoverer,
		localInfo,
	)
	if err != nil {
		logger.Fatal("could not construct exchange", log.Error(err))
	}

	// construct hyperspace discoverer
	hyperspace, err := hyperspace.NewDiscoverer(
		ctx,
		exchange,
		localInfo,
		config.Peer.BootstrapAddresses,
	)
	if err != nil {
		logger.Fatal("could not construct hyperspace", log.Error(err))
	}

	// construct orchestrator
	orchestrator, err := orchestrator.New(
		store,
		exchange,
		nil,
		localInfo,
	)
	if err != nil {
		logger.Fatal("could not construct orchestrator", log.Error(err))
	}

	// add hyperspace provider
	if err := discoverer.AddProvider(hyperspace); err != nil {
		logger.Fatal("could not add hyperspace provider", log.Error(err))
	}

	// print some info
	nlogger := logger.With(
		log.Strings("addresses", localInfo.GetAddresses()),
		log.String("peer", config.Peer.PeerKey.PublicKey().String()),
	)

	ik := config.Peer.IdentityKey
	if ik != "" {
		nlogger = nlogger.With(
			log.String("identity", ik.PublicKey().String()),
		)
	}

	nlogger.Info("starting daemon")

	// construct api server
	apiServer := api.New(
		config,
		config.Peer.PeerKey,
		network,
		discoverer,
		exchange,
		localInfo,
		store,
		orchestrator,
		version.Version,
		version.Commit,
		version.Date,
		config.API.Token,
	)

	apiAddress := fmt.Sprintf("%s:%d", config.API.Host, config.API.Port)
	logger.Info(
		"starting http server",
		log.String("address", apiAddress),
	)
	apiServer.Serve(apiAddress) // nolint: errcheck
}
