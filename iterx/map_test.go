// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len(Map(Empty[int](), toString)))
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, "42", collectStrings(Map(Singleton[int](42), toString)))
	})

	t.Run("many", func(t *testing.T) {
		assert.Equal(t, "1 2 3", collectStrings(Map(slices.Values([]int{1, 2, 3}), toString)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak(t, Map(RepeatTimes(1, 2), toString), 0)
	})
}

func TestMap2(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len2(Map2(Empty2[int, bool](), toKvStrings)))
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, "42:true", collectSeq2(Map2(Singleton2[int, bool](42, true), toKvStrings)))
	})

	t.Run("many", func(t *testing.T) {
		assert.Equal(t, "0:1 1:2 2:3", collectSeq2(Map2(slices.All([]int{1, 2, 3}), toKvStrings)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak2(t, Map2(Repeat2Times(1, true, 2), toKvStrings), 0)
	})
}

func TestMap2To1(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len(Map2To1(Empty2[int, bool](), toKvString)))
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, "42:true", collectStrings(Map2To1(Singleton2[int, bool](42, true), toKvString)))
	})

	t.Run("many", func(t *testing.T) {
		assert.Equal(t, "0:1 1:2 2:3", collectStrings(Map2To1(slices.All([]int{1, 2, 3}), toKvString)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak(t, Map2To1(Repeat2Times(1, true, 2), toKvString), 0)
	})
}

func to2(i int) (int, string) {
	return i, nthChar(i)
}

func nthChar(n int) string {
	return string('a' + rune(n))
}

func TestMapTo2(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len2(MapTo2(Empty[int](), to2)))
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, "1:b", collectSeq2(MapTo2(Singleton[int](1), to2)))
	})

	t.Run("many", func(t *testing.T) {
		assert.Equal(t, "0:a 1:b 2:c", collectSeq2(MapTo2(slices.Values([]int{0, 1, 2}), to2)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak2(t, MapTo2(RepeatTimes(1, 2), to2), 0)
	})
}
