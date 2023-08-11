package kstats

import (
	"golang.org/x/exp/constraints"
)

// MedianAbsoluteDeviationPopulation returns the median of the absolute deviations from the population median.
func MedianAbsoluteDeviationPopulation[F constraints.Float](floats ...F) F {
	if length := len(floats); length == 0 {
		return Empty[F]()
	}

	var (
		clone = Clone(floats)
		media = Median(clone...)
	)

	for key, value := range clone {
		clone[key] = Abs(value - media)
	}

	return Median(clone...)
}

// StandardDeviationPopulation returns the amount of variation from the population.
func StandardDeviationPopulation[F constraints.Float](floats ...F) F {
	if length := len(floats); length == 0 {
		return Empty[F]()
	}

	return Sqrt(PopulationVariance(floats...))
}

// StandardDeviationSample returns the amount of variation from a sample.
func StandardDeviationSample[F constraints.Float](floats ...F) F {
	if length := len(floats); length == 0 {
		return Empty[F]()
	}

	return Sqrt(SampleVariance(floats...))
}
