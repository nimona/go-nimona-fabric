package object

import (
	"nimona.io/pkg/context"
)

type (
	Getter interface {
		Get(
			context.Context,
			Hash,
		) (*Object, error)
	}
	// GetterFunc is an adapter to allow the use of ordinary functions as
	// object.Getter
	GetterFunc func(
		context.Context,
		Hash,
	) (*Object, error)
)
