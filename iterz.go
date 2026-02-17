// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

/*
Package iterz provides the standard library missing functions for working with [iter.Seq] and [iter.Seq2].

Inspired by [Python's itertools].

[Python's itertools]: https://docs.python.org/3/library/itertools.html
*/
package iterz

func Infinite() Seq0 {
	return func(yield func() bool) {
		for {
			if !yield() {
				return
			}
		}
	}
}

type pair[K, V any] struct {
	Key   K
	Value V
}

func newPair[K, V any](key K, value V) pair[K, V] {
	return pair[K, V]{Key: key, Value: value}
}
