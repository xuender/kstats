package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleMode() {
	fmt.Println(kstats.Mode(2.0, 2.0, 2.0, 2.0))
	fmt.Println(kstats.Mode([]float64{5, 3, 4, 2, 1}...))
	fmt.Println(kstats.Mode([]float64{-50, -46.325, -46.325, -.87, 1, 2.1122, 3.20, 5, 15, 15, 15.0001}...))

	fmt.Println(kstats.Mode([]float64{}...))
	fmt.Println(kstats.Mode([]float64{2.0}...))

	// Output:
	// [2]
	// []
	// [-46.325 15]
	// []
	// [2]
}
