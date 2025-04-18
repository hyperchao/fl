package idgen

import (
	"context"
	"github.com/google/uuid"
)

type idGeneratorKey struct{}

type IDGenerator interface {
	Generate() string
}

func WithIDGenerator(ctx context.Context, idGenerator IDGenerator) context.Context {
	return context.WithValue(ctx, idGeneratorKey{}, idGenerator)
}

func GetIDGenerator(ctx context.Context) IDGenerator {
	return ctx.Value(idGeneratorKey{}).(IDGenerator)
}

type v7 struct{}

func (v7) Generate() string {
	return uuid.Must(uuid.NewV7()).String()
}

func NewV7() IDGenerator {
	return v7{}
}
