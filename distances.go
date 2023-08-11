package kstats

import "math"

// Validate data for distance calculation.
func validateData[N Number](dataPointX, dataPointY []N) bool {
	if len(dataPointX) == 0 || len(dataPointY) == 0 {
		return true
	}

	if len(dataPointX) != len(dataPointY) {
		return true
	}

	return false
}

// ChebyshevDistance returns the Chebyshev distance between two data sets.
func ChebyshevDistance[N Number](dataPointX, dataPointY []N) N {
	if validateData(dataPointX, dataPointY) {
		return Empty[N]()
	}

	var tempDistance, distance N
	for i := 0; i < len(dataPointY); i++ {
		tempDistance = Abs(dataPointX[i] - dataPointY[i])
		if distance < tempDistance {
			distance = tempDistance
		}
	}

	return distance
}

// EuclideanDistance returns the Euclidean distance between two data sets.
func EuclideanDistance[N Number](dataPointX, dataPointY []N) float64 {
	if validateData(dataPointX, dataPointY) {
		return math.NaN()
	}

	var distance N
	for i := 0; i < len(dataPointX); i++ {
		distance += ((dataPointX[i] - dataPointY[i]) * (dataPointX[i] - dataPointY[i]))
	}

	return math.Sqrt(float64(distance))
}

// ManhattanDistance returns the Manhattan distance between two data sets.
func ManhattanDistance[N Number](dataPointX, dataPointY []N) N {
	if validateData(dataPointX, dataPointY) {
		return Empty[N]()
	}

	var distance N
	for i := 0; i < len(dataPointX); i++ {
		distance += Abs(dataPointX[i] - dataPointY[i])
	}

	return distance
}

// MinkowskiDistance returns the Minkowski distance between two data sets.
func MinkowskiDistance[N Number](dataPointX, dataPointY []N, lambda float64) float64 {
	if validateData(dataPointX, dataPointY) {
		return math.NaN()
	}

	var distance float64
	for i := 0; i < len(dataPointY); i++ {
		distance += Pow(float64(Abs(dataPointX[i]-dataPointY[i])), lambda)
	}

	distance = Pow(distance, 1/lambda)
	if math.IsInf(distance, 1) {
		return math.NaN()
	}

	return distance
}
