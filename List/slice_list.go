package list

import (
	behavior "data_structure/Behavior"
	"data_structure/fp"
	"errors"
)

// 切片实现的栈
type SliceStack[T any] struct {
	value []T
}

func NewSliceStack[T any]() Stack[T] {
	return &SliceStack[T]{
		value: make([]T, 0),
	}
}

func (stack *SliceStack[T]) Push(val T) {
	(*stack).value = append((*stack).value, val)
} //入栈

func (stack *SliceStack[T]) Pop() behavior.Option[T] {
	if len((*stack).value) == 0 {
		return behavior.None[T]()
	} else {
		topElement := ((*stack).value)[len((*stack).value)-1]
		(*stack).value = ((*stack).value)[:len((*stack).value)-1]
		return behavior.Some(topElement)
	}
} //出栈

// 基于切片的队列
type SliceQueue[T any] struct {
	value []T
}

func NewSliceQueue[T any]() Queue[T] {
	return &SliceQueue[T]{
		value: make([]T, 0),
	}
}

func (queue *SliceQueue[T]) Dequeue() behavior.Option[T] {
	if len((*queue).value) != 0 {
		firstElement := (*queue).value[0]
		(*queue).value = ((*queue).value)[1:]
		return behavior.Some[T](firstElement)
	}
	return behavior.None[T]()
} //出队列

func (queue *SliceQueue[T]) PushDequeue(val T) {
	(*queue).value = append((*queue).value, val)
} //入队列

type SliceList[T any] struct {
	*SliceQueue[T]
	*SliceStack[T]
	valueList []T
}

// 创建新的基于切片的线性表
func NewSliceList[T any]() List[T] {
	slice := make([]T, 0)
	return &SliceList[T]{
		SliceQueue: &SliceQueue[T]{value: slice},
		SliceStack: &SliceStack[T]{value: slice},
		valueList:  slice,
	}
}

func (list *SliceList[T]) Dequeue() behavior.Option[T] {
	return list.SliceQueue.Dequeue()
} //出队列
func (list *SliceList[T]) PushDequeue(val T) {
	list.SliceQueue.PushDequeue(val)
} //入队列

func (list *SliceList[T]) Push(val T) {
	list.SliceStack.Push(val)
} //入栈
func (list *SliceList[T]) Pop() behavior.Option[T] {
	return list.SliceStack.Pop()
} //出栈

func (list *SliceList[T]) Insert(index int, val T) behavior.Result[int] {
	sliceLen := len((*list).valueList)
	if index < 0 || index > sliceLen {
		return behavior.Err[int](errors.New("out of index"))
	} else if index == sliceLen {
		(*list).valueList = append((*list).valueList, val)
		//同步切片
		list.SliceQueue.value = (*list).valueList
		list.SliceStack.value = (*list).valueList
		return behavior.Ok[int](index)
	} else {
		left_slice := ((*list).valueList)[:index]
		right_slice := ((*list).valueList[index+1:])
		left_slice = append(left_slice, val)
		(*list).valueList = append(left_slice, right_slice...)
		//同步切片
		list.SliceQueue.value = (*list).valueList
		list.SliceStack.value = (*list).valueList
		return behavior.Ok[int](index)
	}
} //插入方法
func (list *SliceList[T]) Remove(index int) behavior.Option[T] {
	sliceLen := len((*list).valueList)
	if index < 0 || index >= sliceLen {
		return behavior.None[T]()
	} else {
		selectVal := list.valueList[index]
		left_slice := ((*list).valueList)[:index]
		right_slice := ((*list).valueList[index+1:])
		(*list).valueList = append(left_slice, right_slice...)
		//同步切片
		list.SliceQueue.value = (*list).valueList
		list.SliceStack.value = (*list).valueList
		return behavior.Some[T](selectVal)
	}
} //移除方法
func (list *SliceList[T]) Find(index int) behavior.Option[T] {
	sliceLen := len((*list).valueList)
	if index < 0 || index >= sliceLen {
		return behavior.None[T]()
	} else {
		selectVal := list.valueList[index]
		return behavior.Some[T](selectVal)
	}
}

type SliceIterator[T any] struct {
	currentIndex int
	sliceValues  *SliceList[T]
}

func (list *SliceList[T]) Iter() fp.Iterator[T] {
	return &SliceIterator[T]{
		currentIndex: 0,
		sliceValues:  list,
	}
}

func (iter *SliceIterator[T]) Next() bool {
	if iter.currentIndex >= len((*iter).sliceValues.valueList) || iter.currentIndex < 0 {
		return false
	} else {
		iter.currentIndex++
		return true
	}
} //往后面迭代
func (iter *SliceIterator[T]) Front() bool {
	if iter.currentIndex >= len((*iter).sliceValues.valueList) || iter.currentIndex < 0 {
		return false
	} else {
		iter.currentIndex--
		return true
	}
} //往前面迭代
func (iter *SliceIterator[T]) Value() T {
	return iter.sliceValues.Find(iter.currentIndex).Unwrap()
} //获取当前的值
func (iter *SliceIterator[T]) ReSet() fp.Iterator[T] {
	iter.currentIndex = 0
	return iter
} //回到迭代的起点
func (iter *SliceIterator[T]) Insert(val T) {
	iter.sliceValues.Insert(iter.currentIndex, val)
} //添加元素到当前位置
func (iter *SliceIterator[T]) Remove() T {
	return iter.sliceValues.Remove(iter.currentIndex).Unwrap()
} //移除当前的元素
