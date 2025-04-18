package ex

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

var (
	ErrInvalidCatchType = fmt.Errorf("invalid catch type")
)

func Assert(pass bool, e error) {
	if !pass {
		Trow(errors.WithStack(e))
	}
}

func AssertE(e error) {
	if e != nil {
		Trow(errors.WithStack(e))
	}
}

func Trow(e error) {
	panic(e)
}

type Execution interface {
	Catch(any) Execution
	Finally(func())
}

type execution struct {
	try     func()
	catches []reflect.Value
	finally func()
}

func Try(f func()) Execution {
	return &execution{try: f}
}

func (e *execution) Catch(f any) Execution {
	rt := reflect.TypeOf(f)
	Assert(rt.Kind() == reflect.Func && rt.NumIn() == 1 && rt.NumOut() == 0, ErrInvalidCatchType)
	e.catches = append(e.catches, reflect.ValueOf(f))
	return e
}

func (e *execution) Finally(f func()) {
	e.finally = f
	e.do()
}

func (e *execution) do() {
	if e.finally != nil {
		defer e.finally()
	}

	defer func() {
		if r := recover(); r != nil {
			val := reflect.ValueOf(r)
			valT := val.Type()
			for _, catch := range e.catches {
				argT := catch.Type().In(0)
				if valT.AssignableTo(argT) {
					catch.Call([]reflect.Value{val})
					return
				}
			}
			panic(r)
		}
	}()

	e.try()
}
