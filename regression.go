package kstats

import (
	"math"
)

// Series is a container for a series of data.
type Series[N Number, F Number] []Coordinate[N, F]

// Coordinate holds the data in a series.
type Coordinate[N Number, F Number] struct {
	X N
	Y F
}

// LinearRegression returns the least squares linear regression on data series.
func LinearRegression[N, F Number](series Series[N, F]) Series[N, float64] {
	length := len(series)
	if length == 0 {
		return nil
	}

	var sum [5]float64

	for _, coo := range series {
		sum[0] += float64(coo.X)
		sum[1] += float64(coo.Y)
		sum[2] += float64(coo.X * coo.X)
		sum[3] += float64(coo.X) * float64(coo.Y)
		sum[4] += float64(coo.Y * coo.Y)
	}

	var (
		float       = float64(length)
		gradient    = (float*sum[3] - sum[0]*sum[1]) / (float*sum[2] - sum[0]*sum[0])
		intercept   = (sum[1] / float) - (gradient * sum[0] / float)
		regressions = make([]Coordinate[N, float64], length)
	)

	for i, coord := range series {
		regressions[i] = Coordinate[N, float64]{
			X: coord.X,
			Y: float64(coord.X)*gradient + intercept,
		}
	}

	return regressions
}

// ExponentialRegression returns an exponential regression on data series.
func ExponentialRegression[N, F Number](series Series[N, F]) Series[N, float64] {
	length := len(series)
	if length == 0 {
		return nil
	}

	var sum [6]float64

	for _, coord := range series {
		if coord.Y < 0 {
			return nil
		}

		sum[0] += float64(coord.X)
		sum[1] += float64(coord.Y)
		sum[2] += float64(coord.X*coord.X) * float64(coord.Y)
		sum[3] += float64(coord.Y) * Log(float64(coord.Y))
		sum[4] += float64(coord.X) * float64(coord.Y) * Log(float64(coord.Y))
		sum[5] += float64(coord.X) * float64(coord.Y)
	}

	var (
		denominator = (sum[1]*sum[2] - sum[5]*sum[5])
		left        = Pow(math.E, (sum[2]*sum[3]-sum[5]*sum[4])/denominator)
		right       = (sum[1]*sum[4] - sum[5]*sum[3]) / denominator
		regressions = make([]Coordinate[N, float64], length)
	)

	for i, coord := range series {
		regressions[i] = Coordinate[N, float64]{
			X: coord.X,
			Y: left * Exp(right*float64(coord.X)),
		}
	}

	return regressions
}

// LogarithmicRegression returns an logarithmic regression on data series.
func LogarithmicRegression[N, F Number](series Series[N, F]) Series[N, float64] {
	length := len(series)
	if length == 0 {
		return nil
	}

	var sum [4]float64

	for _, coord := range series {
		sum[0] += Log(float64(coord.X))
		sum[1] += float64(coord.Y) * Log(float64(coord.X))
		sum[2] += float64(coord.Y)
		sum[3] += Pow(Log(float64(coord.X)), _two)
	}

	var (
		float       = float64(length)
		right       = (float*sum[1] - sum[2]*sum[0]) / (float*sum[3] - sum[0]*sum[0])
		left        = (sum[2] - right*sum[0]) / float
		regressions = make([]Coordinate[N, float64], length)
	)

	for i, coord := range series {
		regressions[i] = Coordinate[N, float64]{
			X: coord.X,
			Y: left + right*Log(float64(coord.X)),
		}
	}

	return regressions
}
