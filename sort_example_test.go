package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleSort() {
	ints := []int{4, 3, 1, 2}
	kstats.Sort(ints)
	fmt.Println(ints)

	// Output:
	// [1 2 3 4]
}

func ExampleSorted() {
	fmt.Println(kstats.Sorted([]int{1, 2, 3}))
	fmt.Println(kstats.Sorted([]int{1, 2, 3, 0}))

	// Output:
	// true
	// false
}
