package gptr

func Of[T any](v T) *T {
	return &v
}

func Equal[T comparable](a, b *T) bool {
	return a == b || (a != nil && b != nil && *a == *b)
}
