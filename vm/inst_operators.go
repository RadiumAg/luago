package vm

import "luago/api"

func _binaryArith(i Instruction, vm api.LuaVM, op api.ArithOp) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(b)
	vm.GetRK(c)
	vm.Arith(op)
	vm.Replace(a)
}

func _unaryArith(i Instruction, vm api.LuaVM, op api.ArithOp) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.PushValue(b)
	vm.Arith(op)
	vm.Replace(a)
}

func add(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LUA_OPADD) }  // +
func sub(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LUA_OPSUB) }  // -
func mul(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LUA_OPMUL) }  // ＊
func mod(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LUA_OPMOD) }  // %
func pow(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LUA_OPPOW) }  // ^
func div(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LUA_OPDIV) }  // /
func idiv(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPIDIV) } // //
func band(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPBAND) } // &
func bor(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LUA_OPBOR) }  // |
func bxor(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPBXOR) } // ～
func shl(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LUA_OPSHL) }  // <<
func shr(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LUA_OPSHR) }  // >>
func unm(i Instruction, vm api.LuaVM)  { _unaryArith(i, vm, api.LUA_OPUNM) }   // -
func bnot(i Instruction, vm api.LuaVM) { _unaryArith(i, vm, api.LUA_OPBNOT) }  // ～

func _len(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1

	vm.Len(b)
	vm.Replace(a)
}

func concat(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	c += 1

	n := c - b + 1
	vm.CheckStack(n)
	for i := b; i <= c; i++ {
		vm.PushValue(i)
	}

	vm.Concat(n)
	vm.Replace(a)
}

func _compare(i Instruction, vm api.LuaVM, op api.CompareOp) {
	a, b, c := i.ABC()

	vm.GetRK(b)
	vm.GetRK(c)

	if vm.Compare(-2, -1, op) != (a != 0) {
		vm.AddPC(1)
	}
	vm.Pop(2)
}

func eq(i Instruction, vm api.LuaVM) {
	_compare(i, vm, api.LUA_OPEQ)
}

func lt(i Instruction, vm api.LuaVM) {
	_compare(i, vm, api.LUA_OPLT)
}

func le(i Instruction, vm api.LuaVM) {
	_compare(i, vm, api.LUA_OPLE)
}

func not(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.PushBoolean(!vm.ToBoolean(b))
	vm.Replace(a)
}

func testSet(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	if vm.ToBoolean(b) == (c != 0) {
		vm.Copy(b, a)
	} else {
		vm.AddPC(1)
	}
}

func test(i Instruction, vm api.LuaVM) {
	a, _, c := i.ABC()
	a += 1
	if vm.ToBoolean(a) != (c != 0) {
		vm.AddPC(1)
	}
}

func forPrep(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPSUB)
	vm.Replace(a)
	vm.AddPC(sBx)
}

func forLoop(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1
	vm.PushValue(a + 2)
	vm.PushValue(a)
	vm.Arith(api.LUA_OPADD)
	vm.Replace(a)

	isPositiveStep := vm.ToNumber(a+2) >= 0
	if isPositiveStep && vm.Compare(a, a+1, api.LUA_OPLE) || !isPositiveStep && vm.Compare(a+1, a, api.LUA_OPLE) {
		vm.AddPC(sBx)
		vm.Copy(a, a+3)
	}
}

func jmp(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()

	vm.AddPC(sBx)
	if a != 0 {
		panic("todo：jmp!")
	}
}
