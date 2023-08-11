package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleProbGeom() {
	fmt.Println(kstats.ProbGeom(1, 2, 0.5))
	fmt.Println(kstats.ProbGeom(2, 1, 0.5))

	// Output:
	// 0.25
	// 0
}

func ExampleExpGeom() {
	fmt.Println(kstats.ExpGeom(0.5))
	fmt.Println(kstats.ExpGeom(1.5))

	// Output:
	// 2
	// 0
}

func ExampleVarGeom() {
	fmt.Println(kstats.VarGeom(0.5))
	fmt.Println(kstats.VarGeom(11.0))

	// Output:
	// 2
	// 0
}
