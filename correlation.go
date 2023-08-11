package kstats

import (
	"golang.org/x/exp/constraints"
)

// Correlation returns the degree of relationship between two sets of floats.
func Correlation[F constraints.Float](floats1, floats2 []F) F {
	var (
		length1 = len(floats1)
		length2 = len(floats2)
	)

	if length1 == 0 || length2 == 0 || length1 != length2 {
		return Empty[F]()
	}

	var (
		sdev1 = StandardDeviationPopulation(floats1...)
		sdev2 = StandardDeviationPopulation(floats2...)
	)

	if sdev1 == 0 || sdev2 == 0 {
		return 0
	}

	covp := CovariancePopulation(floats1, floats2)

	return covp / (sdev1 * sdev2)
}

// AutoCorrelation returns the correlation of a signal with a delayed copy of itself as a function of delay.
func AutoCorrelation[F constraints.Float](lags int, data ...F) F {
	length := len(data)
	if length < 1 {
		return Empty[F]()
	}

	var (
		mean   = Mean(data...)
		result F
		mol    F
	)

	for lag := 0; lag < lags; lag++ {
		den := (data[0] - mean) * (data[0] - mean)

		for index := 1; index < len(data); index++ {
			var (
				delta0 = data[index-1] - mean
				delta1 = data[index] - mean
			)

			mol += (delta0*delta1 - mol) / F(index+1)
			den += (delta1*delta1 - den) / F(index+1)
		}

		result = mol / den
	}

	return result
}
