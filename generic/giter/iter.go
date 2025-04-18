package giter

import "iter"

func Map[T, R any](seq iter.Seq[T], f func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func Filter[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}
