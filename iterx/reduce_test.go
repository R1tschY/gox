// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, uint64(42), Reduce(Empty[uint8](), uint64(42), func(acc uint64, val uint8) uint64 { return acc + uint64(val) }))
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, uint64(33), Reduce(Singleton[uint8](11), uint64(22), func(acc uint64, val uint8) uint64 { return acc + uint64(val) }))
	})
}
