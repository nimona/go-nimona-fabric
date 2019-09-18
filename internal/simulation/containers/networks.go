package containers

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Network struct {
	ID string
}

func NewNetwork(ctx context.Context, name string) (*Network, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	network, err := cli.NetworkCreate(
		ctx,
		name,
		types.NetworkCreate{
			CheckDuplicate: true,
			Driver:         "bridge",
		},
	)
	if err != nil {
		return nil, err
	}

	return &Network{ID: network.ID}, nil
}

func (n *Network) Remove(ctx context.Context) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	if err := cli.NetworkRemove(ctx, n.ID); err != nil {
		return err
	}

	return nil
}
