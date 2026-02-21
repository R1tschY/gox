// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package iterx

/*
Seq0 is a type alias representing a generator function.
It can be used by a [range loop], as in:

	func doSomething(it iter.Seq0) {
		for range it {
			// ...
		}
	}

[range loop]: https://go.dev/ref/spec#For_range
*/
type Seq0 = func(yield func() bool)

// Seq is a type alias representing [iter.Seq] iterator function.
// See [iter] package for details.
type Seq[V any] = func(yield func(V) bool)

// Seq2 is a type alias representing [iter.Seq2] iterator function.
// See [iter] package for details.
type Seq2[K any, V any] = func(yield func(K, V) bool)
