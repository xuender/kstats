package stat

import "golang.org/x/exp/constraints"

// Sum returns the sum of ordereds.
func Sum[O constraints.Ordered](ordereds ...O) O {
	var sum O

	for _, o := range ordereds {
		sum += o
	}

	return sum
}
