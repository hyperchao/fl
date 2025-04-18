package idgen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_V7_Generate(t *testing.T) {
	id := NewV7().Generate()
	assert.NotEmpty(t, id)
	t.Log(id)
}
