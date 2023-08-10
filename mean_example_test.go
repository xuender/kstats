package stat_test

import (
	"fmt"

	"github.com/xuender/stat"
)

func ExampleMean() {
	fmt.Println(stat.Mean(1, 2, 3, 6))
	fmt.Println(stat.Mean([]int{}...))

	// Output:
	// 3
	// 0
}

func ExampleGeometricMean() {
	fmt.Println(stat.GeometricMean(2.0, 18.0))
	fmt.Println(stat.GeometricMean([]float32{}...))

	// Output:
	// 6
	// 0
}

func ExampleHarmonicMean() {
	fmt.Println(stat.Round(stat.HarmonicMean(1.0, 2.0, 3.0, 4.0, 5.0), 2))
	fmt.Println(stat.HarmonicMean([]float32{}...))
	fmt.Println(stat.Round(stat.HarmonicMean(0.0), 2))

	// Output:
	// 2.19
	// 0
	// 0
}
