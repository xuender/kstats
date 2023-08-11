package kstats

import "golang.org/x/exp/constraints"

// Sum returns the sum of ordereds.
func Sum[O constraints.Ordered](ordereds ...O) O {
	var sum O

	for _, o := range ordereds {
		sum += o
	}

	return sum
}

// CumulativeSum return the cumulative sum of ordereds.
func CumulativeSum[O constraints.Ordered](ordereds ...O) []O {
	length := len(ordereds)
	if length == 0 {
		return nil
	}

	cumSum := make([]O, length)

	for index, ordered := range ordereds {
		if index == 0 {
			cumSum[index] = ordered
		} else {
			cumSum[index] = cumSum[index-1] + ordered
		}
	}

	return cumSum
}
