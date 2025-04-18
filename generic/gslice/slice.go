package gslice

import (
	"github.com/hyperchao/fl/generic/giter"
	"github.com/hyperchao/fl/generic/optional"
	"iter"
)

func ForEach[T any](slice []T, fn func(T)) {
	for _, v := range slice {
		fn(v)
	}
}

func Filter[T any](slice []T, predict func(T) bool) (result []T) {
	for _, item := range slice {
		if predict(item) {
			result = append(result, item)
		}
	}
	return
}

func FindFirst[T any](slice []T, predict func(T) bool) optional.O[T] {
	for _, item := range slice {
		if predict(item) {
			return optional.Ok(item)
		}
	}
	return optional.Nil[T]()
}

func Contains[T any](slice []T, predict func(T) bool) bool {
	for _, item := range slice {
		if predict(item) {
			return true
		}
	}
	return false
}

func TypeAssert[To any, From any](slice []From) (result []To) {
	for _, item := range slice {
		result = append(result, any(item).(To))
	}
	return
}

func All[T any](slice []T, predict func(T) bool) bool {
	for _, item := range slice {
		if !predict(item) {
			return false
		}
	}
	return true
}

func Any[T any](slice []T, predict func(T) bool) bool {
	for _, item := range slice {
		if predict(item) {
			return true
		}
	}
	return false
}

func Iter[T any](slice []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range slice {
			if !yield(v) {
				return
			}
		}
	}
}

func FromIter[T any](seq iter.Seq[T]) (result []T) {
	for v := range seq {
		result = append(result, v)
	}
	return
}

func Map[T, R any](slice []T, f func(T) R) []R {
	return FromIter(giter.Map(Iter(slice), f))
}

func FilterMap[T, R any](slice []T, f func(T) bool, g func(T) R) []R {
	return FromIter(giter.Map(giter.Filter(Iter(slice), f), g))
}

func MapFilter[T, R any](slice []T, f func(T) R, g func(R) bool) []R {
	return FromIter(giter.Filter(giter.Map(Iter(slice), f), g))
}
