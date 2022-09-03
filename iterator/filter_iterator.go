package iterator

type FilterFunction[T any] func(T) bool

type filterIterator[T any] struct {
	source IIterator[T]
	pred   FilterFunction[T]
}

func NewFilterIter[T any](iter IIterator[T], f FilterFunction[T]) IIterator[T] {
	return &filterIterator[T]{iter, f}
}

func NewFilterIterFromArr[T any](arr []T, f FilterFunction[T]) IIterator[T] {
	return &filterIterator[T]{source: NewIterator(arr), pred: f}
}

func (f *filterIterator[T]) Next() bool {
	for f.source.Next() {
		if f.pred(f.source.Value()) {
			return true
		}
	}
	return false
}

func (f *filterIterator[T]) Value() T {
	return f.source.Value()
}

func (f *filterIterator[T]) Reset() {
	f.source.Reset()
	*f = filterIterator[T]{source: f.source, pred: f.pred}
}

func (f *filterIterator[T]) Collect() []T {
	rawResult := f.source.Collect()
	N := len(rawResult)
	var result []T
	for i := 0; i < N; i++ {
		if f.pred(rawResult[i]) {
			result = append(result, rawResult[i])
		}
	}
	return result
}
