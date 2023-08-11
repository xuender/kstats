package kstats

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Abs[N Number](float N) N {
	return N(math.Abs(float64(float)))
}

func Pow[N Number](x, y N) N {
	return N(math.Pow(float64(x), float64(y)))
}

func Modf[F constraints.Float](float F) (F, F) {
	ret1, ret2 := math.Modf(float64(float))

	return F(ret1), F(ret2)
}

func Ceil[F constraints.Float](float F) F {
	return F(math.Ceil(float64(float)))
}

func Floor[F constraints.Float](float F) F {
	return F(math.Floor(float64(float)))
}

func Exp[N Number](number N) N {
	return N(math.Exp(float64(number)))
}

func Log[F constraints.Float](float F) F {
	return F(math.Log(float64(float)))
}

func Sqrt[F constraints.Float](float F) F {
	return F(math.Sqrt(float64(float)))
}
