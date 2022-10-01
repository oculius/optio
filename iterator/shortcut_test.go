package iterator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	result := Filter([]int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 == 1 })

	assert.Equal(t, 3, len(result))
	assert.Equal(t, []int{1, 3, 5}, result)
}

func TestMap(t *testing.T) {
	result := Map([]string{"abc", "defg", "hijkl"}, func(x string) int { return len(x) })

	assert.Equal(t, 3, len(result))
	assert.Equal(t, []int{3, 4, 5}, result)
}

func TestReduce(t *testing.T) {
	t.Run("homo-type reduce", func(tt *testing.T) {
		result := Reduce([]int{3, 4, 5, 6},
			func(accum int, x int) int {
				return accum + x
			})

		assert.Equal(tt, 18, result)
	})
	t.Run("hetero-type reduce", func(tt *testing.T) {
		result := Reduce([]string{"simple", "test", "yep"},
			func(accum int, x string) int {
				return accum + len(x)
			})

		assert.Equal(tt, 13, result)
	})
}

func TestFilterMapReduce(t *testing.T) {
	t.Run("compose filter map reduce test", func(tt *testing.T) {
		originalArr := []string{"hello", "s", "hey", "soul", "sister"}
		filtered := Filter(originalArr, func(x string) bool {
			return strings.HasPrefix(x, "s")
		})
		assert.Equal(tt, []string{"s", "soul", "sister"}, filtered)
		mapped := Map(filtered, func(x string) int {
			return len(x)
		})
		assert.Equal(tt, []int{1, 4, 6}, mapped)
		reduced := Reduce(mapped, func(result string, x int) string {
			buff := make([]byte, x)
			for a := range buff {
				buff[a] = 'a' + byte(x-1)
			}
			return result + string(buff)
		})
		assert.Equal(tt, "addddffffff", reduced)
	})
}
