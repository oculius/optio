package iterator

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIterator_Integration(t *testing.T) {
	t.Run("when empty", func(tt *testing.T) {
		iter := NewIterator([]int{})

		assert.False(tt, iter.Next())
		assert.Zero(tt, iter.Value())
	})

	t.Run("when single", func(tt *testing.T) {
		iter := NewIterator([]string{"hai"})

		assert.Zero(tt, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, iter.Value(), "hai")
		assert.False(tt, iter.Next())
	})

	t.Run("when multidata and multitype", func(tt *testing.T) {
		iter := NewIterator([]interface{}{1, "lah", 0.5})

		assert.Zero(tt, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, iter.Value(), 1)
		assert.True(tt, iter.Next())
		assert.Equal(tt, iter.Value(), "lah")
		assert.True(tt, iter.Next())
		assert.Equal(tt, iter.Value(), 0.5)
		assert.False(tt, iter.Next())
		assert.Equal(tt, iter.Value(), 0.5)
		assert.False(tt, iter.Next())
	})

	t.Run("when reseted", func(tt *testing.T) {
		iter := NewIterator([]interface{}{1, "lah", 0.5})

		assert.Zero(tt, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, iter.Value(), 1)
		assert.True(tt, iter.Next())
		assert.Equal(tt, iter.Value(), "lah")
		iter.Reset()
		assert.Zero(tt, iter.Value())
		assert.True(tt, iter.Next())
		assert.Equal(tt, iter.Value(), 1)
		assert.True(tt, iter.Next())
		assert.Equal(tt, iter.Value(), "lah")
		assert.True(tt, iter.Next())
		assert.Equal(tt, iter.Value(), 0.5)
		assert.False(tt, iter.Next())
		assert.Equal(tt, iter.Value(), 0.5)
		assert.False(tt, iter.Next())
	})
}

func randArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Int()
	}
	return arr
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterBytesN = len(letterBytes)

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(letterBytesN)]
	}
	return string(b)
}

func randArrayString(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = randString(16)
	}
	return s
}

type sampleStruct struct {
	Name string
	ID   int
}

func newSampleStruct() sampleStruct {
	return sampleStruct{
		Name: randString(16),
		ID:   rand.Int(),
	}
}

func randomArrayStruct(n int) []sampleStruct {
	as := make([]sampleStruct, n)
	for i := 0; i < n; i++ {
		as[i] = newSampleStruct()
	}
	return as
}

func BenchmarkIterator_Integration(b *testing.B) {
	b.Run("integer array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			iter := NewIterator(randArray(500))
			for iter.Next() {
				iter.Value()
			}
		}
	})

	b.Run("string array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			iter := NewIterator(randArrayString(500))
			for iter.Next() {
				iter.Value()
			}
		}
	})

	b.Run("big struct array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			iter := NewIterator(randomArrayStruct(500))
			for iter.Next() {
				iter.Value()
			}
		}
	})
}
