// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

// Empty0 returns an empty unit sequence [Seq0].
func Empty0() Seq0 {
	return func(yield func() bool) {}
}

// Empty returns an empty sequence [Seq].
func Empty[V any]() Seq[V] {
	return func(yield func(V) bool) {}
}

// Empty2 returns an empty key-value sequence [Seq2].
func Empty2[K any, V any]() Seq2[K, V] {
	return func(yield func(K, V) bool) {}
}
