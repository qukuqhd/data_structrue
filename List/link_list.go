package list

type SinglyLinkNode[T any] struct {
	Value T
	Next  *SinglyLinkNode[T]
}

type DoublyLinkNode[T any] struct {
	Value T
	Next  *DoublyLinkNode[T]
	Pre   *DoublyLinkNode[T]
}

type SinglyLinkQueue struct {
}
type SinglyLinkStack struct {
}
type SinglyLinkList struct {
}

type DoublyLinkQueue struct {
}
type DoublyLinkStack struct {
}
type DoublyLinkList struct {
}
