package list_test

import (
	list "data_structure/List"
	"testing"
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
