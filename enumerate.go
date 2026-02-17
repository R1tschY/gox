// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import "github.com/R1tschy/iterz/internal"

func Enumerate[K internal.Integer, V any](it Seq[V], start K) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := start
		for value := range it {
			if !yield(i, value) {
				break
			}
			if i+1 < i {
				panic("integer overflow")
			}
			i++
		}
	}
}
