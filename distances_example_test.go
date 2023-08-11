// nolint: dupword
package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleChebyshevDistance() {
	fmt.Println(kstats.ChebyshevDistance([]int{2, 3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2}))
	fmt.Println(kstats.ChebyshevDistance([]int{3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2}))

	// Output:
	// 6
	// 0
}

func ExampleEuclideanDistance() {
	fmt.Println(kstats.EuclideanDistance([]int{2, 3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2}))
	fmt.Println(kstats.EuclideanDistance([]int{3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2}))

	// Output:
	// 10.583005244258363
	// NaN
}

func ExampleManhattanDistance() {
	fmt.Println(kstats.ManhattanDistance([]int{2, 3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2}))
	fmt.Println(kstats.ManhattanDistance([]int{3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2}))

	// Output:
	// 24
	// 0
}

func ExampleMinkowskiDistance() {
	fmt.Println(kstats.MinkowskiDistance([]int{2, 3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2}, 1.0))
	fmt.Println(kstats.MinkowskiDistance([]int{3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2}, 1.0))
	fmt.Println(kstats.MinkowskiDistance([]int{}, []int{8, 7, 6, 5, 4, 3, 2}, 1.0))
	fmt.Println(kstats.MinkowskiDistance([]int{2, 3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2}, 0))

	// Output:
	// 24
	// NaN
	// NaN
	// NaN
}
