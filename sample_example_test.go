package kstats_test

import (
	"fmt"
	"math/rand"

	"github.com/xuender/kstats"
)

func ExampleSample() {
	rand.Seed(0)
	fmt.Println(kstats.Sample(0, 1, 2, 3))
	fmt.Println(kstats.Sample(1))
	fmt.Println(kstats.Sample([]int{}...))

	// Output:
	// 2
	// 1
	// 0
}

func ExampleSamples() {
	rand.Seed(0)
	fmt.Println(kstats.Samples(3, true, 0, 1, 2, 3))
	fmt.Println(kstats.Samples(3, false, 0, 1, 2, 3))
	fmt.Println(kstats.Samples(30, false, 0, 1, 2, 3))
	fmt.Println(kstats.Samples(30, false, []int{}...))

	// Output:
	// [2 2 1]
	// [0 2 1]
	// [3 0 1 2]
	// []
}

func ExampleStableSamples() {
	rand.Seed(0)
	fmt.Println(kstats.StableSamples(3, 0, 1, 2, 3))
	fmt.Println(kstats.StableSamples(30, 0, 1, 2, 3))
	fmt.Println(kstats.StableSamples(30, []int{}...))

	// Output:
	// [1 2 3]
	// [0 1 2 3]
	// []
}
