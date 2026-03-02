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

func BitXor[T typex.Integer](lhs T, rhs T) T {
	return lhs ^ rhs
}

func BitAnd[T typex.Integer](lhs T, rhs T) T {
	return lhs & rhs
}

func BitOr[T typex.Integer](lhs T, rhs T) T {
	return lhs | rhs
}

func Lshift[T typex.Integer, S typex.Integer](lhs T, rhs S) T {
	return lhs << rhs
}

func Rshift[T typex.Integer, S typex.Integer](lhs T, rhs S) T {
	return lhs >> rhs
}

func Lt[T typex.Ordered](lhs T, rhs T) bool {
	return lhs < rhs
}

func Le[T typex.Ordered](lhs T, rhs T) bool {
	return lhs <= rhs
}

func Gt[T typex.Ordered](lhs T, rhs T) bool {
	return lhs > rhs
}

func Ge[T typex.Ordered](lhs T, rhs T) bool {
	return lhs >= rhs
}

func Eq[T comparable](lhs T, rhs T) bool {
	return lhs == rhs
}

func Ne[T comparable](lhs T, rhs T) bool {
	return lhs != rhs
}

func IsNil[T comparable](op *T) bool {
	return op == nil
}

func IsZero[T comparable](op T) bool {
	var zero T
	return op == zero
}

func Not(op bool) bool {
	return !op
}

func Box[T any](op T) *T {
	return &op
}
