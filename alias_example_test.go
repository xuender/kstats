package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleVariance() {
	fmt.Println(kstats.Round(kstats.Variance(1.0, 2.0, 3.0), 1))
	fmt.Println(kstats.Variance([]float64{}...))

	// Output:
	// 0.7
	// 0
}

func ExamplePopulationVariance() {
	fmt.Println(kstats.Round(kstats.PopulationVariance(1.0, 2.0, 3.0), 1))

	// Output:
	// 0.7
}

func ExampleSampleVariance() {
	fmt.Println(kstats.SampleVariance(1.0, 2.0, 3.0))

	// Output:
	// 1
}

func ExampleMedianAbsoluteDeviation() {
	fmt.Println(kstats.MedianAbsoluteDeviation([]float64{1, 2, 3}...))
	fmt.Println(kstats.MedianAbsoluteDeviation([]float64{}...))

	// Output:
	// 1
	// 0
}

func ExampleStandardDeviation() {
	fmt.Println(kstats.Round(kstats.StandardDeviation([]float64{1, 2, 3}...), 2))
	fmt.Println(kstats.StandardDeviation([]float64{}...))

	// Output:
	// 0.82
	// 0
}

func ExamplePearson() {
	floats1 := []float64{1, 2, 3, 4, 5}
	floats2 := []float64{1, 2, 3, 5, 6}
	fmt.Println(kstats.Round(kstats.Pearson(floats1, floats2), 5))
	fmt.Println(kstats.Pearson(floats1, []float64{}))
	fmt.Println(kstats.Pearson([]float64{0}, []float64{0}))

	// Output:
	// 0.99124
	// 0
	// 0
}

func ExampleVarP() {
	fmt.Println(kstats.Round(kstats.VarP(1.0, 2.0, 3.0), 1))

	// Output:
	// 0.7
}

func ExampleVarS() {
	fmt.Println(kstats.VarS(1.0, 2.0, 3.0))

	// Output:
	// 1
}

func ExampleStdDevP() {
	fmt.Println(kstats.Round(kstats.StdDevP([]float64{1, 2, 3}...), 2))

	// Output:
	// 0.82
}

func ExampleStdDevS() {
	fmt.Println(kstats.StdDevS([]float64{1, 2, 3}...))

	// Output:
	// 1
}

func ExampleLinReg() {
	data := []kstats.Coordinate[int, int]{
		{1, 2},
		{2, 3},
		{3, 3},
	}

	fmt.Println(kstats.LinReg(data))

	// Output:
	// [{1 2.1666666666666665} {2 2.6666666666666665} {3 3.1666666666666665}]
}

func ExampleExpReg() {
	data := []kstats.Coordinate[int, float64]{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	fmt.Println(kstats.ExpReg(data))

	// Output:
	// [{1 2.5150181024736638} {2 3.032084111136781} {3 3.6554544271334493} {4 4.406984298281804} {5 5.313022222665875}]
}

func ExampleLogReg() {
	data := []kstats.Coordinate[int, float64]{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	fmt.Println(kstats.LogReg(data))

	// Output:
	// [{1 2.1520822363811702} {2 3.3305559222492214} {3 4.019918836568674} {4 4.509029608117273} {5 4.888413396683663}]
}
