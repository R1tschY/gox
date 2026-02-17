// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

func SeqOf[V any](values ...V) Seq[V] {
	return func(yield func(V) bool) {
		for _, value := range values {
			if !yield(value) {
				break
			}
		}
	}
}
