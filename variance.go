package kstats

import (
	"golang.org/x/exp/constraints"
)

// VarianceBySample return amount of variation in the floats by sample.
func VarianceBySample[F constraints.Float](sample int, floats ...F) F {
	length := len(floats)
	if length == 0 {
		return Empty[F]()
	}

	var (
		variance F
		mean     = Mean(floats...)
	)

	for _, float := range floats {
		variance += (float - mean) * (float - mean)
	}

	return variance / F(length-sample)
}

// Covariance returns a measure of how much two floats of data change.
func Covariance[F constraints.Float](floats1, floats2 []F) F {
	var (
		length1 = len(floats1)
		length2 = len(floats2)
	)

	if length1 == 0 || length2 == 0 || length1 != length2 {
		return Empty[F]()
	}

	var (
		mean1 = Mean(floats1...)
		mean2 = Mean(floats2...)
		sum   F
	)

	for index := 0; index < length1; index++ {
		delta1 := floats1[index] - mean1
		delta2 := floats2[index] - mean2
		sum += (delta1*delta2 - sum) / F(index+1)
	}

	return sum * F(length1) / F(length1-1)
}

// CovariancePopulation returns covariance for entire population between two floats.
func CovariancePopulation[F constraints.Float](floats1, floats2 []F) F {
	var (
		length1 = len(floats1)
		length2 = len(floats2)
	)

	if length1 == 0 || length2 == 0 || length1 != length2 {
		return Empty[F]()
	}

	var (
		mean1 = Mean(floats1...)
		mean2 = Mean(floats2...)
		sum   F
	)

	for index := 0; index < length1; index++ {
		delta1 := floats1[index] - mean1
		delta2 := floats2[index] - mean2
		sum += delta1 * delta2
	}

	return sum / F(length1)
}
