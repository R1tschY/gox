// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package dsx

import (
	"slices"

	"github.com/R1tschy/gox/iterx"
)

// Vec is a better interface for a slice
type Vec[V any] struct {
	data []V
}

func NewVec[V any]() Vec[V] {
	return Vec[V]{
		data: make([]V, 0),
	}
}

func NewVecWithCapacity[V any](cap int) Vec[V] {
	return Vec[V]{
		data: make([]V, 0, cap),
	}
}

func (v Vec[V]) Append(val V) {
	v.data = append(v.data, val)
}

func (v Vec[V]) AppendSlice(vals []V) {
	v.data = append(v.data, vals...)
}

func (v Vec[V]) AppendSeq(seq iterx.Seq[V]) {
	v.data = slices.AppendSeq(v.data, seq)
}

func (v Vec[V]) Len() int {
	return len(v.data)
}

func (v Vec[V]) Iter() iterx.Seq[V] {
	return slices.Values(v.data)
}

func (v Vec[V]) Clear() {
	clear(v.data)
}
