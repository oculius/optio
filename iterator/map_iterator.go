package iterator

type mapIterator[T any, K any] struct {
	source IIterator[T]
	mapper MapFunction[T, K]
}

type MapFunction[T any, K any] func(T) K

func (it *mapIterator[T, K]) Next() bool {
	return it.source.Next()
}

func (it *mapIterator[T, K]) Value() K {
	return it.mapper(it.source.Value())
}

func (it *mapIterator[T, K]) Reset() {
	it.source.Reset()
	*it = mapIterator[T, K]{source: it.source, mapper: it.mapper}
}

func (it *mapIterator[T, K]) Collect() []K {
	rawResult := it.source.Collect()
	N := len(rawResult)
	result := make([]K, N)
	for i := 0; i < N; i++ {
		result[i] = it.mapper(rawResult[i])
	}
	return result
}

func NewMapIter[T any, K any](iter IIterator[T], f MapFunction[T, K]) IIterator[K] {
	iter.Reset()
	return &mapIterator[T, K]{iter, f}
}

func NewMapIterFromArr[T any, K any](arr []T, f MapFunction[T, K]) IIterator[K] {
	iter := NewIterator(arr)
	return &mapIterator[T, K]{iter, f}
}
