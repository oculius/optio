package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifferenceSet(t *testing.T) {
	t.Run("empty case", func(tt *testing.T) {
		var arr1 []float32
		var arr2 []float32

		diff := DifferenceSet(arr1, arr2)

		assert.Equal(tt, []float32(nil), diff)
		assert.Equal(tt, 0, len(diff))
	})

	t.Run("when second array is empty", func(tt *testing.T) {
		arr1 := []float32{1.0, 6.5}
		var arr2 []float32

		diff := DifferenceSet(arr1, arr2)
		arr1[1] = 3.3

		assert.Equal(tt, []float32{1.0, 6.5}, diff)
		assert.Equal(tt, []float32{1.0, 3.3}, arr1)
		assert.NotSame(tt, &arr1, &diff)
	})

	t.Run("when first array is empty", func(tt *testing.T) {
		var arr1 []float32
		arr2 := []float32{3.3, 6.7, 5.5}

		diff := DifferenceSet(arr1, arr2)
		arr2[0] = 1.2

		assert.Equal(tt, []float32{3.3, 6.7, 5.5}, diff)
		assert.Equal(tt, []float32{1.2, 6.7, 5.5}, arr2)
		assert.NotSame(tt, &arr2, &diff)
	})

	t.Run("normal case", func(tt *testing.T) {
		arr1 := []string{"hai", "hhai", "hello", "hola"}
		arr2 := []string{"hai", "hola", "halo", "heya"}

		diff := DifferenceSet(arr1, arr2)

		assert.Equal(tt, []string{"hhai", "hello"}, diff)
	})

	t.Run("negative case", func(tt *testing.T) {
		arr1 := []int{1, 2, 3}
		arr2 := []int{3, 2, 1}

		diff := DifferenceSet(arr1, arr2)

		assert.Equal(tt, []int(nil), diff)
		assert.Equal(tt, 0, len(diff))
	})

	t.Run("duplicate case", func(tt *testing.T) {
		arr1 := []float64{1, 1, 2.5, 2.5, 3, 3}
		arr2 := []float64{1, 2.5, 4.2}

		diff := DifferenceSet(arr1, arr2)

		assert.Equal(tt, []float64{3}, diff)
	})
}

func TestIntersectSet(t *testing.T) {
	t.Run("empty case", func(tt *testing.T) {
		var arr1 []float32
		var arr2 []float32

		intersect := IntersectSet(arr1, arr2)

		assert.Equal(tt, []float32(nil), intersect)
		assert.Equal(tt, 0, len(intersect))
	})

	t.Run("when second array is empty", func(tt *testing.T) {
		arr1 := []float32{1.0, 6.5}
		var arr2 []float32

		intersect := IntersectSet(arr1, arr2)

		assert.Equal(tt, []float32(nil), intersect)
		assert.Equal(tt, 0, len(intersect))
	})

	t.Run("when first array is empty", func(tt *testing.T) {
		var arr1 []float32
		arr2 := []float32{3.3, 6.7, 5.5}

		intersect := IntersectSet(arr1, arr2)

		assert.Equal(tt, []float32(nil), intersect)
		assert.Equal(tt, 0, len(intersect))
	})

	t.Run("normal case", func(tt *testing.T) {
		arr1 := []string{"hai", "hhai", "hello", "hola"}
		arr2 := []string{"hai", "hola", "halo", "heya"}

		intersect := IntersectSet(arr1, arr2)

		assert.Equal(tt, []string{"hai", "hola"}, intersect)
	})

	t.Run("negative case", func(tt *testing.T) {
		arr1 := []int{1, 3, 5, 6}
		arr2 := []int{2, 4, 7, 11, 13}

		intersect := IntersectSet(arr1, arr2)

		assert.Equal(tt, []int(nil), intersect)
		assert.Equal(tt, 0, len(intersect))
	})

	t.Run("duplicate case", func(tt *testing.T) {
		arr1 := []float64{1, 3.2, 5, 6, 11.5, 11.5, 13}
		arr2 := []float64{2.8, 4, 7, 11.5, 13}

		intersect := IntersectSet(arr1, arr2)

		assert.Equal(tt, []float64{11.5, 13}, intersect)
	})
}

func TestDifferenceUnionSet(t *testing.T) {
	arr1 := []float64{1, 3.2, 5, 6, 11.5, 11.5, 13}
	arr2 := []float64{2.8, 4, 7, 11.5, 13}

	intersect := DifferenceUnionSet(arr1, arr2)

	assert.Equal(t, []float64{1, 3.2, 5, 6, 2.8, 4, 7}, intersect)
}
