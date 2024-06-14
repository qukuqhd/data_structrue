package behavior

//option类型

type Option[T any] struct {
	val T
	set bool
}

//得到一个有值的option
func Some[T any](val T) Option[T] {
	return Option[T]{
		val: val,
		set: true,
	}
}

//得到一个空的option
func None[T any]() Option[T] {
	return Option[T]{
		set: false,
	}
}

//判断option是否是有值的
func (opt *Option[T]) IsSome() bool {
	return opt.set
}

//判断option是否是空的
func (opt *Option[T]) IsNone() bool {
	return !opt.set
}

//解包option如果是空的就panic
func (opt *Option[T]) Unwrap() T {
	if !opt.set {
		panic("unwrap option is none！")
	} else {
		return opt.val
	}
}

//如果option为空使用默认值否则使用自己的值
func (opt *Option[T]) Or(defaultVal T) T {
	if opt.set {
		return opt.val
	} else {
		return defaultVal
	}
}
