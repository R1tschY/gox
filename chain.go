// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

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
