package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleMax() {
	fmt.Println(kstats.Max(9.1, 1.0, 2.2, 30.2, 4.0))
	fmt.Println(kstats.Max([]int{}...))

	// Output:
	// 30.2
	// 0
}
