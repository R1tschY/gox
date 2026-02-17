// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package slicez

import "github.com/R1tschy/iterz/internal"

func Range[I internal.Integer](end I) []I {
	if end <= 0 {
		return nil
	}

	res := make([]I, end)
	for i := I(0); i < end; i++ {
		res[i] = i
	}
	return res
}

func RangeFrom[I internal.Integer](start I, end I) []I {
	if start >= end {
		return nil
	}

	res := make([]I, end-start)
	for i := start; i < end; i++ {
		res[i-start] = i
	}
	return res
}

func RangeBy[I internal.Integer](start I, end I, step I) []I {
	if step == 0 {
		panic("step must not be zero")
	}

	if step > 0 {
		if start >= end {
			return nil
		}

		res := make([]I, (end-start)/step)
		j := 0
		// "end-i > 0" instead of "i < end" to handle overflow properly
		for i := start; end-i > 0; i += step {
			res[j] = i
			j++
		}
		return res
	} else {
		if start <= end {
			return nil
		}

		res := make([]I, (start-end)/step)
		j := 0
		// "end-i < 0" instead of "i > end" to handle overflow properly
		for i := start; end-i < 0; i += step {
			res[j] = i
			j++
		}
		return res
	}
}
