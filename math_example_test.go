package stat_test

import (
	"fmt"

	"github.com/xuender/stat"
)

func ExampleMax() {
	fmt.Println(stat.Max(9.1, 1.0, 2.2, 30.2, 4.0))
	fmt.Println(stat.Max([]int{}...))

	// Output:
	// 30.2
	// 0
}
