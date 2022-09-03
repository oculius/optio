package iterator

func Map[T any, K any](arr []T, fn MapFunction[T, K]) []K {
	return NewMapIterFromArr(arr, fn).Collect()
}

func Filter[T any](arr []T, fn FilterFunction[T]) []T {
	return NewFilterIterFromArr(arr, fn).Collect()
}

type ReducerFunction[T any, K any] func(K, T) K

func Reduce[T any, K any](arr []T, fn ReducerFunction[T, K]) K {
	var result K
	iter := NewIterator(arr)

	for iter.Next() {
		result = fn(result, iter.Value())
	}

	return result
}
