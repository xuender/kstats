package kstats

import (
	"golang.org/x/exp/constraints"
)

// ProbGeom returns the probability for a geometric random variable
// with parameter float to achieve success in the interval of [start, end] trials
// See https://en.wikipedia.org/wiki/Geometric_distribution for more information.
func ProbGeom[F constraints.Float](start int, end int, float F) F {
	if (start > end) || (start < 1) {
		return Empty[F]()
	}

	var (
		prob F
		q    = 1 - float
	)

	for k := start + 1; k <= end; k++ {
		prob += float * Pow(q, F(k-1))
	}

	return prob
}

// ProbGeom returns the expectation or average number of trials
// for a geometric random variable with parameter float.
func ExpGeom[F constraints.Float](float F) F {
	if (float > 1) || (float < 0) {
		return Empty[F]()
	}

	return 1 / float
}

// ProbGeom returns the variance for number for a
// geometric random variable with parameter float.
func VarGeom[F constraints.Float](float F) F {
	if (float > 1) || (float < 0) {
		return Empty[F]()
	}

	return (1 - float) / Pow(float, _two)
}
