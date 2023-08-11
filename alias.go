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

// Pearson returns the Pearson product-moment correlation coefficient between two variables.
func Pearson[F constraints.Float](floats1, floats2 []F) F {
	return Correlation(floats1, floats2)
}

// VarP is a shortcut to PopulationVariance.
func VarP[F constraints.Float](floats ...F) F {
	return PopulationVariance(floats...)
}

// VarS is a shortcut to SampleVariance.
func VarS[F constraints.Float](floats ...F) F {
	return SampleVariance(floats...)
}

// StdDevP is a shortcut to StandardDeviationPopulation.
func StdDevP[F constraints.Float](floats ...F) F {
	return StandardDeviationPopulation(floats...)
}

// StdDevS is a shortcut to StandardDeviationSample.
func StdDevS[F constraints.Float](floats ...F) F {
	return StandardDeviationSample(floats...)
}

// LinReg is a shortcut to LinearRegression.
func LinReg[N, F Number](s []Coordinate[N, F]) []Coordinate[N, float64] {
	return LinearRegression(s)
}

// ExpReg is a shortcut to ExponentialRegression.
func ExpReg[N, F Number](s []Coordinate[N, F]) []Coordinate[N, float64] {
	return ExponentialRegression(s)
}

// LogReg is a shortcut to LogarithmicRegression.
func LogReg[N, F Number](s []Coordinate[N, F]) []Coordinate[N, float64] {
	return LogarithmicRegression(s)
}
