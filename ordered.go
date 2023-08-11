package kstats

import "golang.org/x/exp/constraints"

type Ordered[O constraints.Ordered] []O

func (o Ordered[O]) Len() int           { return len(o) }
func (o Ordered[O]) Less(i, j int) bool { return o[i] < o[j] }
func (o Ordered[O]) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
