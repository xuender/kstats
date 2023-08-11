package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExamplePow() {
	fmt.Println(kstats.Pow(2, 3))
	fmt.Println(kstats.Pow(4.0, 0.5))

	// Output:
	// 8
	// 2
}
