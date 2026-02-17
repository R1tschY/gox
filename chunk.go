// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

// Chunk splits a sequence into chunks of size n.
// It panics if n <= 0.
// When your input is a slice, consider using the faster [slices.Chunk] instead.
func Chunk[V any](it Seq[V], n int) Seq[[]V] {
	if n <= 0 {
		panic("n must be positive")
	}

	return func(yield func([]V) bool) {
		chunk := make([]V, 0, n)

		for v := range it {
			chunk = append(chunk, v)
			if len(chunk) == n {
				if !yield(chunk) {
					chunk = nil
					break
				}
				chunk = make([]V, 0, n)
			}
		}
		if len(chunk) > 0 {
			yield(chunk)
		}
	}
}
