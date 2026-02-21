// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

func Accumulate[V any](it Seq[V], f func(V, V) V) Seq[V] {
	return func(yield func(V) bool) {
		first := true
		var acc V
		for val := range it {
			if first {
				acc = val
				first = false

				if !yield(acc) {
					break
				}
			} else {
				acc = f(acc, val)
				if !yield(acc) {
					break
				}
			}
		}
	}
}
