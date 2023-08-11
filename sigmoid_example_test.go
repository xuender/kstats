package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleSigmoid() {
	fmt.Println(kstats.Sigmoid(3.0, 1.0, 2.1))
	fmt.Println(kstats.Sigmoid([]float32{}...))

	// Output:
	// [0.9525741268224334 0.7310585786300049 0.8909031788043871]
	// []
}
