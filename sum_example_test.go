package stat_test

import (
	"fmt"

	"github.com/xuender/stat"
)

func ExampleSum() {
	fmt.Println(stat.Sum(1, 2, 3, 4))
	fmt.Println(stat.Sum("a", "b", "c"))

	// Output:
	// 10
	// abc
}
