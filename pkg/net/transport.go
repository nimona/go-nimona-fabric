package net

import (
	"context"
)

type Transport interface {
	Dial(ctx context.Context, address string) (*Connection, error)
	Listen(ctx context.Context) (chan *Connection, error)
}
