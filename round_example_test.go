package stat_test

import (
	"fmt"

	"github.com/xuender/stat"
)

func ExampleRound() {
	fmt.Println(stat.Round(3.1415926, 2))
	fmt.Println(stat.Round(-3.005926, 2))

	// Output:
	// 3.14
	// -3.01
}
