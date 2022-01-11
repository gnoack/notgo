package iter

func Fold[T any, R any](i Iter[T], start R, f func(R, T) R) R {
	res := start
	for i.Next() {
		res = f(res, i.Value())
	}
	return res
}
