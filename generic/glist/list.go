package glist

import (
	"container/list"
)

type List[T any] struct {
	l *list.List
}

type Element[T any] struct {
	e *list.Element
}

func (e *Element[T]) Value() T {
	return e.e.Value.(T)
}

func (e *Element[T]) Next() *Element[T] {
	if e.e.Next() == nil {
		return nil
	}
	return &Element[T]{e.e.Next()}
}

func (e *Element[T]) Prev() *Element[T] {
	if e.e.Prev() == nil {
		return nil
	}
	return &Element[T]{e.e.Prev()}
}

func NewList[T any]() *List[T] {
	return &List[T]{list.New()}
}

func (l *List[T]) Len() int {
	return l.l.Len()
}

func (l *List[T]) Init() {
	l.l.Init()
}

func (l *List[T]) PushFront(v T) *Element[T] {
	return &Element[T]{
		l.l.PushFront(v),
	}
}

func (l *List[T]) PushBack(v T) *Element[T] {
	return &Element[T]{
		l.l.PushBack(v),
	}
}

func (l *List[T]) Front() *Element[T] {
	if l.l.Front() == nil {
		return nil
	}
	return &Element[T]{l.l.Front()}
}

func (l *List[T]) Back() *Element[T] {
	if l.l.Back() == nil {
		return nil
	}
	return &Element[T]{l.l.Back()}
}

func (l *List[T]) Remove(e *Element[T]) T {
	return l.l.Remove(e.e).(T)
}

func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	return &Element[T]{l.l.InsertAfter(v, mark.e)}
}

func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	return &Element[T]{l.l.InsertBefore(v, mark.e)}
}

func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.l.MoveBefore(e.e, mark.e)
}

func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	l.l.MoveAfter(e.e, mark.e)
}

func (l *List[T]) PushBackList(other *List[T]) {
	l.l.PushBackList(other.l)
}

func (l *List[T]) PushFrontList(other *List[T]) {
	l.l.PushFrontList(other.l)
}

func (l *List[T]) Find(predict func(T) bool) *Element[T] {
	for e := l.l.Front(); e != nil; e = e.Next() {
		if predict(e.Value.(T)) {
			return &Element[T]{e}
		}
	}
	return nil
}
