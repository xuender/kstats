package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleMin() {
	fmt.Println(kstats.Min(9, 1, 2, 3, 4))
	fmt.Println(kstats.Min([]int{}...))

	// Output:
	// 1
	// 0
}
