package kstats

import (
	"math/rand"
	"sort"
)

// Samples returns sample of elems.
func Sample[T any](elems ...T) T {
	length := len(elems)
	switch length {
	case 0:
		return Empty[T]()
	case 1:
		return elems[0]
	default:
		// nolint: gosec
		return elems[rand.Intn(length)]
	}
}

// Samples returns sample from elems with replacement or without.
func Samples[T any](takenum int, replacement bool, elems ...T) []T {
	length := len(elems)
	if length == 0 {
		return nil
	}

	if replacement {
		result := make([]T, takenum)
		for index := 0; index < takenum; index++ {
			// nolint: gosec
			result[index] = elems[rand.Intn(length)]
		}

		return result
	}

	if takenum > length {
		takenum = length
	}

	var (
		perm   = rand.Perm(length)[:takenum]
		result = make([]T, takenum)
	)

	for index, idx := range perm {
		result[index] = elems[idx]
	}

	return result
}

// StableSamples returns samples keeps the order of original elems.
func StableSamples[T any](takenum int, elems ...T) []T {
	length := len(elems)
	if length == 0 {
		return nil
	}

	if takenum > length {
		takenum = length
	}

	var (
		perm   = rand.Perm(length)[:takenum]
		result = make([]T, takenum)
	)

	sort.Ints(perm)

	for index, idx := range perm {
		result[index] = elems[idx]
	}

	return result
}
