package gonerics

func Slice[T any](e ...T) []T {
	return e
}

func Combine[T any](e ...[]T) []T {
	r := []T{}

	for _, v := range e {
		r = append(r, v...)
	}

	return r
}
