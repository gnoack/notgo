package iter

func Map[T any, U any](i Iter[T], f func(T) U) Iter[U] {
	return &mappingIter[T, U]{i:i, f:f}
}

type mappingIter[T any, U any] struct {
	i Iter[T]
	f func(T) U
}

func (i *mappingIter[T, U]) HasNext() bool {
	return i.i.HasNext()
}

func (i *mappingIter[T, U]) Next() bool {
	return i.i.Next()
}

func (i *mappingIter[T, U]) Value() U {
	return i.f(i.i.Value())
}
