// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

// Len0 returns the length of the [Seq0] sequence.
// Panics if the sequence contains more than 2^31-1 elements.
func Len0(iter Seq0) int {
	acc := 0
	for range iter {
		acc += 1
		if acc < 0 {
			panic("integer overflow")
		}
	}
	return acc
}

// Len returns the length of the [Seq] sequence.
// Panics if the sequence contains more than 2^31-1 elements.
func Len[V any](iter Seq[V]) int {
	return Reduce(iter, 0, func(acc int, _ V) int {
		acc += 1
		if acc < 0 {
			panic("integer overflow")
		}
		return acc
	})
}

// Len2 returns the length of the [Seq2] sequence.
// Panics if the sequence contains more than 2^31-1 elements.
func Len2[K any, V any](iter Seq2[K, V]) int {
	return Reduce2(iter, 0, func(acc int, _ K, _ V) int {
		acc += 1
		if acc < 0 {
			panic("integer overflow")
		}
		return acc
	})
}
