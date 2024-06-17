package list_test

import (
	list "data_structure/List"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceStack(t *testing.T) {
	stack := list.NewSliceStack[int]()
	stack.Push(55)
	stack.Push(66)
	stack.Push(78)
	if stack.Pop().Unwrap() != 78 {
		panic("stack err")
	}
	if stack.Pop().Unwrap() != 66 {
		panic("stack err")
	}
	if stack.Pop().Unwrap() != 55 {
		panic("stack err")
	}
	if stack.Pop().IsSome() {
		panic("stack err")
	}
}
func TestQueue(t *testing.T) {
	queue := list.NewSliceQueue[int]()
	queue.PushDequeue(55)
	queue.PushDequeue(66)
	queue.PushDequeue(78)
	if queue.Dequeue().Unwrap() != 55 {
		panic("queue err")
	}
	if queue.Dequeue().Unwrap() != 66 {
		panic("queue err")
	}
	if queue.Dequeue().Unwrap() != 78 {
		panic("queue err")
	}
	if queue.Dequeue().IsSome() {
		panic("queue err")
	}
}

func TestList(t *testing.T) {
	a := assert.New(t)
	list := list.NewSliceList[int]()
	res := list.Insert(55, 45)
	a.Equal(res.IsOk(), false)
	res = list.Insert(0, 55)
	a.Equal(res.IsOk(), true)
	res = list.Insert(0, 44)
	a.Equal(res.IsOk(), true)
	res = list.Insert(0, 33)
	a.Equal(res.IsOk(), true)
	res = list.Insert(0, 22)
	a.Equal(res.IsOk(), true)
}
