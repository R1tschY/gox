// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"testing"
)

func TestEmpty0(t *testing.T) {
	for range Empty0() {
		t.Fail()
	}
}

func TestEmpty(t *testing.T) {
	for range Empty[int]() {
		t.Fail()
	}
}

func TestEmpty2(t *testing.T) {
	for range Empty2[string, int]() {
		t.Fail()
	}
}
