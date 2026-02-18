// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import (
	"math"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	t.Run("from zero", func(t *testing.T) {
		assert.Equal(t, []int{0, 1, 2}, slices.Collect(Limit(Count(0, 1), 3)))
	})

	t.Run("from one", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3}, slices.Collect(Limit(Count(1, 1), 3)))
	})

	t.Run("big steps", func(t *testing.T) {
		assert.Equal(t, []int{0, 7, 14}, slices.Collect(Limit(Count(0, 7), 3)))
	})

	t.Run("negative steps", func(t *testing.T) {
		assert.Equal(t, []int{0, -1, -2}, slices.Collect(Limit(Count(0, -1), 3)))
	})

	t.Run("zero steps", func(t *testing.T) {
		assert.Equal(t, []int{0, 0, 0}, slices.Collect(Limit(Count(0, 0), 3)))
	})

	t.Run("pos overflow +1", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = slices.Collect(Limit(Count(math.MaxInt, 1), 2))
		})
	})

	t.Run("pos overflow +2", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = slices.Collect(Limit(Count(math.MaxInt, 2), 2))
		})
	})

	t.Run("pos overflow +MaxInt", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = slices.Collect(Limit(Count(math.MaxInt, math.MaxInt), 2))
		})
	})

	t.Run("neg overflow -1", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = slices.Collect(Limit(Count(math.MinInt, -1), 2))
		})
	})

	t.Run("neg overflow -2", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = slices.Collect(Limit(Count(math.MinInt, -2), 2))
		})
	})

	t.Run("neg overflow -MaxInt", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = slices.Collect(Limit(Count(math.MinInt, math.MinInt), 2))
		})
	})

	t.Run("early break +1", func(t *testing.T) {
		checkEarlyBreak(t, Count(0, 1), 0)
	})

	t.Run("early break -1", func(t *testing.T) {
		checkEarlyBreak(t, Count(0, -1), 0)
	})

	t.Run("early break 0", func(t *testing.T) {
		checkEarlyBreak(t, Count(0, 0), 0)
	})
}
