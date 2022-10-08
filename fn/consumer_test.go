package fn

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsumer(t *testing.T) {
	sum := 0
	consumer := Consumer[int](func(i int) error {
		sum += i
		return nil
	})
	someError := errors.New("some error occured")
	errConsumer := Consumer[int](func(int) error {
		return someError
	})

	t.Run("when transformed error", func(tt *testing.T) {
		sum = 0
		var outerErr error
		errorHandler := func(err error) {
			outerErr = err
		}
		consumers := NewEmptyConsumer[int]().AndThen(errConsumer).AndThen(consumer)
		assert.Nil(tt, outerErr)
		silentConsumer := consumers.ToSilentConsumer(errorHandler)
		silentConsumer(5)
		assert.True(tt, errors.Is(outerErr, someError))
		assert.Equal(tt, 0, sum)
	})

	t.Run("when transformed no error", func(tt *testing.T) {
		sum = 0
		var outerErr error
		errorHandler := func(err error) {
			outerErr = err
		}
		consumers := NewEmptyConsumer[int]().AndThen(consumer).AndThen(consumer)
		assert.Nil(tt, outerErr)
		silentConsumer := consumers.ToSilentConsumer(errorHandler)
		silentConsumer(5)
		assert.Nil(tt, outerErr)
		assert.Equal(tt, 10, sum)
	})

	t.Run("when no error", func(tt *testing.T) {
		sum = 0

		consumers := NewEmptyConsumer[int]().AndThen(consumer).AndThen(consumer).AndThen(consumer)
		err := consumers(1)

		assert.Nil(tt, err)
		assert.Equal(tt, 3, sum)
	})

	t.Run("when there is error", func(tt *testing.T) {
		sum = 0

		consumers := NewEmptyConsumer[int]().AndThen(consumer).AndThen(consumer).AndThen(errConsumer).AndThen(consumer)
		err := consumers(1)

		assert.NotNil(tt, err)
		assert.True(tt, errors.Is(err, someError))
		assert.Equal(tt, 2, sum)
	})

	t.Run("when all error", func(tt *testing.T) {
		sum = 0

		consumers := NewEmptyConsumer[int]().AndThen(errConsumer).AndThen(errConsumer)
		err := consumers(1)

		assert.NotNil(tt, err)
		assert.True(tt, errors.Is(err, someError))
		assert.Equal(tt, 0, sum)
	})
}

func TestSilentConsumer(t *testing.T) {
	var result string
	consumer := SilentConsumer[string](func(s string) {
		result += s
	})

	t.Run("single consumer", func(tt *testing.T) {
		result = ""
		assert.Len(tt, result, 0)
		consumer("hai")
		assert.Equal(tt, result, "hai")
	})

	t.Run("multi consumer", func(tt *testing.T) {
		result = ""
		assert.Len(tt, result, 0)
		consumers := NewEmptySilentConsumer[string]().AndThen(consumer).AndThen(consumer)
		consumers("hai")
		assert.Equal(tt, result, "haihai")
	})
}

func TestBiConsumer(t *testing.T) {
	sumI := 0
	sumF := float64(0)
	consumer := BiConsumer[int, float64](func(i int, f float64) error {
		sumI += i
		sumF += f
		return nil
	})
	someError := errors.New("some error occured")
	errConsumer := BiConsumer[int, float64](func(int, float64) error {
		return someError
	})

	t.Run("when transformed error", func(tt *testing.T) {
		sumI = 0
		sumF = 0
		var outerErr error
		errorHandler := func(err error) {
			outerErr = err
		}
		consumers := NewEmptyBiConsumer[int, float64]().AndThen(errConsumer).AndThen(consumer)
		assert.Nil(tt, outerErr)
		silentConsumer := consumers.ToSilentConsumer(errorHandler)
		silentConsumer(5, 1.2)
		assert.True(tt, errors.Is(outerErr, someError))
		assert.Equal(tt, 0, sumI)
		assert.Equal(tt, float64(0), sumF)
	})

	t.Run("when transformed no error", func(tt *testing.T) {
		sumI = 0
		sumF = 0
		var outerErr error
		errorHandler := func(err error) {
			outerErr = err
		}
		consumers := NewEmptyBiConsumer[int, float64]().AndThen(consumer).AndThen(consumer)
		assert.Nil(tt, outerErr)
		silentConsumer := consumers.ToSilentConsumer(errorHandler)
		silentConsumer(5, 1.3)
		assert.Nil(tt, outerErr)
		assert.Equal(tt, 10, sumI)
		assert.InDelta(tt, float64(2.6), sumF, 0.001)
	})

	t.Run("when no error", func(tt *testing.T) {
		sumI = 0
		sumF = 0

		consumers := NewEmptyBiConsumer[int, float64]().
			AndThen(consumer).AndThen(consumer).AndThen(consumer)
		err := consumers(1, 1.2)

		assert.Nil(tt, err)
		assert.Equal(tt, 3, sumI)
		assert.InDelta(tt, float64(3.6), sumF, 0.001)
	})

	t.Run("when there is error", func(tt *testing.T) {
		sumI = 0
		sumF = 0

		consumers := NewEmptyBiConsumer[int, float64]().
			AndThen(consumer).AndThen(consumer).AndThen(errConsumer).AndThen(consumer)
		err := consumers(1, 1.2)

		assert.NotNil(tt, err)
		assert.True(tt, errors.Is(err, someError))
		assert.Equal(tt, 2, sumI)
		assert.InDelta(tt, float64(2.4), sumF, 0.001)
	})

	t.Run("when all error", func(tt *testing.T) {
		sumI = 0
		sumF = 0

		consumers := NewEmptyBiConsumer[int, float64]().AndThen(errConsumer).AndThen(errConsumer)
		err := consumers(1, 1.2)

		assert.NotNil(tt, err)
		assert.True(tt, errors.Is(err, someError))
		assert.Equal(tt, 0, sumI)
		assert.InDelta(tt, float64(0), sumF, 0.001)
	})
}

func TestSilentBiConsumer(t *testing.T) {
	var result string
	var result2 int
	consumer := SilentBiConsumer[string, int](func(s string, i int) {
		result += s
		result2 += i
	})

	t.Run("single consumer", func(tt *testing.T) {
		result = ""
		result2 = 0
		assert.Len(t, result, 0)
		assert.Equal(tt, result2, 0)
		consumer("hai", 3)
		assert.Equal(tt, result, "hai")
		assert.Equal(tt, result2, 3)
	})

	t.Run("multi consumer", func(tt *testing.T) {
		result = ""
		result2 = 0
		assert.Len(t, result, 0)
		assert.Equal(tt, result2, 0)
		consumers := NewEmptySilentBiConsumer[string, int]().AndThen(consumer).AndThen(consumer)
		consumers("hai", 3)
		assert.Equal(tt, result, "haihai")
		assert.Equal(tt, result2, 6)
	})
}
