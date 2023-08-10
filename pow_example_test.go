package stat_test

import (
	"fmt"

	"github.com/xuender/stat"
)

func ExamplePow() {
	fmt.Println(stat.Pow(2, 3))
	fmt.Println(stat.Pow(4.0, 0.5))

	// Output:
	// 8
	// 2
}
