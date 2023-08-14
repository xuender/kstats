package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleDescribe() {
	describe, err := kstats.Describe([]float64{1.0, 2.0, 3.0}, true, []float64{25.0, 50.0, 75.0})

	fmt.Println(err)
	fmt.Println(describe.Max)
	fmt.Println(describe.String())

	_, err = kstats.Describe([]float64{}, false, nil)
	fmt.Println(err)

	// Output:
	// <nil>
	// 3
	// count	3
	// count	3
	// mean	2.00
	// std	0.82
	// max	3.00
	// min	1.00
	// 25.00%	0.00
	// 50.00%	1.50
	// 75.00%	2.50
	// NaN OK	true
	// input must not be empty
}
