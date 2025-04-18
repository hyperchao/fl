package gset

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	return s[v] != struct{}{}
}

func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for v := range s {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

func (s Set[T]) Size() int {
	return len(s)
}
