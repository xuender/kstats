package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleCovariance() {
	fmt.Println(kstats.Covariance([]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 5, 6}))
	fmt.Println(kstats.Covariance([]float64{2, 3, 4, 5}, []float64{1, 2, 3, 5, 6}))

	// Output:
	// 3.2499999999999996
	// 0
}

func ExampleCovariancePopulation() {
	fmt.Println(kstats.CovariancePopulation([]float64{1, 2, 3.5, 3.7, 8, 12}, []float64{0.5, 1, 2.1, 3.4, 3.4, 4}))
	fmt.Println(kstats.CovariancePopulation([]float64{2, 3.5, 3.7, 8, 12}, []float64{0.5, 1, 2.1, 3.4, 3.4, 4}))

	// Output:
	// 4.191666666666666
	// 0
}
