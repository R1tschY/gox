// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycle(t *testing.T) {
	t.Run("cycle", func(t *testing.T) {
		assert.Equal(t, []int{0, 1, 0, 1, 0}, slices.Collect(Limit(Cycle(Range(2)), 5)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak(t, Cycle(Range(2)), 0)
	})

	t.Run("late early break", func(t *testing.T) {
		checkEarlyBreak(t, Cycle(Range(2)), 7)
	})
}

func TestCycle2(t *testing.T) {
	t.Run("cycle", func(t *testing.T) {
		assert.Equal(t, "0:0 1:1 0:0 1:1 0:0", collectSeq2(Limit2(Cycle2(range2(2)), 5)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak2(t, Cycle2(range2(2)), 0)
	})

	t.Run("late early break", func(t *testing.T) {
		checkEarlyBreak2(t, Cycle2(range2(2)), 7)
	})
}
