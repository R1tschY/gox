// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import "github.com/R1tschy/iterz/internal"

func Range[I internal.Integer](end I) Seq[I] {
	return RangeFrom[I](0, end)
}

func RangeFrom[I internal.Integer](start I, end I) Seq[I] {
	return func(yield func(I) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				break
			}
		}
	}
}

func RangeBy[I internal.Integer](start I, end I, step I) Seq[I] {
	if step == 0 {
		panic("step must not be zero")
	}

	return func(yield func(I) bool) {
		if step > 0 {
			if start >= end {
				return
			}

			// "end-i > 0" instead of "i < end" to handle overflow properly
			for i := start; end-i > 0; i += step {
				if !yield(i) {
					break
				}
			}
		} else {
			if start <= end {
				return
			}

			// "end-i < 0" instead of "i > end" to handle overflow properly
			for i := start; end-i < 0; i += step {
				if !yield(i) {
					break
				}
			}
		}
	}
}
