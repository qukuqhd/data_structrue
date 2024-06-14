package fp

//迭代器接口
//T 迭代获取元素的类型
type Iterator[T any] interface {
	Next() bool
	Value() T
}

//转换数据结构为迭代对象的接口
type Iterable[T any] interface {
	iter() Iterator[T]
}
