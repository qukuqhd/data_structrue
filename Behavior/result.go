package behavior

type Result[T any] struct {
	Value T
	Error error
}

func Ok[T any](val T) Result[T] {
	return Result[T]{
		Value: val,
		Error: nil,
	}
}
func Err[T any](err error) Result[T] {
	return Result[T]{
		Error: err,
	}
}
func (r Result[T]) Unwrap() T {
	if r.Error == nil {
		return r.Value
	} else {
		panic("result is err! Error is " + r.Error.Error())
	}
}
func (r Result[T]) Or(defaultVal T) T {
	if r.Error == nil {
		return r.Value
	} else {
		return defaultVal
	}
}
func (r Result[T]) IsOk() bool {
	if r.Error != nil {
		return false
	} else {
		return true
	}
}
