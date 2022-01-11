package iter

type Iter[T any] interface {
	HasNext() bool
	Next() bool
	Value() T
}
