package stat

import "golang.org/x/exp/constraints"

// Max returns the highest elem in ordereds.
func Max[O constraints.Ordered](ordereds ...O) O {
	var (
		max    O
		length = len(ordereds)
	)

	if length == 0 {
		return max
	}

	max = ordereds[0]

	for i := 1; i < length; i++ {
		if ordereds[i] > max {
			max = ordereds[i]
		}
	}

	return max
}
