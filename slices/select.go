package slices

// Select returns the slice with the elements e where f(e) was true.
func Select[T any](s []T, f func(T) bool) []T {
	var res []T
	for _, e := range s {
		if f(e) {
			res = append(res, e)
		}
	}
	return res
}

// Reject returns the slice with the elements e where f(e) was false.
// This is the opposite of Select().
func Reject[T any](s []T, f func(T) bool) []T {
	var res []T
	for _, e := range s {
		if !f(e) {
			res = append(res, e)
		}
	}
	return res
}
