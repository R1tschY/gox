// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package ops

import "github.com/R1tschy/gox/typex"

func Add[T typex.Numeric](lhs T, rhs T) T {
	return lhs + rhs
}

func Sub[T typex.Numeric](lhs T, rhs T) T {
	return lhs - rhs
}

func Mul[T typex.Numeric](lhs T, rhs T) T {
	return lhs * rhs
}

func Div[T typex.Numeric](lhs T, rhs T) T {
	return lhs / rhs
}

func Mod[T typex.Integer](lhs T, rhs T) T {
	return lhs % rhs
}
