package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleLinearRegression() {
	data := []kstats.Coordinate[float64]{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
	}

	fmt.Println(kstats.LinearRegression(data))
	fmt.Println(kstats.LinearRegression([]kstats.Coordinate[float64]{}))

	// Output:
	// [{1 2.400000000000001} {2 3.1} {3 3.7999999999999994}]
	// []
}

func ExampleLinearRegression_int() {
	data := []kstats.Coordinate[int]{
		{1, 2},
		{2, 3},
		{3, 3},
	}

	fmt.Println(kstats.LinearRegression(data))

	// Output:
	// [{1 2.1666666666666665} {2 2.6666666666666665} {3 3.1666666666666665}]
}

func ExampleExponentialRegression() {
	data := []kstats.Coordinate[float64]{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	fmt.Println(kstats.ExponentialRegression(data))
	fmt.Println(kstats.ExponentialRegression([]kstats.Coordinate[float64]{}))
	fmt.Println(kstats.ExponentialRegression([]kstats.Coordinate[float64]{{1, -1}}))

	// Output:
	// [{1 2.5150181024736638} {2 3.032084111136781} {3 3.6554544271334493} {4 4.406984298281804} {5 5.313022222665875}]
	// []
	// []
}

func ExampleLogarithmicRegression() {
	data := []kstats.Coordinate[float64]{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	fmt.Println(kstats.LogarithmicRegression(data))
	fmt.Println(kstats.LogarithmicRegression([]kstats.Coordinate[float64]{}))

	// Output:
	// [{1 2.1520822363811702} {2 3.3305559222492214} {3 4.019918836568674} {4 4.509029608117273} {5 4.888413396683663}]
	// []
}
