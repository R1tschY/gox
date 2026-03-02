// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len(Chain[int]()))
	})

	t.Run("chain empty seqs", func(t *testing.T) {
		assert.Equal(t, 0, Len(Chain(Empty[int](), Empty[int]())))
	})

	t.Run("chain seqs", func(t *testing.T) {
		assert.Equal(t, []int{1, 2}, slices.Collect(Chain(Singleton(1), Singleton(2))))
	})

	t.Run("early break at 0", func(t *testing.T) {
		checkEarlyBreak(t, Chain(Singleton(1), Singleton(2)), 0)
	})

	t.Run("early break at 1", func(t *testing.T) {
		checkEarlyBreak(t, Chain(Singleton(1), Singleton(2)), 1)
	})
}

func TestChainFrom(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len(ChainFrom(Empty[Seq[int]]())))
	})

	t.Run("chain empty seqs", func(t *testing.T) {
		assert.Equal(t, 0, Len(ChainFrom(SeqOf(Empty[int](), Empty[int]()))))
	})

	t.Run("chain seqs", func(t *testing.T) {
		assert.Equal(t, []int{1, 2}, slices.Collect(ChainFrom(SeqOf(Singleton(1), Singleton(2)))))
	})

	t.Run("early break at 0", func(t *testing.T) {
		checkEarlyBreak(t, ChainFrom(SeqOf(Singleton(1), Singleton(2))), 0)
	})

	t.Run("early break at 1", func(t *testing.T) {
		checkEarlyBreak(t, ChainFrom(SeqOf(Singleton(1), Singleton(2))), 1)
	})
}

func TestChain2(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len2(Chain2[string, int]()))
	})

	t.Run("chain empty seqs", func(t *testing.T) {
		assert.Equal(t, 0, Len2(Chain2(Empty2[string, int](), Empty2[string, int]())))
	})

	t.Run("chain seqs", func(t *testing.T) {
		assert.Equal(t, "1:1 2:2", collectSeq2(Chain2(Singleton2("1", 1), Singleton2("2", 2))))
	})

	t.Run("early break at 0", func(t *testing.T) {
		checkEarlyBreak2(t, Chain2(Singleton2("1", 1), Singleton2("2", 2)), 0)
	})

	t.Run("early break at 1", func(t *testing.T) {
		checkEarlyBreak2(t, Chain2(Singleton2("1", 1), Singleton2("2", 2)), 1)
	})
}

func TestChainFrom2(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, Len2(ChainFrom2(Empty[Seq2[string, int]]())))
	})

	t.Run("chain empty seqs", func(t *testing.T) {
		assert.Equal(t, 0, Len2(ChainFrom2(SeqOf(Empty2[string, int](), Empty2[string, int]()))))
	})

	t.Run("chain seqs", func(t *testing.T) {
		assert.Equal(t, "1:1 2:2", collectSeq2(ChainFrom2(SeqOf(Singleton2("1", 1), Singleton2("2", 2)))))
	})

	t.Run("early break at 0", func(t *testing.T) {
		checkEarlyBreak2(t, ChainFrom2(SeqOf(Singleton2("1", 1), Singleton2("2", 2))), 0)
	})

	t.Run("early break at 1", func(t *testing.T) {
		checkEarlyBreak2(t, ChainFrom2(SeqOf(Singleton2("1", 1), Singleton2("2", 2))), 1)
	})
}
