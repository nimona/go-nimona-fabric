package daemon

import (
	"database/sql"

	"nimona.io/internal/net"
	"nimona.io/pkg/blob"
	"nimona.io/pkg/config"
	"nimona.io/pkg/context"
	"nimona.io/pkg/hyperspace/resolver"
	"nimona.io/pkg/localpeer"
	"nimona.io/pkg/log"
	"nimona.io/pkg/network"
	"nimona.io/pkg/objectmanager"
	"nimona.io/pkg/objectstore"
	"nimona.io/pkg/peer"
	"nimona.io/pkg/sqlobjectstore"
)

type localService struct {
	local         localpeer.LocalPeer
	objectmanager objectmanager.ObjectManager
	objectstore   objectstore.Store
	blobmanager   blob.Manager
	resolver      resolver.Resolver
	listener      net.Listener
	config        *config.Config
}

func newLocalService(
	ctx context.Context,
	cfg *config.Config,
	logger log.Logger,
) *localService {
	lcsrv := &localService{}
	lcsrv.config = cfg
	// construct local peer
	local := localpeer.New()
	// attach peer private key from config
	local.PutPrimaryPeerKey(cfg.Peer.PrivateKey)
	local.PutContentTypes(
		// new(File).Type(),
		new(blob.Blob).Type(),
		new(blob.Chunk).Type(),
	)
	lcsrv.local = local

	// construct new network
	net := network.New(
		ctx,
		network.WithLocalPeer(local),
	)

	if cfg.Peer.BindAddress != "" {
		// start listening
		lis, err := net.Listen(
			ctx,
			cfg.Peer.BindAddress,
			network.ListenOnLocalIPs,
			network.ListenOnPrivateIPs,
		)
		if err != nil {
			logger.Fatal("error while listening", log.Error(err))
		}
		lcsrv.listener = lis
	}

	// convert shorthands into peers
	bootstrapPeers := []*peer.ConnectionInfo{}
	for _, s := range cfg.Peer.Bootstraps {
		bootstrapPeer, err := s.ConnectionInfo()
		if err != nil {
			logger.Fatal("error parsing bootstrap peer", log.Error(err))
		}
		bootstrapPeers = append(bootstrapPeers, bootstrapPeer)
	}

	// add bootstrap peers as relays
	local.PutRelays(bootstrapPeers...)

	// construct new resolver
	res := resolver.New(
		ctx,
		net,
		resolver.WithBoostrapPeers(bootstrapPeers...),
	)
	lcsrv.resolver = res

	logger = logger.With(
		log.String("peer.publicKey", local.GetPrimaryPeerKey().PublicKey().String()),
		log.Strings("peer.addresses", local.GetAddresses()),
	)

	// construct object store
	db, err := sql.Open("sqlite3", "file_transfer.db")
	if err != nil {
		logger.Fatal("error opening sql file", log.Error(err))
	}

	str, err := sqlobjectstore.New(db)
	if err != nil {
		logger.Fatal("error starting sql store", log.Error(err))
	}
	lcsrv.objectstore = str

	// construct object manager
	man := objectmanager.New(
		ctx,
		net,
		res,
		str,
	)
	lcsrv.objectmanager = man

	// construct blob manager
	bm := blob.NewManager(
		ctx,
		blob.WithObjectManager(man),
		blob.WithResolver(res),
	)
	lcsrv.blobmanager = bm

	return lcsrv
}
