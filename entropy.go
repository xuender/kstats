package kstats

import (
	"golang.org/x/exp/constraints"
)

// Entropy returns entropy of floats.
func Entropy[F constraints.Float](floats ...F) F {
	Normalize(floats)

	var result F

	for _, float := range floats {
		if float == 0 {
			continue
		}

		result += (float * Log(float))
	}

	return -result
}

// Normalize floats to normalized.
func Normalize[F constraints.Float](floats []F) {
	sum := Sum(floats...)
	for index, float := range floats {
		floats[index] = float / sum
	}
}
