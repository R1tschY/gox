// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package slicez

func MapValues[V any, W any](slice []V, f func(V) W) []W {
	res := make([]W, len(slice))
	for i, v := range slice {
		res[i] = f(v)
	}
	return res
}

func Map[V any, W any](slice []V, f func(int, V) W) []W {
	res := make([]W, len(slice))
	for i, v := range slice {
		res[i] = f(i, v)
	}
	return res
}

func Filter[V any](slice []V, f func(int, V) bool) []V {
	res := make([]V, 0)
	for i, v := range slice {
		if f(i, v) {
			res = append(res, v)
		}
	}
	return res
}

func FilterValues[V any](slice []V, f func(V) bool) []V {
	res := make([]V, 0)
	for _, v := range slice {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
