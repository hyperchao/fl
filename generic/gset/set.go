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

func (s Set[T]) IsSubSet(other Set[T]) bool {
	for v := range s {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

func (s Set[T]) IsSuperSet(other Set[T]) bool {
	return other.IsSubSet(s)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := New[T]()
	for v := range s {
		result.Add(v)
	}
	for v := range other {
		result.Add(v)
	}
	return result
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	result := New[T]()
	for v := range s {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}
