// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package slicez

func ReduceAll[V any, A any](slice []V, initial A, f func(A, int, V) A) A {
	acc := initial
	for i, v := range slice {
		acc = f(acc, i, v)
	}
	return acc
}

func ReduceValues[V any, A any](slice []V, initial A, f func(A, V) A) A {
	acc := initial
	for _, v := range slice {
		acc = f(acc, v)
	}
	return acc
}
