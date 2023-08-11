package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleMedian() {
	fmt.Println(kstats.Median(5, 3, 4, 2, 1))
	fmt.Println(kstats.Median(6.0, 3.0, 2.0, 4.0, 5.0, 1.0))
	fmt.Println(kstats.Median(1))
	fmt.Println(kstats.Median([]int{}...))

	// Output:
	// 3
	// 3.5
	// 1
	// 0
}
