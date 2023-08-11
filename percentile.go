package kstats

import (
	"math"
)

// Percentile returns the relative standing in a slice of numbers.
func Percentile[N Number](percent float64, numbers ...N) N {
	length := len(numbers)
	if length == 0 {
		return Empty[N]()
	}

	if length == 1 {
		return numbers[0]
	}

	if percent <= 0 || percent > _hundred {
		return Empty[N]()
	}

	var (
		clone      = SortedClone(numbers)
		index      = (percent / _hundred) * float64(length)
		percentile N
	)

	switch {
	case index == float64(int64(index)):
		percentile = clone[int(index)-1]
	case index > 1:
		i := int(index)
		percentile = Mean(clone[i-1], clone[i])
	default:
		return Empty[N]()
	}

	return percentile
}

// PercentileNearestRank returns the relative standing in a slice of floats using the Nearest Rank method.
func PercentileNearestRank[N Number](percent float64, numbers ...N) N {
	length := len(numbers)
	if length == 0 {
		return Empty[N]()
	}

	if percent < 0 || percent > _hundred {
		return Empty[N]()
	}

	clone := SortedClone(numbers)

	if percent == float64(_hundred) {
		return clone[length-1]
	}

	index := int(math.Ceil(float64(length) * percent / _hundred))
	if index == 0 {
		return clone[0]
	}

	return clone[index-1]
}
