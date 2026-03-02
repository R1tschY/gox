// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

func Chain[V any](iterables ...Seq[V]) Seq[V] {
	return func(yield func(V) bool) {
		for _, iterable := range iterables {
			for val := range iterable {
				if !yield(val) {
					return
				}
			}
		}
	}
}

func ChainFrom[V any](iterables Seq[Seq[V]]) Seq[V] {
	return func(yield func(V) bool) {
		for iterable := range iterables {
			for val := range iterable {
				if !yield(val) {
					return
				}
			}
		}
	}
}

func Chain2[K any, V any](iterables ...Seq2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, iterable := range iterables {
			for key, val := range iterable {
				if !yield(key, val) {
					return
				}
			}
		}
	}
}

func ChainFrom2[K any, V any](iterables Seq[Seq2[K, V]]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for iterable := range iterables {
			for key, val := range iterable {
				if !yield(key, val) {
					return
				}
			}
		}
	}
}
