package main

import (
	"fmt"
	"io/ioutil"
	. "luago/api"
	"luago/binchunk"
	"luago/state"
	"luago/vm"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		proto := binchunk.Undump(data)
		luaMain(proto)
	}
}

func luaMain(proto *binchunk.Prototype) {
	nRegs := int(proto.MaxStackSize)
	ls := state.New(nRegs+8, proto)
	ls.SetTop(nRegs)
	for {
		fmt.Print(ls.GetTop())
		pc := ls.PC()
		inst := vm.Instruction(ls.Fetch())
		if inst.Opcode() != vm.OP_RETURN {
			inst.Execute(ls)

			fmt.Printf("[%02d] %s ", pc+1, inst.OpName())
			printStack(ls)
		} else {
			break
		}
	}
}

func list(f *binchunk.Prototype) {
	printHeader(f)
	printCode(f)
	printDetail(f)

	for _, p := range f.Protos {
		list(p)
	}
}

func printStack(ls LuaState) {
	top := ls.GetTop()
	for i := 1; i <= top; i++ {
		t := ls.Type(i)
		switch t {
		case LUA_TBOOLEAN:
			fmt.Printf("[%t]", ls.ToBoolean(i))

		case LUA_TNUMBER:
			fmt.Printf("[%g]", ls.ToNumber(i))

		case LUA_TSTRING:
			fmt.Printf("[%q]", ls.ToString(i))

		default:
			fmt.Printf("[%s]", ls.TypeName(t))

		}
	}

	fmt.Println()
}

func printHeader(f *binchunk.Prototype) {
	funcType := "main"
	if f.LineDefined > 0 {
		funcType = "function"
	}

	varargFlag := ""
	if f.IsVararg > 0 {
		varargFlag = "+"
	}

	fmt.Printf("\n%s <%s:%d,%d> (%d instructions) \n", funcType, f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))

	fmt.Printf("%d%s params, %d slots, %d upvalues.", f.NumParams, varargFlag, f.MaxStackSize, len(f.Upvalues))

	fmt.Printf("%d locals,%d constants, %d functions\n", len(f.LocVars), len(f.Constants), len(f.Protos))
}

func printDetail(f *binchunk.Prototype) {
	fmt.Printf("constants (%d):\n", len(f.Constants))
	for i, k := range f.Constants {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i+1, constantToString(k))
	}
}

func constantToString(k interface{}) string {
	switch k.(type) {
	case nil:
		return "nil"

	case bool:
		return fmt.Sprintf("%t", k)

	case float64:
		return fmt.Sprintf("%g", k)

	case int64:
		return fmt.Sprintf("%d", k)

	case string:
		return fmt.Sprintf("%q", k)

	default:
		return "?"
	}
}

func upvalName(f *binchunk.Prototype, idx int) string {
	if len(f.UpvalueNames) > 0 {
		return f.UpvalueNames[idx]
	}
	return "-"
}

func printCode(f *binchunk.Prototype) {
	for pc, c := range f.Code {
		line := "-"
		if len(f.LineInfo) > 0 {
			line = fmt.Sprintf("%d", f.LineInfo[pc])
		}
		i := vm.Instruction(c)
		fmt.Printf("\t%d\t[%s]\t%s\t", pc+1, line, i.OpName())
		printOperands(i)
		fmt.Printf("\n")
	}
}

func printOperands(i vm.Instruction) {
	switch i.OpMode() {
	case vm.IABC:
		a, b, c := i.ABC()
		fmt.Printf("%d", a)
		if i.BMode() != vm.OpArgN {
			if b > 0xff {
				fmt.Printf("%d", -1-b&0xff)
			} else {
				fmt.Printf("%d", b)
			}
		}

		if i.CMode() != vm.OpArgN {
			if c > 0xff {
				fmt.Printf("%d", -1-c&0xff)
			} else {
				fmt.Printf("%d", c)
			}
		}

	case vm.IABx:
		a, bx := i.ABx()
		fmt.Printf("%d", a)
		if i.BMode() == vm.OpArgK {
			fmt.Printf("%d", -1-bx)
		} else if i.BMode() == vm.OpArgU {
			fmt.Printf("%d", bx)
		}

	case vm.IAsBx:
		a, sbx := i.AsBx()
		fmt.Printf("%d", a, sbx)

	case vm.IAx:
		ax := i.Ax()
		fmt.Printf("%d", -1-ax)
	}

}
