package gfunc

func Equal[T comparable](val T) func(T) bool {
	return func(input T) bool {
		return val == input
	}
}

func Identity[T any]() func(T) T {
	return func(input T) T {
		return input
	}
}
