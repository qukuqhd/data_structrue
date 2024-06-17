package fp

//迭代器接口
//T 迭代获取元素的类型
type Iterator[T any] interface {
	Next() bool         //往后面迭代
	Front() bool        //往前面迭代
	Value() T           //获取当前的值
	ReSet() Iterator[T] //回到迭代的起点
	Insert(T)           //添加元素到当前位置
	Remove() T          //移除当前的元素
}

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type IteratorObj[T any] struct {
	Iterator[T]
}

type DefaultIter[T any] struct {
	val   []T
	index int
}

// Front implements Iterator.
func (d *DefaultIter[T]) Front() bool {
	if d.index <= 0 {
		return false
	} else {
		d.index--
		return true
	}
}

// Insert implements Iterator.
func (d *DefaultIter[T]) Insert(T) {
	panic("unimplemented")
}

// Next implements Iterator.
func (d *DefaultIter[T]) Next() bool {
	if d.index > len(d.val) {
		return false
	} else {
		d.index++
		return true
	}
}

// ReSet implements Iterator.
func (d *DefaultIter[T]) ReSet() Iterator[T] {
	d.index = 0
	return d
}

// Remove implements Iterator.
func (d *DefaultIter[T]) Remove() T {
	panic("unimplemented")
}

// Value implements Iterator.
func (d *DefaultIter[T]) Value() T {
	return d.val[d.index]
}

func newIter[T any]() Iterator[T] {
	return &DefaultIter[T]{
		val:   make([]T, 0),
		index: 0,
	}
}
func (iter IteratorObj[T]) Foreach(handle func(val any)) Iterator[T] {
	for iter.Next() {
		handle(iter.Value())
	}
	iter.ReSet()
	return iter
}

func (iter IteratorObj[T]) Filter(filter func(T) bool) Iterator[T] {
	for iter.Next() {
		if !filter(iter.Value()) {
			iter.Remove()
		}
	}
	return iter
}

type KVBucket[K, V any] struct {
	Key K
	Val V
}

func NewKVBucket[K, V any](key K, val V) KVBucket[K, V] {
	return KVBucket[K, V]{
		Key: key,
		Val: val,
	}
}
func Zip[T comparable, V any](keyIter Iterator[T], valIter Iterator[V]) Iterator[KVBucket[T, V]] {
	newIter := newIter[KVBucket[T, V]]()
	for keyIter.Next() && valIter.Next() {
		newIter.Insert(NewKVBucket(keyIter.Value(), valIter.Value()))
	}
	return newIter
}
func Map[T, V any](iter Iterator[T], map_func func(T) V) Iterator[V] {
	newIter := newIter[V]()
	for iter.Next() {
		newIter.Insert(map_func(iter.Value()))
	}
	return newIter
}
