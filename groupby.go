// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import "iter"

func GroupBy[K comparable, V any](it Seq[V], key func(V) K) Seq2[K, Seq[V]] {
	return func(yield func(K, Seq[V]) bool) {
		exhausted := false
		var currValue V
		var currKey K

		iterator, stop := iter.Pull(it)
		defer stop()

		_grouper := func(targetKey K) Seq[V] {
			return func(yield func(V) bool) {
				if !yield(currValue) {
					return
				}

				for {
					var ok bool
					currValue, ok = iterator()
					if !ok {
						exhausted = true
						return
					}
					currKey = key(currValue)
					if currKey != targetKey {
						return
					}
					if !yield(currValue) {
						return
					}
				}
			}
		}

		var ok bool
		currValue, ok = iterator()
		if !ok {
			return
		}
		currKey = key(currValue)

		for !exhausted {
			targetKey := currKey
			currGroup := _grouper(targetKey)
			if !yield(targetKey, currGroup) {
				return
			}
			if currKey == targetKey {
				for range currGroup {
				}
				// TODO: faster
				//for {
				//	val, ok := iterator()
				//	if !ok {
				//		exhausted = true
				//		return
				//	}
				//	currKey = key(val)
				//	if currKey != targetKey {
				//		return
				//	}
				//}
			}
		}
	}
}
