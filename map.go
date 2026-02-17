// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

// Map applies a function to each element of the sequence.
func Map[V any, U any](seq Seq[V], f func(value V) U) Seq[U] {
	return func(yield func(U) bool) {
		for val := range seq {
			if !yield(f(val)) {
				break
			}
		}
	}
}

// Map2 applies a function to each pair element of the sequence, returning a key-value sequence.
func Map2[K any, V any, J any, U any](seq Seq2[K, V], f func(key K, value V) (J, U)) Seq2[J, U] {
	return func(yield func(J, U) bool) {
		for key, value := range seq {
			if !yield(f(key, value)) {
				break
			}
		}
	}
}

// Map2To1 applies a function to each pair element of the sequence, returning a sequence.
func Map2To1[K any, V any, U any](seq Seq2[K, V], f func(key K, value V) U) Seq[U] {
	return func(yield func(U) bool) {
		for key, value := range seq {
			if !yield(f(key, value)) {
				break
			}
		}
	}
}

// MapTo2 applies a function to each element of the sequence, returning a key-value sequence.
func MapTo2[T any, K any, V any](seq Seq[T], f func(value T) (K, V)) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for val := range seq {
			if !yield(f(val)) {
				break
			}
		}
	}
}
