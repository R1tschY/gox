// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len(Chunk(Range(0), 1)))
	})

	t.Run("n=1", func(t *testing.T) {
		assert.Equal(t, [][]int{{0}, {1}, {2}}, slices.Collect(Chunk(Range(3), 1)))
	})

	t.Run("n=2", func(t *testing.T) {
		assert.Equal(t, [][]int{{0, 1}, {2}}, slices.Collect(Chunk(Range(3), 2)))
	})

	t.Run("one incomplete", func(t *testing.T) {
		assert.Equal(t, [][]int{{0, 1, 2}}, slices.Collect(Chunk(Range(3), 200)))
	})

	t.Run("early break", func(t *testing.T) {
		checkEarlyBreak(t, Chunk(Range(3), 2), 0)
	})
}

func chunkSlice[V any](seq []V, n int) [][]V {
	res := make([][]V, 0)
	chunk := make([]V, 0, n)

	for _, v := range seq {
		chunk = append(chunk, v)
		if len(chunk) == n {
			res = append(res, chunk)
			chunk = make([]V, 0, n)
		}
	}
	if len(chunk) > 0 {
		res = append(res, chunk)
	}
	return res
}

func chunkSliceV2[V any](seq []V, n int) [][]V {
	res := make([][]V, 0)
	chunk := []V(nil)

	for _, v := range seq {
		if chunk == nil {
			chunk = make([]V, 0, n)
		}
		chunk = append(chunk, v)
		if len(chunk) == n {
			res = append(res, chunk)
			chunk = nil
		}
	}
	if len(chunk) > 0 {
		res = append(res, chunk)
	}
	return res
}

func chunkV2[V any](it Seq[V], n int) Seq[[]V] {
	if n <= 0 {
		panic("n must be positive")
	}

	return func(yield func([]V) bool) {
		chunk := make([]V, 0, n)

		it(func(v V) bool {
			chunk = append(chunk, v)
			if len(chunk) == n {
				if !yield(chunk) {
					chunk = nil
					return false
				}
				chunk = make([]V, 0, n)
			}
			return true
		})

		if len(chunk) > 0 {
			yield(chunk)
		}
	}
}

func chunkV3[V any](it Seq[V], n int) Seq[[]V] {
	if n <= 0 {
		panic("n must be positive")
	}

	return func(yield func([]V) bool) {
		chunk := []V(nil)

		it(func(v V) bool {
			if chunk == nil {
				chunk = make([]V, 0, n)
			}
			chunk = append(chunk, v)
			if len(chunk) == n {
				if !yield(chunk) {
					chunk = nil
					return false
				}
				chunk = nil
			}
			return true
		})

		if len(chunk) > 0 {
			yield(chunk)
		}
	}
}

func chunkSeq(slice []byte) [][]byte {
	return slices.Collect(Chunk(slices.Values(slice), 32))
}

func chunkSeqPlain(slice []byte) [][]byte {
	return slices.Collect(chunkV2(slices.Values(slice), 32))
}

func chunkSeqAlloc(slice []byte) [][]byte {
	return slices.Collect(chunkV3(slices.Values(slice), 32))
}

func BenchmarkChunk(b *testing.B) {
	slice := make([]byte, 320)

	b.Run("slice builtin", func(b *testing.B) {
		for b.Loop() {
			_ = func() [][]byte {
				return slices.Collect(slices.Chunk(slice, 32))
			}()
		}
	})

	b.Run("slice for", func(b *testing.B) {
		for b.Loop() {
			_ = chunkSlice(slice, 32)
		}
	})

	b.Run("slice for alloc", func(b *testing.B) {
		for b.Loop() {
			_ = chunkSliceV2(slice, 32)
		}
	})

	b.Run("seq", func(b *testing.B) {
		for b.Loop() {
			_ = chunkSeq(slice)
		}
	})

	b.Run("seq plain", func(b *testing.B) {
		for b.Loop() {
			_ = chunkSeqPlain(slice)
		}
	})

	b.Run("seq alloc", func(b *testing.B) {
		for b.Loop() {
			_ = chunkSeqAlloc(slice)
		}
	})
}
