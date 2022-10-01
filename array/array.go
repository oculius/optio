package array

import (
	"golang.org/x/exp/constraints"
	"optio/fn"
)

func Fill[T any](arr []T, value T) {
	for i := range arr {
		arr[i] = value
	}
}

func CopyArray[T any](arr []T) []T {
	if arr == nil {
		return nil
	}

	result := make([]T, len(arr))
	copy(result, arr)
	return result
}

func CutArray[T any](arr []T, idx int) []T {
	N := len(arr)
	if N == 0 || idx < 0 || idx >= N || (N == 1 && idx == 0) {
		return nil
	} else if idx == 0 {
		return CopyArray(arr[(idx + 1):])
	} else if idx == (N - 1) {
		return CopyArray(arr[:idx])
	}

	result := make([]T, 0, N-1)
	result = append(result, arr[:idx]...)
	result = append(result, arr[(idx+1):]...)
	return result
}

func Union[T any](arrays ...[]T) []T {
	N := 0
	for i := range arrays {
		N += len(arrays[i])
	}

	if N == 0 {
		return nil
	}

	result := make([]T, N)
	c := 0
	for i := range arrays {
		for j := 0; j < len(arrays[i]); j++ {
			result[c] = arrays[i][j]
			c++
		}
	}
	return result
}

func ForEach[T any](arr []T, fn fn.SilentConsumer[T]) {
	for i := range arr {
		fn(arr[i])
	}
}

func Max[T constraints.Ordered](arr ...[]T) T {
	var max T
	if len(arr) > 0 {
		if len(arr[0]) > 0 {
			max = arr[0][0]
		}
	}
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] > max {
				max = arr[i][j]
			}
		}
	}
	return max
}

func Min[T constraints.Ordered](arr ...[]T) T {
	var min T
	if len(arr) > 0 {
		if len(arr[0]) > 0 {
			min = arr[0][0]
		}
	}
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] < min {
				min = arr[i][j]
			}
		}
	}
	return min
}

func Find[T any](arr []T, pred fn.Predicate[T]) (idx int, found bool) {
	idx = -1
	for i, val := range arr {
		if pred(val) {
			found = true
			idx = i
			return
		}
	}
	return
}

func FindAndCut[T any](arr []T, pred fn.Predicate[T]) (idx int, result []T, found bool) {
	found = false
	idx = -1
	for i, val := range arr {
		if pred(val) {
			idx = i
			found = true
			result = CutArray(arr, i)
			return
		}
	}
	result = CopyArray(arr)
	return
}
