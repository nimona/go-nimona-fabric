package gomockutil

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"

	"nimona.io/pkg/object"
)

// TODO is this even needed after objects v6?

func ObjectEq(x *object.Object) gomock.Matcher {
	return objectEqMatcher{x}
}

type objectEqMatcher struct {
	x *object.Object
}

func (e objectEqMatcher) Matches(x interface{}) bool {
	o, ok := x.(*object.Object)
	if !ok {
		return false
	}
	return reflect.DeepEqual(
		e.x,
		o,
	)
}

func (e objectEqMatcher) String() string {
	return fmt.Sprintf("is equal to %v", e.x)
}
