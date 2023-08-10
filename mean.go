package stat

import "golang.org/x/exp/constraints"

// Mean returns the average of numbers.
func Mean[N Number](numbers ...N) N {
	var (
		zero   N
		length = len(numbers)
	)

	if length == 0 {
		return zero
	}

	return Sum[N](numbers...) / N(length)
}

// GeometricMean returns the geometric mean for floats.
func GeometricMean[F constraints.Float](floats ...F) F {
	var (
		zero   F
		length = len(floats)
	)

	if length == 0 {
		return zero
	}

	var sum F

	for index, num := range floats {
		if index == 0 {
			sum = num

			continue
		}

		sum *= num
	}

	return Pow(sum, 1/F(length))
}

// HarmonicMean return the harmonic mean for floats.
func HarmonicMean[F constraints.Float](floats ...F) F {
	var (
		zero   F
		length = len(floats)
	)

	if length == 0 {
		return zero
	}

	var sum F

	for _, num := range floats {
		if num <= 0 {
			return zero
		}

		sum += (1 / num)
	}

	return F(length) / sum
}
