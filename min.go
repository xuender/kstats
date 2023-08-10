package stat

import "golang.org/x/exp/constraints"

// Min returns the lowest elem in ordereds.
func Min[O constraints.Ordered](ordereds ...O) O {
	var (
		min    O
		length = len(ordereds)
	)

	if length == 0 {
		return min
	}

	min = ordereds[0]

	for i := 1; i < length; i++ {
		if ordereds[i] < min {
			min = ordereds[i]
		}
	}

	return min
}
