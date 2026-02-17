// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import "github.com/R1tschy/iterz/internal"

// Repeat0Times returns a sequence of length count.
func Repeat0Times[I internal.Integer](count I) Seq0 {
	if count < 0 {
		panic("negative count")
	}

	return func(yield func() bool) {
		for i := I(0); i < count; i++ {
			if !yield() {
				break
			}
		}
	}
}

// Repeat returns a sequence of length count, each element equal to value.
func RepeatTimes[V any, I internal.Integer](value V, count I) Seq[V] {
	if count < 0 {
		panic("negative count")
	}

	return func(yield func(V) bool) {
		for i := I(0); i < count; i++ {
			if !yield(value) {
				break
			}
		}
	}
}

// Repeat returns a sequence of length count, each element equal to value.
func Repeat[V any](value V) Seq[V] {
	return func(yield func(V) bool) {
		for {
			if !yield(value) {
				break
			}
		}
	}
}

// Repeat2Times generates a sequence of key-value pairs repeated a specified number of times
// Panics if `count` is negative.
func Repeat2Times[K any, V any, I internal.Integer](key K, value V, count I) Seq2[K, V] {
	if count < 0 {
		panic("negative count")
	}

	return func(yield func(K, V) bool) {
		for i := I(0); i < count; i++ {
			if !yield(key, value) {
				break
			}
		}
	}
}
