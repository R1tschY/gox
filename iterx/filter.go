// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

func Filter[T any](it Seq[T], pred func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		for val := range it {
			if pred(val) {
				if !yield(val) {
					break
				}
			}
		}
	}
}

func Filter2[K any, V any](it Seq2[K, V], pred func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range it {
			if pred(key, val) {
				if !yield(key, val) {
					break
				}
			}
		}
	}
}

func DropWhile[T any](it Seq[T], pred func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		dropping := true
		for val := range it {
			if dropping {
				if !pred(val) {
					dropping = false
				} else {
					continue
				}
			}

			if !yield(val) {
				break
			}
		}
	}
}

func DropWhile2[K any, V any](it Seq2[K, V], pred func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		dropping := true
		for key, val := range it {
			if dropping {
				if !pred(key, val) {
					dropping = false
				} else {
					continue
				}
			}

			if !yield(key, val) {
				break
			}
		}
	}
}

func TakeWhile[T any](it Seq[T], pred func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		for val := range it {
			if !pred(val) || !yield(val) {
				break
			}
		}
	}
}

func TakeWhile2[K any, V any](it Seq2[K, V], pred func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range it {
			if !pred(key, val) || !yield(key, val) {
				break
			}
		}
	}
}
