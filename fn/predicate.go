package fn

type GenericPredicate[T any] interface {
	Negate() T
	Or(T) T
	And(T) T
	Xnor(T) T
	Xor(T) T
}

type ISilentPredicate[T any, V GenericPredicate[V]] interface {
	GenericPredicate[T]
	ToPredicate() V
}

type IPredicate[T any, V GenericPredicate[V]] interface {
	GenericPredicate[T]
	ToSilentPredicate(ErrorHandler) V
}

type SilentPredicate[T any] func(T) bool
type SilentBiPredicate[T, V any] func(T, V) bool
type Predicate[T any] func(T) (bool, error)
type BiPredicate[T, V any] func(T, V) (bool, error)

var _ ISilentPredicate[
	SilentPredicate[int], Predicate[int],
] = SilentPredicate[int](
	func(int) bool { return true },
)

var _ ISilentPredicate[
	SilentBiPredicate[int, string], BiPredicate[int, string],
] = SilentBiPredicate[int, string](
	func(int, string) bool { return true },
)

var _ IPredicate[
	Predicate[int], SilentPredicate[int],
] = Predicate[int](
	func(int) (bool, error) { return true, nil },
)

var _ IPredicate[
	BiPredicate[int64, float32], SilentBiPredicate[int64, float32],
] = BiPredicate[int64, float32](
	func(int64, float32) (bool, error) { return true, nil },
)

func (p SilentBiPredicate[T, V]) ToPredicate() BiPredicate[T, V] {
	return func(v1 T, v2 V) (bool, error) {
		result := p(v1, v2)
		return result, nil
	}
}

func (p SilentBiPredicate[T, V]) Negate() SilentBiPredicate[T, V] {
	return func(v1 T, v2 V) bool {
		result := p(v1, v2)
		return !result
	}
}

func (p SilentBiPredicate[T, V]) Or(other SilentBiPredicate[T, V]) SilentBiPredicate[T, V] {
	return func(v1 T, v2 V) bool {
		result := p(v1, v2)
		if result {
			return true
		}

		result2 := other(v1, v2)
		return result2
	}
}

func (p SilentBiPredicate[T, V]) And(other SilentBiPredicate[T, V]) SilentBiPredicate[T, V] {
	return func(v1 T, v2 V) bool {
		result := p(v1, v2)
		if !result {
			return false
		}

		result2 := other(v1, v2)
		return result2
	}
}

func (p SilentBiPredicate[T, V]) Xnor(other SilentBiPredicate[T, V]) SilentBiPredicate[T, V] {
	return func(v1 T, v2 V) bool {
		result := p(v1, v2)
		result2 := other(v1, v2)
		return result == result2
	}
}

func (p SilentBiPredicate[T, V]) Xor(other SilentBiPredicate[T, V]) SilentBiPredicate[T, V] {
	return func(v1 T, v2 V) bool {
		result := p(v1, v2)
		result2 := other(v1, v2)
		return result != result2
	}
}

func (p SilentPredicate[T]) ToPredicate() Predicate[T] {
	return func(v1 T) (bool, error) {
		result := p(v1)
		return result, nil
	}
}

func (p SilentPredicate[T]) And(other SilentPredicate[T]) SilentPredicate[T] {
	return func(v1 T) bool {
		result := p(v1)
		if !result {
			return false
		}

		result2 := other(v1)
		return result2
	}
}

func (p SilentPredicate[T]) Xnor(other SilentPredicate[T]) SilentPredicate[T] {
	return func(v1 T) bool {
		result := p(v1)
		result2 := other(v1)
		return result == result2
	}
}

func (p SilentPredicate[T]) Xor(other SilentPredicate[T]) SilentPredicate[T] {
	return func(v1 T) bool {
		result := p(v1)
		result2 := other(v1)
		return result != result2
	}
}

func (p SilentPredicate[T]) Negate() SilentPredicate[T] {
	return func(v1 T) bool {
		result := p(v1)
		return !result
	}
}

func (p SilentPredicate[T]) Or(other SilentPredicate[T]) SilentPredicate[T] {
	return func(v1 T) bool {
		result := p(v1)
		if result {
			return true
		}

		result2 := other(v1)
		return result2
	}
}

func (p BiPredicate[T, V]) ToSilentPredicate(errHandler ErrorHandler) SilentBiPredicate[T, V] {
	return func(v1 T, v2 V) bool {
		result, err := p(v1, v2)
		if err != nil {
			if errHandler != nil {
				errHandler(err)
			}
			return false
		}
		return result
	}
}

func (p BiPredicate[T, V]) Negate() BiPredicate[T, V] {
	return func(v1 T, v2 V) (bool, error) {
		result, err := p(v1, v2)
		if err != nil {
			return false, err
		}
		return !result, nil
	}
}

func (p BiPredicate[T, V]) Or(other BiPredicate[T, V]) BiPredicate[T, V] {
	return func(v1 T, v2 V) (bool, error) {
		result, err := p(v1, v2)
		if err != nil {
			return false, err
		}
		if result {
			return true, nil
		}

		result2, err2 := other(v1, v2)
		if err2 != nil {
			return false, err2
		}
		return result2, nil
	}
}

func (p BiPredicate[T, V]) And(other BiPredicate[T, V]) BiPredicate[T, V] {
	return func(v1 T, v2 V) (bool, error) {
		result, err := p(v1, v2)
		if err != nil {
			return false, err
		}

		if !result {
			return false, nil
		}

		result2, err2 := other(v1, v2)
		if err2 != nil {
			return false, err2
		}
		return result2, nil
	}
}

func (p BiPredicate[T, V]) Xnor(other BiPredicate[T, V]) BiPredicate[T, V] {
	return func(v1 T, v2 V) (bool, error) {
		result, err := p(v1, v2)
		if err != nil {
			return false, err
		}
		result2, err2 := other(v1, v2)
		if err2 != nil {
			return false, err2
		}
		return result == result2, nil
	}
}

func (p BiPredicate[T, V]) Xor(other BiPredicate[T, V]) BiPredicate[T, V] {
	return func(v1 T, v2 V) (bool, error) {
		result, err := p(v1, v2)
		if err != nil {
			return false, err
		}
		result2, err2 := other(v1, v2)
		if err2 != nil {
			return false, err2
		}
		return result != result2, nil
	}
}

func (p Predicate[T]) Negate() Predicate[T] {
	return func(v1 T) (bool, error) {
		result, err := p(v1)
		if err != nil {
			return false, err
		}
		return !result, nil
	}
}

func (p Predicate[T]) Or(other Predicate[T]) Predicate[T] {
	return func(v1 T) (bool, error) {
		result, err := p(v1)
		if err != nil {
			return false, err
		}
		if result {
			return true, nil
		}

		result2, err2 := other(v1)
		if err2 != nil {
			return false, err2
		}
		return result2, nil
	}
}

func (p Predicate[T]) And(other Predicate[T]) Predicate[T] {
	return func(v1 T) (bool, error) {
		result, err := p(v1)
		if err != nil {
			return false, err
		}

		if !result {
			return false, nil
		}

		result2, err2 := other(v1)
		if err2 != nil {
			return false, err2
		}
		return result2, nil
	}
}

func (p Predicate[T]) Xnor(other Predicate[T]) Predicate[T] {
	return func(v1 T) (bool, error) {
		result, err := p(v1)
		if err != nil {
			return false, err
		}
		result2, err2 := other(v1)
		if err2 != nil {
			return false, err2
		}
		return result == result2, nil
	}
}

func (p Predicate[T]) Xor(other Predicate[T]) Predicate[T] {
	return func(v1 T) (bool, error) {
		result, err := p(v1)
		if err != nil {
			return false, err
		}
		result2, err2 := other(v1)
		if err2 != nil {
			return false, err2
		}
		return result != result2, nil
	}
}

func (p Predicate[T]) ToSilentPredicate(errHandler ErrorHandler) SilentPredicate[T] {
	return func(v1 T) bool {
		result, err := p(v1)
		if err != nil {
			if errHandler != nil {
				errHandler(err)
			}
			return false
		}
		return result
	}
}
