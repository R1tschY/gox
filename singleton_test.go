// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestSingleton(t *testing.T) {
	i := 0
	for val := range Singleton[int](-42) {
		assert.Equal(t, -42, val)
		i += 1
	}
	assert.Equal(t, i, 1)
}

func TestSingleton2(t *testing.T) {
	i := 0
	for key, val := range Singleton2[string, int]("TEST", -42) {
		assert.Equal(t, "TEST", key)
		assert.Equal(t, -42, val)
		i += 1
	}
	assert.Equal(t, i, 1)
}
