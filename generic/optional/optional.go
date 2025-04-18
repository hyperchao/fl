package optional

type O[T any] struct {
	val T
	ok  bool
}

func OfPtr[T any](ptr *T) O[T] {
	if ptr == nil {
		return Nil[T]()
	}
	return Ok(*ptr)
}

func Nil[T any]() O[T] {
	return O[T]{}
}

func Ok[T any](val T) O[T] {
	return O[T]{
		val: val,
		ok:  true,
	}
}

func (o O[T]) Ptr() *T {
	if o.ok {
		return &o.val
	}
	return nil
}

func (o O[T]) Val() T {
	return o.val
}

func (o O[T]) Ok() bool {
	return o.ok
}

func (o O[T]) ValOr(val T) T {
	if o.ok {
		return o.val
	}
	return val
}
