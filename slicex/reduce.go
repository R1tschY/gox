// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package slicex

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

func Find[V any](slice []V, f func(V) bool) (V, bool) {
	for _, v := range slice {
		if f(v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

func CountIf[V any](slice []V, pred func(int, V) bool) int {
	return ReduceAll(slice, 0, func(acc int, idx int, value V) int {
		if pred(idx, value) {
			acc += 1
		}
		return acc
	})
}

func Any[V any](slice []V, pred func(int, V) bool) bool {
	for idx, val := range slice {
		if pred(idx, val) {
			return true
		}
	}
	return false
}

func All[V any](slice []V, pred func(int, V) bool) bool {
	for idx, val := range slice {
		if !pred(idx, val) {
			return false
		}
	}
	return true
}

func None[V any](slice []V, pred func(int, V) bool) bool {
	return !Any(slice, pred)
}
