package fn

type GenericConsumer[X any] interface {
	AndThen(X) X
}

type IConsumer[T any, V GenericConsumer[V]] interface {
	GenericConsumer[T]
	ToSilentConsumer(ErrorHandler) V
}

type ISilentConsumer[T any, V GenericConsumer[V]] interface {
	GenericConsumer[T]
	ToConsumer() V
}

type SilentConsumer[T any] func(T)
type SilentBiConsumer[T, V any] func(T, V)
type Consumer[T any] func(T) error
type BiConsumer[T, V any] func(T, V) error

var _ IConsumer[
	Consumer[int], SilentConsumer[int],
] = Consumer[int](
	func(int) error { return nil },
)

var _ IConsumer[
	BiConsumer[int, string], SilentBiConsumer[int, string],
] = BiConsumer[int, string](
	func(int, string) error { return nil },
)

var _ ISilentConsumer[
	SilentConsumer[int], Consumer[int],
] = SilentConsumer[int](
	func(int) {},
)

var _ ISilentConsumer[
	SilentBiConsumer[int, string], BiConsumer[int, string],
] = SilentBiConsumer[int, string](
	func(int, string) {},
)

func NewEmptyConsumer[T any]() Consumer[T] {
	return func(T) error { return nil }
}

func NewEmptyBiConsumer[T, V any]() BiConsumer[T, V] {
	return func(T, V) error { return nil }
}

func NewEmptySilentConsumer[T any]() SilentConsumer[T] {
	return func(T) {}
}

func NewEmptySilentBiConsumer[T, V any]() SilentBiConsumer[T, V] {
	return func(T, V) {}
}

func (bc SilentBiConsumer[T, V]) ToConsumer() BiConsumer[T, V] {
	return func(v1 T, v2 V) error {
		bc(v1, v2)
		return nil
	}
}

func (bc SilentBiConsumer[T, V]) AndThen(after SilentBiConsumer[T, V]) SilentBiConsumer[T, V] {
	if after == nil {
		return bc
	}

	return func(v1 T, v2 V) {
		bc(v1, v2)
		after(v1, v2)
	}
}

func (c SilentConsumer[T]) AndThen(after SilentConsumer[T]) SilentConsumer[T] {
	if after == nil {
		return c
	}

	return func(v1 T) {
		c(v1)
		after(v1)
	}
}

func (c SilentConsumer[T]) ToConsumer() Consumer[T] {
	return func(v1 T) error {
		c(v1)
		return nil
	}
}

func (c Consumer[T]) AndThen(after Consumer[T]) Consumer[T] {
	if after == nil {
		return c
	}

	return func(v1 T) error {
		err := c(v1)
		if err != nil {
			return err
		}
		return after(v1)
	}
}

func (c Consumer[T]) ToSilentConsumer(errHandler ErrorHandler) SilentConsumer[T] {
	return func(v1 T) {
		err := c(v1)
		if err != nil && errHandler != nil {
			errHandler(err)
		}
	}
}

func (bc BiConsumer[T, V]) ToSilentConsumer(errHandler ErrorHandler) SilentBiConsumer[T, V] {
	return func(v1 T, v2 V) {
		err := bc(v1, v2)
		if err != nil && errHandler != nil {
			errHandler(err)
		}
	}
}

func (bc BiConsumer[T, V]) AndThen(after BiConsumer[T, V]) BiConsumer[T, V] {
	if after == nil {
		return bc
	}

	return func(v1 T, v2 V) error {
		err := bc(v1, v2)
		if err != nil {
			return err
		}
		return after(v1, v2)
	}
}
