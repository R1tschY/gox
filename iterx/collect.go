// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"maps"
	"slices"
)

type Extend[T any] interface {
	AppendSeq(seq Seq[T])
}

func Collect[V any, T Extend[V]](target T, seq Seq[V]) T {
	target.AppendSeq(seq)
	return target
}

func CollectSlice[V any](seq Seq[V]) []V {
	return slices.Collect(seq)
}

func CollectMap[K comparable, V any](seq Seq2[K, V]) map[K]V {
	return maps.Collect(seq)
}

func TryAppendToSlice[V any](s []V, seq Seq2[V, error]) ([]V, error) {
	for v, err := range seq {
		if err != nil {
			return nil, err
		}
		s = append(s, v)
	}
	return s, nil
}

func TryCollectSlice[V any](seq Seq2[V, error]) ([]V, error) {
	return TryAppendToSlice(nil, seq)
}

func First[V any](seq Seq[V]) (V, bool) {
	var val V
	seen := false

	for val = range seq {
		seen = true
		break
	}

	return val, seen
}

func Single[V any](seq Seq[V]) (V, bool) {
	var val V
	seen := false

	for val = range seq {
		if seen {
			return val, false
		}
		seen = true
	}

	return val, seen
}

func Consume[V any](seq Seq[V]) {
	for range seq {
	}
}

func Consume2[K any, V any](seq Seq2[K, V]) {
	for range seq {
	}
}
