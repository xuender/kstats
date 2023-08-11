package kstats

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Empty[T any]() T {
	var zero T

	return zero
}
