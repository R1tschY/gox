// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

func Compress[V any](seq Seq[V], pred func(V) bool) Seq[V] {
	return func(yield func(V) bool) {
		for val := range seq {
			if pred(val) {
				if !yield(val) {
					break
				}
			}
		}
	}
}

func Compress2[K any, V any](seq Seq2[K, V], pred func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range seq {
			if pred(key, val) {
				if !yield(key, val) {
					break
				}
			}
		}
	}
}
