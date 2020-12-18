package daemon

import (
	"fmt"

	"nimona.io/pkg/config"
	"nimona.io/pkg/context"
	"nimona.io/pkg/log"
	"nimona.io/pkg/object"
)

type (
	Server interface {
		PutObject(req PutObjectRequest, res *PutObjectResponse) error
	}
	PutObjectRequest struct {
		// Ctx context.Context
		Obj *object.Object
	}
	PutObjectResponse struct {
		Obj *object.Object
	}
	RPCServer struct {
		localService *localService
	}
)

func NewRPCServer(
	ctx context.Context,
	cfg *config.Config,
	logger log.Logger,
) Server {
	return &RPCServer{
		localService: newLocalService(ctx, cfg, logger),
	}
}

func (srv *RPCServer) PutObject(
	req PutObjectRequest,
	res *PutObjectResponse,
) error {
	fmt.Println(req)
	obj, err := srv.localService.objectmanager.Put(
		context.Background(),
		req.Obj,
	)
	if err != nil {
		return err
	}
	res.Obj = obj
	return nil
}
