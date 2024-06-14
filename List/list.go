package list

import behavior "data_structure/Behavior"

//队列接口的定义
type Queue[T any] interface {
	Dequeue() behavior.Option[T] //出队列
	Queue(T)                     //入队列
}

//栈接口的定义
type Stack[T any] interface {
	Push(T)                  //入栈
	Pop() behavior.Option[T] //出栈
}

type List[T any] interface {
	Stack[T]          //继承栈接口
	Queue[T]          //继承队列接口
	Insert(uint64, T) //插入方法
	Remove(uint64) T  //移除方法
}
