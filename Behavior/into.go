package behavior

//转换的接口根据转换数据从一种类型到另外一种类型
/// T:target type
type From[T any] interface {
	into() T
}
