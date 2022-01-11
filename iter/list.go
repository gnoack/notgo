package iter

func Slice[T any](s []T) Iter[T] {
	return &sliceIter[T]{s:s}
}

type sliceIter[T any] struct {
	s []T
	v T
}

func (s *sliceIter[T]) HasNext() bool {
	return len(s.s) > 0
}

func (s *sliceIter[T]) Next() bool {
	if len(s.s) == 0 {
		return false
	}
	s.v = s.s[0]
	s.s = s.s[1:]
	return true
}

func (s *sliceIter[T]) Value() T {
	return s.v
}

func ToSlice[T any](i Iter[T]) []T {
	var res []T
	for i.Next() {
		res = append(res, i.Value())
	}
	return res
}
