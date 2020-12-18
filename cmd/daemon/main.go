package main

import (
	"io"
	"net/http"
	"net/rpc"

	"nimona.io/internal/version"
	"nimona.io/pkg/config"
	"nimona.io/pkg/context"
	"nimona.io/pkg/daemon"
	"nimona.io/pkg/log"
)

func main() {
	ctx := context.Background()
	logger := log.FromContext(ctx).With(
		log.String("build.version", version.Version),
		log.String("build.commit", version.Commit),
		log.String("build.timestamp", version.Date),
	)
	ncfg, err := config.New()
	if err != nil {
		logger.Fatal("error parsing config", log.Error(err))
	}

	dmn := daemon.NewRPCServer(ctx, ncfg, logger)

	if err := rpc.Register(dmn); err != nil {
		logger.Fatal("error parsing config", log.Error(err))
	}

	rpc.HandleHTTP()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "ok")
	})

	http.ListenAndServe(":9000", nil)
}
