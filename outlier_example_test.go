package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleQuartileOutliers() {
	fmt.Println(kstats.QuartileOutliers[float64](-1000, 1, 3, 4, 4, 6, 6, 6, 6, 7, 8, 15, 18, 100))
	fmt.Println(kstats.QuartileOutliers([]float64{}...))

	// Output:
	// {[15 18] [-1000 100]}
	// {[] []}
}
