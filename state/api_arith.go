package state

import (
	"luago/number"
	"math"
)

var (
	iadd = func(a, b int64) int64 {
		return a + b
	}

	fadd = func(a, b float64) float64 {
		return a + b
	}

	isub = func(a, b int64) int64 {
		return a - b
	}

	fsub = func(a, b float64) float64 {
		return a - b
	}
	imul = func(a, b int64) int64 {
		return a * b
	}
	fmul = func(a, b float64) float64 {
		return a * b
	}
	imod = number.IMode
	fmod = number.FMode
	pow  = math.Pow
	div  = func(a, b float64) float64 {
		return a / b
	}
	iidiv = number.IFloorDiv
	fidiv = number.FFloorDiv
	band  = func(a, b int64) int64 {
		return a & b
	}
	bor = func(a, b int64) int64 {
		return a | b
	}
	bxor = func(a, b int64) int64 {
		return a ^ b
	}
	shl  = number.ShiftLeft
	shr  = number.ShiftRight
	iunm = func(a, _ int64) int64 {
		return -a
	}
	funm = func(a, _ float64) float64 {
		return -a
	}
	bnot = func(a, _ int64) int64 {
		return ^a
	}
)

type operator struct {
	integerFunc func(int64, int64)
	floatFunc   func(float64, float64) float64
}
