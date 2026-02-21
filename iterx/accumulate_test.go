// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"slices"
	"testing"

	"github.com/R1tschy/gox/ops"
	"github.com/stretchr/testify/assert"
)

func TestAccumulate(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len(Accumulate(Empty[int](), ops.Add)))
	})

	t.Run("one", func(t *testing.T) {
		assert.Equal(t, []int{0}, slices.Collect(Accumulate(Range(1), ops.Add)))
	})

	t.Run("add", func(t *testing.T) {
		assert.Equal(t, []int{0, 1, 3}, slices.Collect(Accumulate(Range(3), ops.Add)))
	})

	t.Run("early break 1", func(t *testing.T) {
		checkEarlyBreak(t, Accumulate(Range(3), ops.Add), 0)
	})

	t.Run("early break 2", func(t *testing.T) {
		checkEarlyBreak(t, Accumulate(Range(3), ops.Add), 1)
	})
}
