package stat_test

import (
	"fmt"

	"github.com/xuender/stat"
)

func ExampleMin() {
	fmt.Println(stat.Min(9, 1, 2, 3, 4))
	fmt.Println(stat.Min([]int{}...))

	// Output:
	// 1
	// 0
}
