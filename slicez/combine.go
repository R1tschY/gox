// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package slicez

func Chain[V any](iterables ...[]V) []V {
	return Flatten(iterables)
}

func Flatten[V any](iterables [][]V) []V {
	newLen := 0
	for _, val := range iterables {
		newLen += len(val)
		if newLen < 0 {
			panic("length overflow")
		}
	}

	res := make([]V, newLen)
	k := 0
	for _, vs := range iterables {
		for _, v := range vs {
			res[k] = v
			k++
		}
	}
	return res
}
