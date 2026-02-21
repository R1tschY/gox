// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

func Pairwise[V any](seq Seq[V]) Seq2[V, V] {
	return func(yield func(V, V) bool) {
		seen := false
		var prev V

		for val := range seq {
			if seen {
				if !yield(prev, val) {
					break
				}
			} else {
				seen = true
			}
			prev = val
		}
	}
}
