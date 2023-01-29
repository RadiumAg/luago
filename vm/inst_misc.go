package vm

import . "luago/api"

func move(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	vm.Copy(b, a)
}
