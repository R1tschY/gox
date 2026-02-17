// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

// Reduce applies a function to each element of the sequence, accumulating the result from an initial value.
func Reduce[V any, A any](iter Seq[V], acc A, f func(acc A, val V) A) A {
	for val := range iter {
		acc = f(acc, val)
	}
	return acc
}

// Reduce2 applies a function to each key-value element of the sequence, accumulating the result from an initial value.
func Reduce2[K any, V any, A any](iter Seq2[K, V], acc A, f func(acc A, key K, value V) A) A {
	for key, val := range iter {
		acc = f(acc, key, val)
	}
	return acc
}

func CountIf[V any](iter Seq[V], pred func(V) bool) int {
	return Reduce(iter, 0, func(acc int, value V) int {
		if pred(value) {
			acc += 1
		}
		return acc
	})
}

func Count2If[K any, V any](iter Seq2[K, V], pred func(K, V) bool) int {
	return Reduce2(iter, 0, func(acc int, key K, value V) int {
		if pred(key, value) {
			acc += 1
		}
		return acc
	})
}

func Any[V any](iter Seq[V], pred func(V) bool) bool {
	for val := range iter {
		if pred(val) {
			return true
		}
	}
	return false
}

func Any2[K any, V any](iter Seq2[K, V], pred func(K, V) bool) bool {
	for key, val := range iter {
		if pred(key, val) {
			return true
		}
	}
	return false
}

func All[V any](iter Seq[V], pred func(V) bool) bool {
	for val := range iter {
		if !pred(val) {
			return false
		}
	}
	return true
}

func All2[K any, V any](iter Seq2[K, V], pred func(K, V) bool) bool {
	for key, val := range iter {
		if !pred(key, val) {
			return false
		}
	}
	return true
}

func None[V any](iter Seq[V], pred func(V) bool) bool {
	return !Any(iter, pred)
}

func None2[K any, V any](iter Seq2[K, V], pred func(K, V) bool) bool {
	return !Any2(iter, pred)
}
