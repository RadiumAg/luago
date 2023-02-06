package vm

import . "luago/api"

func newTable(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.CreateTable(Fb2int(b), Fb2int(c))
	vm.Replace(a)
}

func getTable(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
}
