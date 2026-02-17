// Copyright (c) 2026 Richard Liebscher
// SPDX-License-Identifier: MIT

package ops

import "github.com/R1tschy/iterz/internal"

func Add[T internal.Numeric](lhs T, rhs T) T {
	return lhs + rhs
}

func Sub[T internal.Numeric](lhs T, rhs T) T {
	return lhs - rhs
}

func Mul[T internal.Numeric](lhs T, rhs T) T {
	return lhs * rhs
}

func Div[T internal.Numeric](lhs T, rhs T) T {
	return lhs / rhs
}

func Mod[T internal.Integer](lhs T, rhs T) T {
	return lhs % rhs
}
