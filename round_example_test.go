package kstats_test

import (
	"fmt"

	"github.com/xuender/kstats"
)

func ExampleRound() {
	fmt.Println(kstats.Round(3.1415926, 2))
	fmt.Println(kstats.Round(-3.005926, 2))

	// Output:
	// 3.14
	// -3.01
}
