package iterator

type IIterator[T any] interface {
	Next() bool
	Value() T
	Reset()
	Collect() []T
}
