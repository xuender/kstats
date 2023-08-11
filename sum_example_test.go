package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleSum() {
	fmt.Println(kstats.Sum(1, 2, 3, 4))
	fmt.Println(kstats.Sum("a", "b", "c"))

	// Output:
	// 10
	// abc
}

func ExampleCumulativeSum() {
	fmt.Println(kstats.CumulativeSum(1.0, 2.1, 3.2, 4.823, 4.1, 5.8))
	fmt.Println(kstats.CumulativeSum([]int{}...))
	fmt.Println(kstats.CumulativeSum("a", "b", "c"))

	// Output:
	// [1 3.1 6.300000000000001 11.123000000000001 15.223 21.023]
	// []
	// [a ab abc]
}
