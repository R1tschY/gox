// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import (
	"maps"
	"slices"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestLen0(t *testing.T) {
	type testCase struct {
		name string
		arg  Seq0
		want int
	}
	tests := []testCase{
		{
			name: "empty",
			arg:  Empty0(),
			want: 0,
		},
		{
			name: "one",
			arg:  Repeat0Times(1),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Len0(tt.arg))
		})
	}
}

func TestLen(t *testing.T) {
	type testCase[V any] struct {
		name string
		arg  Seq[V]
		want int
	}
	tests := []testCase[int]{
		{
			name: "empty",
			arg:  Empty[int](),
			want: 0,
		},
		{
			name: "one",
			arg:  Singleton[int](42),
			want: 1,
		},
		{
			name: "two",
			arg:  slices.Values([]int{-24, -42}),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Len(tt.arg))
		})
	}
}

//func TestLenOverflow(t *testing.T) {
//	assert.Panics(t, func() { Len(Count[uint64]()) })
//}

func TestLen2(t *testing.T) {
	type testCase struct {
		name string
		arg  Seq2[string, int]
		want int
	}
	tests := []testCase{
		{
			name: "empty",
			arg:  Empty2[string, int](),
			want: 0,
		},
		{
			name: "one",
			arg:  Singleton2[string, int]("TEST", 42),
			want: 1,
		},
		{
			name: "two",
			arg:  maps.All(map[string]int{"A": -24, "B": -42}),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Len2(tt.arg))
		})
	}
}
