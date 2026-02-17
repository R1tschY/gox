// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestRepeat0(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		assert.Equal(t, 0, Len0(Repeat0Times(0)))
	})

	t.Run("more", func(t *testing.T) {
		assert.Equal(t, 42, Len0(Repeat0Times(42)))
	})

	t.Run("uint64", func(t *testing.T) {
		assert.Equal(t, 1, Len0(Repeat0Times(uint64(1))))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak0(t, Repeat0Times(5), 0)
	})
}

func TestRepeat(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		assert.Equal(t, 0, Len(RepeatTimes("VALUE", 0)))
	})

	t.Run("more", func(t *testing.T) {
		assert.Equal(t, 42, Len(RepeatTimes("VALUE", 42)))
	})

	t.Run("value", func(t *testing.T) {
		for value := range RepeatTimes("VALUE", 2) {
			assert.Equal(t, "VALUE", value)
		}
	})

	t.Run("uint64", func(t *testing.T) {
		assert.Equal(t, 1, Len(RepeatTimes("VALUE", uint64(1))))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak(t, RepeatTimes("VALUE", 2), 0)
	})
}

func TestRepeat2(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		assert.Equal(t, 0, Len2(Repeat2Times(1, "VALUE", 0)))
	})

	t.Run("more", func(t *testing.T) {
		assert.Equal(t, 42, Len2(Repeat2Times(24, "VALUE", 42)))
	})

	t.Run("key", func(t *testing.T) {
		for key, _ := range Repeat2Times("KEY", "VALUE", 2) {
			assert.Equal(t, "KEY", key)
		}
	})

	t.Run("value", func(t *testing.T) {
		for _, value := range Repeat2Times("KEY", "VALUE", 2) {
			assert.Equal(t, "VALUE", value)
		}
	})

	t.Run("uint64", func(t *testing.T) {
		assert.Equal(t, 1, Len2(Repeat2Times("KEY", "VALUE", uint64(1))))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak2(t, Repeat2Times(24, "VALUE", 2), 0)
	})
}
