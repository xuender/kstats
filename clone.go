package kstats

func Clone[T any](elems []T) []T {
	ret := make([]T, len(elems))

	copy(ret, elems)

	return ret
}
