package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFill(t *testing.T) {
	t.Run("non empty array", func(t *testing.T) {
		arr := []int{1, 2, 3}
		Fill(arr, 0)

		assert.Equal(t, []int{0, 0, 0}, arr)
	})

	t.Run("empty array", func(t *testing.T) {
		var arr []int
		Fill(arr, 0)

		assert.Equal(t, []int(nil), arr)
	})
}

func TestCopyArray(t *testing.T) {
	t.Run("empty array", func(tt *testing.T) {
		var arr []int
		copied := CopyArray(arr)

		assert.Equal(tt, arr, copied)
		assert.NotSame(tt, &arr, &copied)
	})

	t.Run("normal case", func(tt *testing.T) {
		arr := []int{1, 2, 3}
		copied := CopyArray(arr)

		assert.Equal(tt, arr, copied)
		assert.NotSame(tt, &arr, &copied)
	})

	t.Run("normal case (2)", func(tt *testing.T) {
		arr := []int{1, 2, 3}
		copied := CopyArray(arr)
		arr[1] = 5

		assert.Equal(tt, []int{1, 5, 3}, arr)
		assert.Equal(tt, []int{1, 2, 3}, copied)
		assert.NotSame(tt, &arr, &copied)
	})
}

func TestUnion(t *testing.T) {
	t.Run("single array", func(tt *testing.T) {
		arr1 := []int{1, 2, 3}
		union := Union(arr1)

		assert.Equal(tt, arr1, union)
	})
	t.Run("single empty array", func(tt *testing.T) {
		var arr1 []int
		union := Union(arr1)

		assert.Equal(tt, arr1, union)
	})
	t.Run("multi empty array", func(tt *testing.T) {
		var arr1 []int
		var arr2 []int
		union := Union(arr1, arr2)

		assert.Equal(tt, arr1, union)
		assert.Equal(tt, arr2, union)
	})
	t.Run("multi array", func(tt *testing.T) {
		arr1 := []int{1, 2}
		arr2 := []int{3, 4, 5}
		union := Union(arr1, arr2)

		assert.Equal(tt, []int{1, 2, 3, 4, 5}, union)
	})
	t.Run("multi array (2)", func(tt *testing.T) {
		arr1 := []int{1, 2}
		arr2 := []int{3, 4, 5}
		arr3 := []int{6}
		var arr4 []int
		arr5 := []int{7, 8, 9, 10, 11}
		union := Union(arr1, arr2, arr3, arr4, arr5)

		assert.Equal(tt, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, union)
	})
}

func TestForEach(t *testing.T) {
	var arr []int
	fn := func(x int) {
		arr = append(arr, x)
	}
	data := []int{1, 2, 3}

	ForEach(data, fn)

	assert.Equal(t, data, arr)
	assert.NotSame(t, &data, &arr)
}

func TestMin(t *testing.T) {
	t.Run("empty array", func(tt *testing.T) {
		var arr1, arr2 []int

		result := Min(arr1, arr2)

		assert.Equal(tt, 0, result)
	})

	t.Run("array zero values", func(tt *testing.T) {
		var arr1 []int
		arr2 := []int{0, 0, 0}

		result := Min(arr1, arr2)

		assert.Equal(tt, 0, result)
	})

	t.Run("case negative numbers", func(tt *testing.T) {
		arr1 := []int{-1, -2}
		arr2 := []int{-2, -5, -3}

		result := Min(arr1, arr2)

		assert.Equal(tt, -5, result)
	})

	t.Run("case positive numbers", func(tt *testing.T) {
		arr1 := []int{3, 2, 1}
		arr2 := []int{5, 2, 3, 4}

		result := Min(arr1, arr2)

		assert.Equal(tt, 1, result)
	})

	t.Run("case all numbers", func(tt *testing.T) {
		arr1 := []int{3, -2, 1, 7, 11}
		arr2 := []int{5, 2, 0, -1}

		result := Min(arr1, arr2)

		assert.Equal(tt, -2, result)
	})

	t.Run("case string", func(tt *testing.T) {
		arr1 := []string{"ab", "b", "a"}
		arr2 := []string{"c", "defg"}

		result := Min(arr1, arr2)

		assert.Equal(tt, "a", result)
	})
}

func TestMax(t *testing.T) {
	t.Run("empty array", func(tt *testing.T) {
		var arr1, arr2 []string

		result := Max(arr1, arr2)

		assert.Equal(tt, "", result)
	})

	t.Run("array zero values", func(tt *testing.T) {
		var arr1 []string
		arr2 := []string{"", "", ""}

		result := Max(arr1, arr2)

		assert.Equal(tt, "", result)
	})

	t.Run("case negative numbers", func(tt *testing.T) {
		arr1 := []int{-5, -2}
		arr2 := []int{-6, -5, -1, -3}

		result := Max(arr1, arr2)

		assert.Equal(tt, -1, result)
	})

	t.Run("case positive numbers", func(tt *testing.T) {
		arr1 := []int{3, 2, 1}
		arr2 := []int{5, 2, 3, 4}

		result := Max(arr1, arr2)

		assert.Equal(tt, 5, result)
	})

	t.Run("case all numbers", func(tt *testing.T) {
		arr1 := []int{3, -2, 1, 7, 11}
		arr2 := []int{5, 2, 0, -1}

		result := Max(arr1, arr2)

		assert.Equal(tt, 11, result)
	})

	t.Run("case string", func(tt *testing.T) {
		arr1 := []string{"ab", "b", "a"}
		arr2 := []string{"c", "defg"}

		result := Max(arr1, arr2)

		assert.Equal(tt, "defg", result)
	})
}

func TestCutArray(t *testing.T) {
	t.Run("empty array", func(tt *testing.T) {
		var arr []string

		result := CutArray(arr, 0)

		assert.Equal(tt, []string(nil), result)
	})

	t.Run("one element array", func(tt *testing.T) {
		arr := []string{"asd"}

		result := CutArray(arr, 0)

		assert.Equal(tt, []string(nil), result)
	})

	t.Run("two elements array - head", func(tt *testing.T) {
		arr := []string{"asd", "def"}

		result := CutArray(arr, 0)

		assert.Equal(tt, []string{"def"}, result)
	})

	t.Run("two elements array - tail", func(tt *testing.T) {
		arr := []string{"asd", "def"}

		result := CutArray(arr, 1)

		assert.Equal(tt, []string{"asd"}, result)
	})

	t.Run("three elements array - head", func(tt *testing.T) {
		arr := []string{"asd", "mid", "def"}

		result := CutArray(arr, 0)

		assert.Equal(tt, []string{"mid", "def"}, result)
	})

	t.Run("three elements array - tail", func(tt *testing.T) {
		arr := []string{"asd", "mid", "def"}

		result := CutArray(arr, 2)

		assert.Equal(tt, []string{"asd", "mid"}, result)
	})

	t.Run("three elements array - mid", func(tt *testing.T) {
		arr := []string{"asd", "mid", "def"}

		result := CutArray(arr, 1)

		assert.Equal(tt, []string{"asd", "def"}, result)
	})

	t.Run("N elements array (1)", func(tt *testing.T) {
		arr := []string{"asd", "mid", "huag", "koal", "tek", "def"}

		result := CutArray(arr, 3)
		arr[0] = "beda"

		assert.Equal(tt, []string{"beda", "mid", "huag", "koal", "tek", "def"}, arr)
		assert.Equal(tt, []string{"asd", "mid", "huag", "tek", "def"}, result)
	})

	t.Run("N elements array (2)", func(tt *testing.T) {
		arr := []string{"asd", "mid", "huag", "koal", "tek", "def"}

		result := CutArray(arr, 0)
		arr[0] = "beda"

		assert.Equal(tt, []string{"beda", "mid", "huag", "koal", "tek", "def"}, arr)
		assert.Equal(tt, []string{"mid", "huag", "koal", "tek", "def"}, result)
	})

}

func TestFind(t *testing.T) {
	t.Run("empty array", func(tt *testing.T) {
		var arr []string

		result, found := Find(arr, func(x string) bool { return x == "halo" })

		assert.Equal(tt, -1, result)
		assert.False(tt, found)
	})

	t.Run("case negative", func(tt *testing.T) {
		arr := []string{"hey", "hai", "hao"}

		result, found := Find(arr, func(x string) bool { return x == "halo" })

		assert.Equal(tt, -1, result)
		assert.False(tt, found)
	})

	t.Run("case positive", func(tt *testing.T) {
		arr := []string{"hey", "halo", "hai", "hao"}

		result, found := Find(arr, func(x string) bool { return x == "halo" })

		assert.Equal(tt, 1, result)
		assert.True(tt, found)
	})
}

func TestFindAndCut(t *testing.T) {
	t.Run("empty array", func(tt *testing.T) {
		var arr []string

		idx, result, found := FindAndCut(arr, func(x string) bool { return x == "halo" })

		assert.Equal(tt, -1, idx)
		assert.Equal(tt, []string(nil), result)
		assert.False(tt, found)
	})

	t.Run("case negative", func(tt *testing.T) {
		arr := []string{"hey", "hai", "hao"}

		idx, result, found := FindAndCut(arr, func(x string) bool { return x == "halo" })
		arr[0] = "uy"

		assert.Equal(tt, -1, idx)
		assert.Equal(tt, []string{"hey", "hai", "hao"}, result)
		assert.Equal(tt, []string{"uy", "hai", "hao"}, arr)
		assert.False(tt, found)
	})

	t.Run("case positive", func(tt *testing.T) {
		arr := []string{"hey", "halo", "hai", "hao"}

		idx, result, found := FindAndCut(arr, func(x string) bool { return x == "halo" })
		arr[0] = "uy"

		assert.Equal(tt, 1, idx)
		assert.Equal(tt, []string{"hey", "hai", "hao"}, result)
		assert.Equal(tt, []string{"uy", "halo", "hai", "hao"}, arr)
		assert.True(tt, found)
	})
}
