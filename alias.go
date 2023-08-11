package kstats

import "golang.org/x/exp/constraints"

// Variance returns amount of variation in the floats.
func Variance[F constraints.Float](floats ...F) F {
	return VarianceBySample(0, floats...)
}

// PopulationVariance returns the amount of floats within a population.
func PopulationVariance[F constraints.Float](floats ...F) F {
	return VarianceBySample(0, floats...)
}

// SampleVariance returns the amount of floats within a sample.
func SampleVariance[F constraints.Float](floats ...F) F {
	return VarianceBySample(1, floats...)
}

// MedianAbsoluteDeviation returns the median of the absolute deviations from the floats median.
func MedianAbsoluteDeviation[F constraints.Float](floats ...F) F {
	return MedianAbsoluteDeviationPopulation(floats...)
}

// StandardDeviation returns the amount of variation in the floats.
func StandardDeviation[F constraints.Float](floats ...F) F {
	return StandardDeviationPopulation(floats...)
}
