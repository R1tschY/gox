// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len(Filter(Empty[int](), odd)))
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, "1", collectSeq(Filter(Singleton[int](1), odd)))
	})

	t.Run("many", func(t *testing.T) {
		assert.Equal(t, "1 3", collectSeq(Filter(slices.Values([]int{1, 2, 3}), odd)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak(t, Filter(RepeatTimes(1, 2), all), 0)
	})
}

func TestFilter2(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len2(Filter2(Empty2[int, int](), oddKey)))
	})

	t.Run("key", func(t *testing.T) {
		assert.Equal(t, "1:2", collectSeq2(Filter2(slices.All([]int{1, 2, 3}), oddKey)))
	})

	t.Run("value", func(t *testing.T) {
		assert.Equal(t, "0:1 2:3", collectSeq2(Filter2(slices.All([]int{1, 2, 3}), oddValue)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak2(t, Filter2(Repeat2Times(1, 2, 2), all2), 0)
	})
}
