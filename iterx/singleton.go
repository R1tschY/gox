// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

// Singleton creates a iter.Seq iterator containing a single value.
func Singleton[V any](value V) Seq[V] {
	return func(yield func(V) bool) {
		yield(value)
	}
}

// Singleton2 creates a iter.Seq2 iterator containing a single value.
func Singleton2[K any, V any](key K, value V) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		yield(key, value)
	}
}
