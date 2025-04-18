package glist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Find(t *testing.T) {
	l := NewList[int]()
	l.PushBack(12)
	e := l.Find(func(v int) bool {
		return v == 12
	})
	assert.Equal(t, 12, e.Value())
}
