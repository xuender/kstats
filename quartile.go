package kstats

// Quartiles holds the three quartile points.
type Quartiles[N Number] struct {
	Q1 N
	Q2 N
	Q3 N
}

// Quartile returns the three quartile points from numbers.
func Quartile[N Number](numbers ...N) Quartiles[N] {
	length := len(numbers)
	if length == 0 {
		return Quartiles[N]{}
	}

	var (
		clone = SortedClone(numbers)
		right int
		left  int
	)

	if length%_two == 0 {
		right = length / _two
		left = length / _two
	} else {
		right = (length - 1) / _two
		left = right + 1
	}

	return Quartiles[N]{
		Q1: Median(clone[:right]...),
		Q2: Median(clone...),
		Q3: Median(clone[left:]...),
	}
}

// InterQuartileRange returns the range between Q1 and Q3.
func InterQuartileRange[N Number](numbers ...N) N {
	if len(numbers) == 0 {
		return Empty[N]()
	}

	qua := Quartile(numbers...)

	return qua.Q3 - qua.Q1
}

// Midhinge returns the average of the first and third quartiles.
func Midhinge[N Number](numbers ...N) N {
	if len(numbers) == 0 {
		return Empty[N]()
	}

	qua := Quartile(numbers...)

	return (qua.Q1 + qua.Q3) / _two
}

// Trimean returns the average of the median and the midhinge.
func Trimean[N Number](input ...N) N {
	if len(input) == 0 {
		return Empty[N]()
	}

	qua := Quartile(SortedClone(input)...)

	return (qua.Q1 + (qua.Q2 * _two) + qua.Q3) / _four
}
