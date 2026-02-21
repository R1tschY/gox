// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"slices"

	"github.com/R1tschy/gox/slicex"
)

func mapIndices[V any](indices []int, pool []V) []V {
	res := make([]V, len(indices))
	for i, idx := range indices {
		res[i] = pool[idx]
	}
	return res
}

func Combinations[V any](pool []V, r int) Seq[[]V] {
	return func(yield func([]V) bool) {
		n := len(pool)
		if r > n {
			return
		}

		indices := slicex.Range(r)

		if !yield(mapIndices(indices, pool)) {
			return
		}

		for {
			var i int
			for i = r - 1; i >= 0; i-- {
				if indices[i] != i+n-r {
					break
				}
			}
			if i < 0 {
				return
			}

			indices[i] += 1
			for j := i + 1; j < r; j++ {
				indices[j] = indices[j-1] + 1
			}

			if !yield(mapIndices(indices, pool)) {
				return
			}
		}
	}
}

func CombinationsWithReplacement[V any](pool []V, r int) Seq[[]V] {
	return func(yield func([]V) bool) {
		n := len(pool)
		if n == 0 && r != 0 {
			return
		}

		indices := make([]int, r)

		if !yield(mapIndices(indices, pool)) {
			return
		}

		for {
			var i int
			for i = r - 1; i >= 0; i-- {
				if indices[i] != n-1 {
					break
				}
			}
			if i < 0 {
				return
			}

			v := indices[i] + 1
			for j := i; j < r; j++ {
				indices[j] = v
			}

			if !yield(mapIndices(indices, pool)) {
				return
			}
		}
	}
}

func Permutations[V any](pool []V, r int) Seq[[]V] {
	return func(yield func([]V) bool) {
		n := len(pool)
		if r > n {
			return
		}

		indices := slicex.Range(r)
		cycles := slicex.RangeBy(n, n-r, -1)
		if !yield(mapIndices(indices, pool)) {
			return
		}

		for n != 0 {
			var i int
			for i = r - 1; i >= 0; i-- {
				cycles[i] -= 1
				if cycles[i] == 0 {
					last := indices[i]
					// TODO: try to use copy and benchmark: copy(indices[i+1:r-1], indices[i:r-2])
					for j := i; j < r-1; j++ {
						indices[j] = indices[j+1]
					}
					indices[r-1] = last
					cycles[i] = n - i
				} else {
					j := cycles[i]
					indices[i], indices[r-j] = indices[r-j], indices[i]
					if !yield(mapIndices(indices, pool)) {
						return
					}
					break
				}
			}
			if i < 0 {
				return
			}
		}
	}
}

func Product[V any](pools [][]V) Seq[[]V] {
	return func(yield func([]V) bool) {
		indices := make([]int, len(pools))

		res := make([]V, len(pools))
		for i, pool := range pools {
			res[i] = pool[0]
		}
		if !yield(res) {
			return
		}

		for {
			res = slices.Clone(res)

			var i int
			for i = len(pools) - 1; i >= 0; i-- {
				pool := pools[i]
				indices[i] += 1
				if indices[i] == len(pool) {
					indices[i] = 0
					res[i] = pool[0]
				} else {
					res[i] = pool[indices[i]]
					break
				}
			}
			if i < 0 {
				return
			}

			if !yield(res) {
				break
			}
		}
	}
}
