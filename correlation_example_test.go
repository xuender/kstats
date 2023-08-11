package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleAutoCorrelation() {
	fmt.Println(kstats.AutoCorrelation(1, []float64{1, 2, 3, 4, 5}...))
	fmt.Println(kstats.AutoCorrelation(1, []float64{}...))

	// Output:
	// 0.4
	// 0
}
