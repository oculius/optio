package iterator

type Iterator[T any] struct {
	elements []T
	value    T
	index    int
	maxIndex int
}

func NewIterator[T any](arr []T) IIterator[T] {
	return &Iterator[T]{
		elements: arr,
		index:    0,
		maxIndex: len(arr),
	}
}

func (it *Iterator[T]) Next() bool {
	if it.index < it.maxIndex {
		it.value = it.elements[it.index]
		it.index++
		return true
	}
	return false
}

func (it *Iterator[T]) Value() T {
	return it.value
}

func (it *Iterator[T]) Reset() {
	*it = Iterator[T]{
		elements: it.elements,
		index:    0,
		maxIndex: len(it.elements),
	}
}

func (it *Iterator[T]) Collect() []T {
	result := make([]T, it.maxIndex)
	for i := 0; i < it.maxIndex; i++ {
		result[i] = it.elements[i]
	}
	return result
}
