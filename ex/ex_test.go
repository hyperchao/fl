package ex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryFinally(t *testing.T) {
	called := 0
	cleaned := 0

	Try(func() {
		called++
	}).Finally(func() {
		cleaned++
	})

	assert.Equal(t, 1, called)
	assert.Equal(t, 1, cleaned)
}

type myErr string

func (e myErr) Error() string {
	return string(e)
}

func TestTryCatchFinally(t *testing.T) {
	called := 0
	cleaned := 0
	catched := 0
	var myErrVal myErr
	var err error

	Try(func() {
		called++
		Trow(myErr("test1"))
	}).Catch(func(err myErr) {
		catched++
		myErrVal = err
	}).Finally(func() {
		cleaned++
	})

	assert.Equal(t, 1, called)
	assert.Equal(t, 1, catched)
	assert.Equal(t, 1, cleaned)
	assert.Equal(t, "test1", myErrVal.Error())

	Try(func() {
		Trow(myErr("test2"))
		called++
	}).Catch(func(_ string) {
		catched++
	}).Catch(func(e error) {
		err = e
	}).Catch(func(i any) {
		catched++
	}).Finally(nil)

	assert.Equal(t, 1, called)
	assert.Equal(t, 1, catched)
	assert.Equal(t, 1, cleaned)
	assert.Equal(t, "test2", err.Error())
}
