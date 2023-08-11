package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleSoftMax() {
	fmt.Println(kstats.SoftMax(3.0, 1.0, 0.2))
	fmt.Println(kstats.SoftMax([]float32{}...))

	// Output:
	// [0.8360188027814407 0.11314284146556013 0.05083835575299916]
	// []
}
