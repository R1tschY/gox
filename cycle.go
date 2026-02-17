// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

// Cycle returns a sequence that repeats the elements of the given sequence indefinitely.
// It copies the elements of the sequence to an in-memory buffer in the first iteration and returns copies of them in
// later iterations.
func Cycle[T any](it Seq[T]) Seq[T] {
	return func(yield func(T) bool) {
		buf := make([]T, 0, 16)
		for value := range it {
			if !yield(value) {
				break
			}
			buf = append(buf, value)
		}

		if len(buf) > 0 {
		outer:
			for {
				for _, value := range buf {
					if !yield(value) {
						break outer
					}
				}
			}
		}
	}
}

// Cycle2 returns a key-value sequence that repeats the elements of the given key-value sequence indefinitely.
// It copies the elements of the sequence to an in-memory buffer in the first iteration and returns copies of them in
// later iterations.
func Cycle2[K any, V any](it Seq2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		buf := make([]pair[K, V], 0, 16)
		for key, value := range it {
			if !yield(key, value) {
				break
			}
			buf = append(buf, newPair(key, value))
		}

		if len(buf) > 0 {
		outer:
			for {
				for _, entry := range buf {
					if !yield(entry.Key, entry.Value) {
						break outer
					}
				}
			}
		}
	}
}
