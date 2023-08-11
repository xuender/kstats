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
