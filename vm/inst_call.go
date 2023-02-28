package vm

import (
	. "luago/api"
)

func closure(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1

	vm.LoadProto(bx)
	vm.Replace(a)
}

func call(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1

	nArgs := _pushFuncAndArgs(a, b, vm)
}
