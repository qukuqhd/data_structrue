package list

import data_structure "data_structure/Behavior"

//切片实现的栈
type SliceStack[T any] []T

func NewSliceStack[T any]() SliceStack[T] {
	return make([]T, 0)
}

func (stack *SliceQueue[T]) Push(val T) {
	*stack = append(*stack, val)
} //入栈
func (stack *SliceQueue[T]) Pop() data_structure.Option[T] {
	if len(*stack) == 0 {
		return data_structure.None[T]()
	} else {
		topElement := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		return data_structure.Some(topElement)
	}
} //出栈
