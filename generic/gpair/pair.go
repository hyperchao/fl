package gpair

type Pair[T, R any] struct {
	Left  T
	Right R
}

func Of[T, R any](left T, right R) Pair[T, R] {
	return Pair[T, R]{
		Left:  left,
		Right: right,
	}
}

func Left[T, R any](pair Pair[T, R]) T {
	return pair.Left
}

func Right[T, R any](pair Pair[T, R]) R {
	return pair.Right
}
