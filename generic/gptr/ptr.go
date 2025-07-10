package gptr

func Of[T any](v T) *T {
	return &v
}

func Indirect[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

func Equal[T comparable](a, b *T) bool {
	return a == b || (a != nil && b != nil && *a == *b)
}
