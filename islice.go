// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import "github.com/R1tschy/iterz/internal"

func Limit[V any, I internal.Integer](it Seq[V], n I) Seq[V] {
	return func(yield func(V) bool) {
		i := I(0)
		for val := range it {
			if i >= n {
				break
			}
			if !yield(val) {
				break
			}
			i += 1
		}
	}
}

func Limit2[K any, V any, I internal.Integer](it Seq2[K, V], n I) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := I(0)
		for key, val := range it {
			if i >= n {
				break
			}
			if !yield(key, val) {
				break
			}
			i += 1
		}
	}
}

func Slice[V any, I internal.Integer](it Seq[V], start I, end I) Seq[V] {
	return func(yield func(V) bool) {
		i := I(0)
		for val := range it {
			if i >= end {
				break
			}
			if i >= start {
				if !yield(val) {
					break
				}
			}
			i += 1
		}
	}
}

func Slice2[K any, V any, I internal.Integer](it Seq2[K, V], start I, end I) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := I(0)
		for key, val := range it {
			if i < start {
				continue
			}
			if i >= end || !yield(key, val) {
				break
			}
			i += 1
		}
	}
}
