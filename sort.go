package kstats

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[O constraints.Ordered](ordereds []O) {
	sort.Sort(Ordered[O](ordereds))
}

func Sorted[O constraints.Ordered](ordereds []O) bool {
	return sort.IsSorted(Ordered[O](ordereds))
}

func SortedClone[O constraints.Ordered](ordereds []O) []O {
	clone := Clone(ordereds)

	Sort(clone)

	return clone
}

func SortedCloneDif[O constraints.Ordered](ordereds []O) []O {
	if Sorted(ordereds) {
		return ordereds
	}

	return SortedClone(ordereds)
}
