// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func collectStrings(it Seq[string]) string {
	return strings.Join(slices.Collect(it), " ")
}

func collectSeq2[K any, V any](it Seq2[K, V]) string {
	res := make([]string, 0, 16)
	for k, v := range it {
		res = append(res, toKvString(k, v))
	}
	return strings.Join(res, " ")
}

func collectSeq[V any](it Seq[V]) string {
	return strings.Join(slices.Collect(Map(it, toString)), " ")
}

func toString[V any](v V) string {
	return fmt.Sprintf("%v", v)
}

func toKvString[K any, V any](k K, v V) string {
	return fmt.Sprintf("%v:%v", k, v)
}

func toKvStrings[K any, V any](k K, v V) (string, string) {
	return fmt.Sprintf("%v", k), fmt.Sprintf("%v", v)
}

func range2(end int) Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i := 0; i < end; i++ {
			if !yield(i, fmt.Sprintf("%d", i)) {
				break
			}
		}
	}
}

func checkEarlyBreak0(t *testing.T, it Seq0, n int) {
	i := 0
	for range it {
		if i == n {
			break
		}
		i++
	}
	assert.Equal(t, n, i)
}

func checkEarlyBreak[V any](t *testing.T, it Seq[V], n int) {
	i := 0
	for range it {
		if i == n {
			break
		}
		i++
	}
	assert.Equal(t, n, i)
}

func checkEarlyBreak2[K any, V any](t *testing.T, it Seq2[K, V], n int) {
	i := 0
	for range it {
		if i == n {
			break
		}
		i++
	}
	assert.Equal(t, n, i)
}

// Predicates

func all[V any](_ V) bool              { return true }
func all2[K any, V any](_ K, _ V) bool { return true }
func odd(i int) bool                   { return i%2 == 1 }
func oddValue(_ int, v int) bool       { return v%2 == 1 }
func oddKey(k int, _ int) bool         { return k%2 == 1 }
