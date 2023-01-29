package vm

import "luago/api"

func _binaryArith(i Instruction, vm api.LuaVM, op api.ArithOp) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRk(b)
	vm.GetRk(c)
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
	vm.Len(b)
	vm.Replace(a)
}

func concat(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	n := c - b + 1
	vm.CheckStack(n)
	for i := b; i <= c; i++ {
		vm.PushValue(i)
	}

	vm.Concat(n)
	vm.Replace(a)
}
