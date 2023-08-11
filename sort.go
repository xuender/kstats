package kstats

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[O constraints.Ordered](ordereds []O) {
	sort.Slice(ordereds, func(i, j int) bool { return ordereds[i] < ordereds[j] })
}
