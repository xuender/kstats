package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleMean() {
	fmt.Println(kstats.Mean(1, 2, 3, 6))
	fmt.Println(kstats.Mean([]int{}...))

	// Output:
	// 3
	// 0
}

func ExampleGeometricMean() {
	fmt.Println(kstats.GeometricMean(2.0, 18.0))
	fmt.Println(kstats.GeometricMean([]float32{}...))

	// Output:
	// 6
	// 0
}

func ExampleHarmonicMean() {
	fmt.Println(kstats.Round(kstats.HarmonicMean(1.0, 2.0, 3.0, 4.0, 5.0), 2))
	fmt.Println(kstats.HarmonicMean([]float32{}...))
	fmt.Println(kstats.Round(kstats.HarmonicMean(0.0), 2))

	// Output:
	// 2.19
	// 0
	// 0
}
