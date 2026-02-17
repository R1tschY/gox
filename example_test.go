// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterz

import (
	"fmt"
	"maps"
	"slices"

	"github.com/R1tschy/iterz/ops"
)

// Len

func ExampleLen0() {
	fmt.Println(Len0(Repeat0Times(3)))
	// Output: 3
}

func ExampleLen() {
	fmt.Println(Len(slices.Values([]int{1, 2, 3})))
	// Output: 3
}

func ExampleLen2() {
	fmt.Println(Len2(slices.All([]int{1, 2, 3})))
	// Output: 3
}

// Reduce

func ExampleReduce() {
	fmt.Println(Reduce(slices.Values([]int{1, 2, 3}), 0, func(acc, val int) int { return acc + val }))
	// Output: 6
}

func ExampleReduce2() {
	fmt.Println(Reduce2(slices.All([]int{1, 2, 3}), 0, func(acc, key int, val int) int { return acc + key*val }))
	// Output: 8
}

// Repeat

func ExampleRepeat() {
	fmt.Println(slices.Collect(RepeatTimes("A", 3)))
	// Output: [A A A]
}

func ExampleRepeat2Times() {
	fmt.Println(Reduce2(Repeat2Times("K", "V", 3), "", func(acc, k string, v string) string { return acc + "[" + k + " " + v + "]" }))
	// Output: [K V][K V][K V]
}

// Map

func ExampleMap() {
	fmt.Println(slices.Collect(Map(slices.Values([]int{1, 2, 3}), func(i int) string { return fmt.Sprintf("<%d>", i) })))
	// Output: [<1> <2> <3>]
}

func ExampleMap2() {
	fmt.Println(maps.Collect(Map2(slices.All([]string{"A", "B", "C"}), func(i int, s string) (int, string) { return i + 1, s })))
	// Output: map[1:A 2:B 3:C]
}

func ExampleMap2To1() {
	fmt.Println(slices.Collect(Map2To1(slices.All([]string{"A", "B", "C"}), func(i int, s string) int { return i + 1 })))
	// Output: [1 2 3]
}

func ExampleMapTo2() {
	type Pair[F any, S any] struct {
		First  F
		Second S
	}

	pairs := []Pair[string, int]{
		{
			First:  "A",
			Second: 1,
		},
	}

	fmt.Println(maps.Collect(MapTo2(slices.Values(pairs), func(pair Pair[string, int]) (string, int) {
		return pair.First, pair.Second
	})))
	// Output: map[A:1]
}

// Filter

func ExampleFilter() {
	fmt.Println(slices.Collect(Filter(slices.Values([]int{1, 2, 3}), func(i int) bool { return i%2 == 0 })))
	// Output: [2]
}

func ExampleFilter2() {
	fmt.Println(maps.Collect(Filter2(slices.All([]int{1, 2, 3}), func(i int, val int) bool { return i%2 == 0 })))
	// Output: map[0:1 2:3]
}

// Count

func ExampleCount_positive() {
	fmt.Println(slices.Collect(Limit(Count(1, 1), 3)))
	// Output: [1 2 3]
}

func ExampleCount_negative() {
	fmt.Println(slices.Collect(Limit(Count(1, -1), 3)))
	// Output: [1 0 -1]
}

// Accumulate

func ExampleAccumulate_runningMaximum() {
	data := []int{3, 4, 6, 2, 1, 9, 0, 7, 5, 8}
	fmt.Println(slices.Collect(Accumulate(slices.Values(data), func(l int, r int) int {
		return max(l, r)
	})))
	// Output: [3 4 6 6 6 9 9 9 9 9]
}

func ExampleAccumulate_runningProduct() {
	data := []int{3, 4, 6, 2, 1, 9, 0, 7, 5, 8}
	fmt.Println(slices.Collect(Accumulate(slices.Values(data), ops.Mul)))
	// Output: [3 12 72 144 144 1296 0 0 0 0]
}
