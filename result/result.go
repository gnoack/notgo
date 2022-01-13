package result

// Result wraps either a value or an error.
//
// This is a bad idea because it's already idiomatic in Go to do it
// differently (to return a value and an error from a function).
type Result[T any] struct {
	v   T
	err error
}

func Of[T any](v T) Result[T] {
	return Result[T]{v: v}
}

func Error[T any](e error) Result[T] {
	return Result[T]{err: e}
}

func (r Result[T]) Value() T {
	return r.v
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) Must() T {
	if r.Err() != nil {
		panic(r.Err())
	}
	return r.Value()
}
