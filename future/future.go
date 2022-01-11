// Package future comes from the future.
package future

type Fulfil[T any] func(T)

func New[T any]() (*Future[T], Fulfil[T]) {
	f := Future[T]{
		done: make(chan struct{}),
	}
	return &f, f.fulfil
}

func Go[T any](f func() T) *Future[T] {
	fut, fulfil := New[T]()
	go func() {
		v := f()
		fulfil(v)
	}()
	return fut
}

type Future[T any] struct {
	v T
	done chan struct{}
}

func (f*Future[T]) Value() T {
	<-f.done
	return f.v
}

func (f*Future[T]) fulfil(v T) {
	f.v = v
	close(f.done)
}
