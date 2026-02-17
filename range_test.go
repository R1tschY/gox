// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import (
	"math"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeFrom(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, []int(nil), slices.Collect(RangeFrom(0, 0)))
	})

	t.Run("0..3", func(t *testing.T) {
		assert.Equal(t, []int{0, 1, 2}, slices.Collect(RangeFrom(0, 3)))
	})

	t.Run("max-1..max", func(t *testing.T) {
		assert.Equal(t, []int{math.MaxInt - 1}, slices.Collect(RangeFrom(math.MaxInt-1, math.MaxInt)))
	})

	t.Run("min..min+1", func(t *testing.T) {
		assert.Equal(t, []int{math.MinInt}, slices.Collect(RangeFrom(math.MinInt, math.MinInt+1)))
	})

	t.Run("1..-1", func(t *testing.T) {
		assert.Equal(t, []int(nil), slices.Collect(RangeFrom(1, -1)))
	})

	t.Run("max..min", func(t *testing.T) {
		assert.Equal(t, []int(nil), slices.Collect(RangeFrom(math.MaxInt, math.MinInt)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak(t, RangeFrom(0, 3), 0)
	})
}

func TestRangeBy(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, []int(nil), slices.Collect(RangeBy(0, 0, 1)))
	})

	t.Run("0..3", func(t *testing.T) {
		assert.Equal(t, []int{0, 1, 2}, slices.Collect(RangeBy(0, 3, 1)))
	})

	t.Run("0..3 by 2", func(t *testing.T) {
		assert.Equal(t, []int{0, 2}, slices.Collect(RangeBy(0, 3, 2)))
	})

	t.Run("3..0 by -1", func(t *testing.T) {
		assert.Equal(t, []int{3, 2, 1}, slices.Collect(RangeBy(3, 0, -1)))
	})

	t.Run("3..0 by -2", func(t *testing.T) {
		assert.Equal(t, []int{3, 1}, slices.Collect(RangeBy(3, 0, -2)))
	})

	t.Run("0..max", func(t *testing.T) {
		assert.Equal(t, []int{math.MaxInt - 1}, slices.Collect(RangeBy(math.MaxInt-1, math.MaxInt, 1)))
	})

	t.Run("min..min+1", func(t *testing.T) {
		assert.Equal(t, []int{math.MinInt}, slices.Collect(RangeBy(math.MinInt, math.MinInt+1, 1)))
	})

	t.Run("max-1..max by max", func(t *testing.T) {
		assert.Equal(t, []int{math.MaxInt - 1}, slices.Collect(RangeBy(math.MaxInt-1, math.MaxInt, math.MaxInt)))
	})

	t.Run("1..-1", func(t *testing.T) {
		assert.Equal(t, []int(nil), slices.Collect(RangeBy(1, -1, 1)))
	})

	t.Run("max..min", func(t *testing.T) {
		assert.Equal(t, []int(nil), slices.Collect(RangeBy(math.MaxInt, math.MinInt, 1)))
	})

	t.Run("min..max by -1", func(t *testing.T) {
		assert.Equal(t, []int(nil), slices.Collect(RangeBy(math.MinInt, math.MaxInt, -1)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak(t, RangeBy(0, 3, 1), 0)
	})
}
