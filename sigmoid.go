package kstats

import (
	"golang.org/x/exp/constraints"
)

// Sigmoid returns the input values in the range of -1 to 1
// along the sigmoid or s-shaped curve, commonly used in
// machine learning while training neural networks as an
// activation function.
func Sigmoid[F constraints.Float](floats ...F) []F {
	length := len(floats)
	if length == 0 {
		return nil
	}

	ret := make([]F, length)
	for index, float := range floats {
		ret[index] = 1 / (1 + Exp(-float))
	}

	return ret
}
