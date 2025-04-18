package statemachine

import (
	"github.com/hyperchao/fl/generic/gfunc"
	"github.com/hyperchao/fl/generic/gslice"
)

type FSM[T comparable] map[T][]T

func (m FSM[T]) Transition(old T, new T) bool {
	return gslice.Any(m[old], gfunc.Equal(new))
}
