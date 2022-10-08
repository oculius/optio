package fn

type Supplier[T any] func() (T, error)
type SilentSupplier[T any] func() T
type BiSupplier[T, V any] func() (T, V, error)
type SilentBiSupplier[T, V any] func() (T, V)
