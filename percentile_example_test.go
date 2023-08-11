package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExamplePercentile() {
	fmt.Println(kstats.Percentile(90, 63, 54, 61, 62, 66))
	fmt.Println(kstats.Percentile(90, 43))
	fmt.Println(kstats.Percentile(90, []int{}...))
	fmt.Println(kstats.Percentile(190, 3, 2))
	fmt.Println(kstats.Percentile(0.2, 3, 2))
	fmt.Println(kstats.Percentile(10, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9))

	// Output:
	// 64
	// 43
	// 0
	// 0
	// 0
	// 1
}

func ExamplePercentileNearestRank() {
	fmt.Println(kstats.PercentileNearestRank(0, 1, 2, 3))
	fmt.Println(kstats.PercentileNearestRank(30, 35, 20, 15, 40, 50))
	fmt.Println(kstats.PercentileNearestRank(190, 3, 2))
	fmt.Println(kstats.PercentileNearestRank(90, []int{}...))
	fmt.Println(kstats.PercentileNearestRank(100, 3, 2))

	// Output:
	// 1
	// 20
	// 0
	// 0
	// 3
}
