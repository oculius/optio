package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapIterator_Integration(to *testing.T) {
	to.Run("map hetero-type from array constructor", func(t *testing.T) {
		mapResult := NewMapIterFromArr[int, float64]([]int{3, 7}, func(t int) float64 {
			return float64(t) / 2
		})

		assert.Zero(t, mapResult.Value())
		assert.True(t, mapResult.Next())
		assert.Equal(t, 1.5, mapResult.Value())
		assert.Equal(t, []float64{1.5, 3.5}, mapResult.Collect())
		assert.True(t, mapResult.Next())
		assert.Equal(t, 3.5, mapResult.Value())
		assert.False(t, mapResult.Next())
	})

	to.Run("map hetero-type", func(t *testing.T) {
		mapResult := NewMapIter[string, int](NewIterator([]string{"hai", "gegeh", ""}), func(t string) int {
			return len(t)
		})

		assert.Zero(t, mapResult.Value())
		assert.True(t, mapResult.Next())
		assert.Equal(t, 3, mapResult.Value())
		assert.True(t, mapResult.Next())
		assert.Equal(t, 5, mapResult.Value())
		assert.True(t, mapResult.Next())
		assert.Equal(t, 0, mapResult.Value())
		assert.False(t, mapResult.Next())
		assert.Equal(t, []int{3, 5, 0}, mapResult.Collect())
	})

	to.Run("nested map hetero-type", func(t *testing.T) {
		firstMap := NewMapIterFromArr([]string{"yes", "no", "release"}, func(t string) int {
			return len(t)
		})
		secondMap := NewMapIter(firstMap, func(t int) float64 {
			return float64(t) / 2
		})

		assert.Zero(t, secondMap.Value())
		assert.Equal(t, []float64{1.5, 1, 3.5}, secondMap.Collect())
		assert.True(t, secondMap.Next())
		assert.Equal(t, secondMap.Value(), 1.5)
		assert.True(t, secondMap.Next())
		assert.Equal(t, secondMap.Value(), 1.0)
		secondMap.Reset()
		assert.Zero(t, secondMap.Value())
		assert.True(t, secondMap.Next())
		assert.Equal(t, secondMap.Value(), 1.5)
		assert.True(t, secondMap.Next())
		assert.Equal(t, secondMap.Value(), 1.0)
		assert.True(t, secondMap.Next())
		assert.Equal(t, secondMap.Value(), 3.5)
		assert.False(t, secondMap.Next())
	})

	to.Run("map homo-type", func(t *testing.T) {
		mapResult := NewMapIter[int, int](NewIterator([]int{1, 2, 3}), func(t int) int {
			return t * 3
		})

		assert.Zero(t, mapResult.Value())
		assert.True(t, mapResult.Next())
		assert.Equal(t, 3, mapResult.Value())
		assert.True(t, mapResult.Next())
		assert.Equal(t, 6, mapResult.Value())
		assert.True(t, mapResult.Next())
		assert.Equal(t, 9, mapResult.Value())
		assert.False(t, mapResult.Next())
		assert.Equal(t, mapResult.Collect(), []int{3, 6, 9})
	})

	to.Run("map struct hetero-type", func(t *testing.T) {
		randomStruct := randomArrayStruct(3)
		mapResult := NewMapIterFromArr[sampleStruct, string](randomStruct, func(t sampleStruct) string {
			return t.Name
		})

		assert.Zero(t, mapResult.Value())
		assert.True(t, mapResult.Next())
		collected := mapResult.Collect()
		assert.Equal(t, collected[0], mapResult.Value())
		assert.True(t, mapResult.Next())
		assert.Equal(t, collected[1], mapResult.Value())
		assert.True(t, mapResult.Next())
		assert.Equal(t, collected[2], mapResult.Value())
		assert.False(t, mapResult.Next())
	})
}
