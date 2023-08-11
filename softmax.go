package kstats

import (
	"golang.org/x/exp/constraints"
)

// SoftMax returns the input values in the range of 0 to 1
// with sum of all the probabilities being equal to one. It
// is commonly used in machine learning neural networks.
func SoftMax[F constraints.Float](floats ...F) []F {
	length := len(floats)
	if length == 0 {
		return nil
	}

	var (
		sum F
		max = Max(floats...)
	)

	for _, e := range floats {
		sum += Exp(e - max)
	}

	softMax := make([]F, length)
	for i, v := range floats {
		softMax[i] = Exp(v-max) / sum
	}

	return softMax
}
