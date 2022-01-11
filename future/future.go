// Package future comes from the future.
package future

type Fulfil[T any] func(T)

// New returns a new Future[T] and a function that fulfils it.
func New[T any]() (*Future[T], Fulfil[T]) {
	f := Future[T]{
		done: make(chan struct{}),
	}
	return &f, f.fulfil
}

// Go runs the given function f in a background goroutine,
// and immediately returns a future that will be fulfilled when f returns.
func Go[T any](f func() T) *Future[T] {
	fut, fulfil := New[T]()
	go func() {
		v := f()
		fulfil(v)
	}()
	return fut
}

// Future represents a value of T that will become available at some
// point in the future.
type Future[T any] struct {
	v T
	done chan struct{}
}

// Value returns the future's value, waiting for fulfilment if needed.
func (f *Future[T]) Value() T {
	<-f.done
	return f.v
}

func (f *Future[T]) fulfil(v T) {
	f.v = v
	close(f.done)
}
