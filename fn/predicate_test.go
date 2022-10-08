package fn

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPredicate(t *testing.T) {
	call := 0
	callTrue := 0
	callFalse := 0
	callError := 0
	someError := errors.New("some error occured")
	trueFn := Predicate[int](func(int) (bool, error) {
		callTrue += 1
		call += 1
		return true, nil
	})
	falseFn := Predicate[int](func(int) (bool, error) {
		callFalse += 1
		call += 1
		return false, nil
	})

	errorFn := Predicate[int](func(int) (bool, error) {
		call += 1
		callError += 1
		return false, someError
	})

	beforeEach := func() {
		call = 0
		callTrue = 0
		callFalse = 0
		callError = 0
	}

	type TestCase struct {
		Title      string
		Predicate  Predicate[int]
		Result     bool
		IsError    bool
		TotalCount int
		FalseCount int
		TrueCount  int
		ErrorCount int
	}

	transformer := func(tc TestCase) (string, func(*testing.T)) {
		return tc.Title, func(x *testing.T) {
			beforeEach()
			result, err := tc.Predicate(0)

			if tc.IsError {
				assert.True(x, errors.Is(err, someError))
			} else {
				assert.Nil(x, err)
			}

			assert.Equal(x, tc.Result, result)
			assert.Equal(x, tc.TotalCount, call)
			assert.Equal(x, tc.FalseCount, callFalse)
			assert.Equal(x, tc.TrueCount, callTrue)
			assert.Equal(x, tc.ErrorCount, callError)
		}
	}

	t.Run("Negate", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true function",
				Predicate:  trueFn.Negate(),
				Result:     false,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "false function",
				Predicate:  falseFn.Negate(),
				Result:     true,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "error",
				Predicate:  errorFn.Negate(),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("And", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.And(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.And(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.And(trueFn),
				Result:     false,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.And(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "error first",
				Predicate:  errorFn.And(trueFn),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
			{
				Title:      "error last",
				Predicate:  trueFn.And(errorFn),
				Result:     false,
				IsError:    true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Or", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Or(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Or(falseFn),
				Result:     true,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Or(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Or(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
				ErrorCount: 0,
			},
			{
				Title:      "error first",
				Predicate:  errorFn.Or(trueFn),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
			{
				Title:      "error last",
				Predicate:  trueFn.Or(errorFn),
				Result:     true,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 0,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Xor", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Xor(trueFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Xor(falseFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Xor(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Xor(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
				ErrorCount: 0,
			},
			{
				Title:      "error first",
				Predicate:  errorFn.Xor(trueFn),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
			{
				Title:      "error last",
				Predicate:  trueFn.Xor(errorFn),
				Result:     false,
				IsError:    true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Xnor", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Xnor(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Xnor(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Xnor(trueFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Xnor(falseFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
				ErrorCount: 0,
			},
			{
				Title:      "error first",
				Predicate:  errorFn.Xnor(trueFn),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
			{
				Title:      "error last",
				Predicate:  trueFn.Xnor(errorFn),
				Result:     false,
				IsError:    true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})
}

func TestBiPredicate(t *testing.T) {
	call := 0
	callTrue := 0
	callFalse := 0
	callError := 0
	someError := errors.New("some error occured")
	trueFn := BiPredicate[int, string](func(int, string) (bool, error) {
		callTrue += 1
		call += 1
		return true, nil
	})
	falseFn := BiPredicate[int, string](func(int, string) (bool, error) {
		callFalse += 1
		call += 1
		return false, nil
	})

	errorFn := BiPredicate[int, string](func(int, string) (bool, error) {
		call += 1
		callError += 1
		return false, someError
	})

	beforeEach := func() {
		call = 0
		callTrue = 0
		callFalse = 0
		callError = 0
	}

	type TestCase struct {
		Title      string
		Predicate  BiPredicate[int, string]
		Result     bool
		IsError    bool
		TotalCount int
		FalseCount int
		TrueCount  int
		ErrorCount int
	}

	transformer := func(tc TestCase) (string, func(*testing.T)) {
		return tc.Title, func(x *testing.T) {
			beforeEach()
			result, err := tc.Predicate(0, "test")

			if tc.IsError {
				assert.True(x, errors.Is(err, someError))
			} else {
				assert.Nil(x, err)
			}

			assert.Equal(x, tc.Result, result)
			assert.Equal(x, tc.TotalCount, call)
			assert.Equal(x, tc.FalseCount, callFalse)
			assert.Equal(x, tc.TrueCount, callTrue)
			assert.Equal(x, tc.ErrorCount, callError)
		}
	}

	t.Run("Negate", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true function",
				Predicate:  trueFn.Negate(),
				Result:     false,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "false function",
				Predicate:  falseFn.Negate(),
				Result:     true,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "error",
				Predicate:  errorFn.Negate(),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("And", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.And(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.And(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.And(trueFn),
				Result:     false,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.And(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "error first",
				Predicate:  errorFn.And(trueFn),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
			{
				Title:      "error last",
				Predicate:  trueFn.And(errorFn),
				Result:     false,
				IsError:    true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Or", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Or(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Or(falseFn),
				Result:     true,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Or(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Or(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
				ErrorCount: 0,
			},
			{
				Title:      "error first",
				Predicate:  errorFn.Or(trueFn),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
			{
				Title:      "error last",
				Predicate:  trueFn.Or(errorFn),
				Result:     true,
				IsError:    false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 0,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Xor", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Xor(trueFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Xor(falseFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Xor(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Xor(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
				ErrorCount: 0,
			},
			{
				Title:      "error first",
				Predicate:  errorFn.Xor(trueFn),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
			{
				Title:      "error last",
				Predicate:  trueFn.Xor(errorFn),
				Result:     false,
				IsError:    true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Xnor", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Xnor(trueFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
				ErrorCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Xnor(falseFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Xnor(trueFn),
				Result:     false,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
				ErrorCount: 0,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Xnor(falseFn),
				Result:     true,
				IsError:    false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
				ErrorCount: 0,
			},
			{
				Title:      "error first",
				Predicate:  errorFn.Xnor(trueFn),
				Result:     false,
				IsError:    true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 0,
				ErrorCount: 1,
			},
			{
				Title:      "error last",
				Predicate:  trueFn.Xnor(errorFn),
				Result:     false,
				IsError:    true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 0,
				ErrorCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})
}

func TestSilentPredicate(t *testing.T) {
	call := 0
	callTrue := 0
	callFalse := 0
	trueFn := SilentPredicate[int](func(int) bool {
		callTrue += 1
		call += 1
		return true
	})
	falseFn := SilentPredicate[int](func(int) bool {
		callFalse += 1
		call += 1
		return false
	})

	beforeEach := func() {
		call = 0
		callTrue = 0
		callFalse = 0
	}

	type TestCase struct {
		Title      string
		Predicate  SilentPredicate[int]
		Result     bool
		TotalCount int
		FalseCount int
		TrueCount  int
	}

	transformer := func(tc TestCase) (string, func(*testing.T)) {
		return tc.Title, func(x *testing.T) {
			beforeEach()
			result := tc.Predicate(0)

			assert.Equal(x, tc.Result, result)
			assert.Equal(x, tc.TotalCount, call)
			assert.Equal(x, tc.FalseCount, callFalse)
			assert.Equal(x, tc.TrueCount, callTrue)
		}
	}

	t.Run("Negate", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true function",
				Predicate:  trueFn.Negate(),
				Result:     false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
			},
			{
				Title:      "false function",
				Predicate:  falseFn.Negate(),
				Result:     true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("And", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.And(trueFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.And(falseFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.And(trueFn),
				Result:     false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.And(falseFn),
				Result:     false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Or", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Or(trueFn),
				Result:     true,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Or(falseFn),
				Result:     true,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Or(trueFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Or(falseFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Xor", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Xor(trueFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Xor(falseFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Xor(trueFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Xor(falseFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Xnor", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Xnor(trueFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Xnor(falseFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Xnor(trueFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Xnor(falseFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})
}

func TestSilentBiPredicate(t *testing.T) {
	call := 0
	callTrue := 0
	callFalse := 0
	trueFn := SilentBiPredicate[int, func()](func(int, func()) bool {
		callTrue += 1
		call += 1
		return true
	})
	falseFn := SilentBiPredicate[int, func()](func(int, func()) bool {
		callFalse += 1
		call += 1
		return false
	})

	beforeEach := func() {
		call = 0
		callTrue = 0
		callFalse = 0
	}

	type TestCase struct {
		Title      string
		Predicate  SilentBiPredicate[int, func()]
		Result     bool
		TotalCount int
		FalseCount int
		TrueCount  int
	}

	transformer := func(tc TestCase) (string, func(*testing.T)) {
		return tc.Title, func(x *testing.T) {
			beforeEach()
			result := tc.Predicate(0, func() {})

			assert.Equal(x, tc.Result, result)
			assert.Equal(x, tc.TotalCount, call)
			assert.Equal(x, tc.FalseCount, callFalse)
			assert.Equal(x, tc.TrueCount, callTrue)
		}
	}

	t.Run("Negate", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true function",
				Predicate:  trueFn.Negate(),
				Result:     false,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
			},
			{
				Title:      "false function",
				Predicate:  falseFn.Negate(),
				Result:     true,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("And", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.And(trueFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.And(falseFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.And(trueFn),
				Result:     false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.And(falseFn),
				Result:     false,
				TotalCount: 1,
				TrueCount:  0,
				FalseCount: 1,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Or", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Or(trueFn),
				Result:     true,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Or(falseFn),
				Result:     true,
				TotalCount: 1,
				TrueCount:  1,
				FalseCount: 0,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Or(trueFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Or(falseFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Xor", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Xor(trueFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Xor(falseFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Xor(trueFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Xor(falseFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})

	t.Run("Xnor", func(tt *testing.T) {
		testCases := []TestCase{
			{
				Title:      "true and true",
				Predicate:  trueFn.Xnor(trueFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  2,
				FalseCount: 0,
			},
			{
				Title:      "true and false",
				Predicate:  trueFn.Xnor(falseFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and true",
				Predicate:  falseFn.Xnor(trueFn),
				Result:     false,
				TotalCount: 2,
				TrueCount:  1,
				FalseCount: 1,
			},
			{
				Title:      "false and false",
				Predicate:  falseFn.Xnor(falseFn),
				Result:     true,
				TotalCount: 2,
				TrueCount:  0,
				FalseCount: 2,
			},
		}

		for _, each := range testCases {
			title, fn := transformer(each)
			tt.Run(title, fn)
		}
	})
}
