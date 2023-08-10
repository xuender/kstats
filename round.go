package stat

import (
	"golang.org/x/exp/constraints"
)

// Round returns a float to a specific decimal place or precision.
func Round[F constraints.Float](float F, places int) F {
	var sign F = 1
	if float < 0 {
		sign = -1
		float *= -1
	}

	var (
		rounded    F
		precision  = Pow(F(_ten), F(places))
		digit      = float * precision
		_, decimal = Modf(digit)
	)

	if decimal >= _half {
		rounded = Ceil(digit)
	} else {
		rounded = Floor(digit)
	}

	return rounded / precision * sign
}
