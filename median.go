package kstats

// Median return the median elem of ordereds.
func Median[O Number](ordereds ...O) O {
	length := len(ordereds)
	switch length {
	case 0:
		return Empty[O]()
	case 1:
		return ordereds[0]
	default:
		clone := Clone(ordereds)
		Sort(clone)

		if length%2 == 0 {
			return Mean(clone[length/2-1 : length/2+1]...)
		}

		return clone[length/2]
	}
}
