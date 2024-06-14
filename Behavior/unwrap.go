package behavior

type Unwrap[T any] interface {
	Unwrap() T
	Or(defaultVal T) T
}
