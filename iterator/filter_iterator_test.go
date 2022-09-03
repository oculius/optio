package iterator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFilterIterator_Integration(t *testing.T) {
	t.Run("when predicate all true for all elements", func(tt *testing.T) {
		iter := NewFilterIterFromArr([]int{1, 2, 3, 4}, func(x int) bool { return true })

		assert.Equal(tt, []int{1, 2, 3, 4}, iter.Collect())
		assert.Zero(tt, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, 1, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, 2, iter.Value())
		iter.Reset()
		assert.Zero(tt, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, 1, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, 2, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, 3, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, 4, iter.Value())
		assert.False(tt, iter.Next())
	})

	t.Run("when predicate true for half elements", func(tt *testing.T) {
		iter := NewFilterIterFromArr([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })

		assert.Equal(tt, []int{2, 4}, iter.Collect())
		assert.Zero(tt, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, 2, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, 4, iter.Value())
		assert.False(tt, iter.Next())
	})

	t.Run("nested filter", func(tt *testing.T) {
		firstFilter := NewFilterIterFromArr(
			[]string{"hey", "train", "trample", "after", "say"},
			func(x string) bool {
				return len(x) > 3
			},
		)
		assert.Equal(tt, []string{"train", "trample", "after"}, firstFilter.Collect())

		secondFilter := NewFilterIter(firstFilter,
			func(x string) bool {
				return strings.HasPrefix(x, "tr")
			},
		)

		assert.Equal(tt, []string{"train", "trample"}, secondFilter.Collect())
	})
}
