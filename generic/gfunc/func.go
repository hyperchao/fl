package gfunc

func Equal[T comparable](val T) func(T) bool {
	return func(input T) bool {
		return val == input
	}
}
