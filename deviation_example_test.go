package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleStandardDeviationSample() {
	fmt.Println(kstats.StandardDeviationSample([]float64{1, 2, 3}...))
	fmt.Println(kstats.StandardDeviationSample([]float64{}...))

	// Output:
	// 1
	// 0
}
