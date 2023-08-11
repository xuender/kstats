package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleEntropy() {
	fmt.Println(kstats.Entropy(0.0, 1.1, 2.2, 3.3))
	// Output:
	// 1.0114042647073518
}

func ExampleNormalize() {
	list := []float32{2, 2, 4}
	kstats.Normalize(list)
	fmt.Println(list)

	// Output:
	// [0.25 0.25 0.5]
}

func ExampleNormalize_inf() {
	list := []float32{2, -2}
	kstats.Normalize(list)
	fmt.Println(list)

	// Output:
	// [+Inf -Inf]
}
