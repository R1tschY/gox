// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"github.com/R1tschy/gox/typex"
)

// Count returns a sequence of integers starting from from and incremented by step.
// It panics if the sequence overflows an integer.
func Count[I typex.Integer](from I, step I) Seq[I] {
	if step > 0 {
		return func(yield func(I) bool) {
			i := from
			for {
				if !yield(i) {
					return
				}

				i += step
				if i <= from {
					panic("integer overflow")
				}
			}
		}
	} else if step < 0 {
		return func(yield func(I) bool) {
			i := from
			for {
				if !yield(i) {
					return
				}

				i += step
				if i >= from {
					panic("integer overflow")
				}
			}
		}
	}

	return func(yield func(I) bool) {
		for {
			if !yield(from) {
				return
			}
		}
	}
}
