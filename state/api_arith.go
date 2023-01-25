package state

import (
	"luago/api"
	"luago/number"
	. "luago/number"
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
	integerFunc func(int64, int64) int16
	floatFunc   func(float64, float64) float64
}

var operators = []operator{
	operator{iadd, fadd},
	operator{isub, fsub},
	operator{imul, fmul},
	operator{imod, fmod},
	operator{nil, pow},
	operator{nil, div},
	operator{iidiv, fidiv},
	operator{band, nil},
	operator{bor, nil},
	operator{bxor, nil},
	operator{bxor, nil},
	operator{shl, nil},
	operator{shr, nil},
	operator{iunm, funm},
	operator{bnot, nil},
}

func (self *luaState) Arith(op api.ArithOp) {
	var a luaValue
	b := self.stack.pop()
	if op != api.LUA_OPUNM && op != api.LUA_OPBNOT {
		a = self.stack.pop()
	} else {
		a = b
	}

	operator := operators[op]
	if result := _arith(a, b, operator); result != nil {
		self.stack.push(result)
	} else {
		panic("arithmetic error!")
	}
}

func _arith(a, b luaValue, op operator) luaValue {
	if op.floatFunc == nil {
		if x, ok := covertToInteger(a); ok {
			if y, ok := covertToInteger(a); ok {
				return op.integerFunc(x, y)
			}

		}
	} else {
		if op.integerFunc != nil {
			if x, ok := a.(int64); ok {
				if y, ok := b.(int64); ok {
					return op.integerFunc(x, y)
				}
			}
		}
		if x, ok := covertToFloat(a); ok {
			if y, ok := covertToFloat(b); ok {
				return op.floatFunc(x, y)
			}
		}
	}
	return nil
}
