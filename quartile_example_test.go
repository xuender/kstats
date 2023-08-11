package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleQuartile() {
	fmt.Println(kstats.Quartile(6, 7, 15, 36, 39, 40, 41, 42, 43, 47, 49))
	fmt.Println(kstats.Quartile[float64](7, 15, 36, 39, 40, 41))
	fmt.Println(kstats.Quartile([]int{}...))

	// Output:
	// {15 40 43}
	// {15 37.5 40}
	// {0 0 0}
}

func ExampleInterQuartileRange() {
	fmt.Println(kstats.InterQuartileRange(102, 104, 105, 107, 108, 109, 110, 112, 115, 116, 118))
	fmt.Println(kstats.InterQuartileRange([]int{}...))

	// Output:
	// 10
	// 0
}

func ExampleMidhinge() {
	fmt.Println(kstats.Midhinge[float32](1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13))
	fmt.Println(kstats.Midhinge([]int{}...))

	// Output:
	// 7.5
	// 0
}

func ExampleTrimean() {
	fmt.Println(kstats.Trimean[float64](1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13))
	fmt.Println(kstats.Trimean([]int{}...))

	// Output:
	// 7.25
	// 0
}
