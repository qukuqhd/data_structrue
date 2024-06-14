package list

import behavior "data_structure/Behavior"

//基于切片的队列
type SliceQueue[T any] []T

func NewSliceQueue[T any]() SliceQueue[T] {
	return make([]T, 0)
}
func (queue *SliceQueue[T]) Dequeue() behavior.Option[T] {
	if len(*queue) != 0 {
		firstElement := (*queue)[0]
		*queue = (*queue)[1:]
		return behavior.Some[T](firstElement)
	}
	return behavior.None[T]()
} //出队列
func (queue *SliceQueue[T]) Queue(val T) {
	*queue = append(*queue, val)
} //入队列
