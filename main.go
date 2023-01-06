package main

import (
	"fmt"
	"io/ioutil"
	binchunk "luago/bingchunk"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		proto := binchunk.Undump(data)
		list(proto)
	}
}

func list(f *binchunk.Prototype) {

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

func printCode(f *binchunk.Prototype) {
	for pc, c := range f.Code {
		line := "-"
		if len(f.LineInfo) > 0 {
			line = fmt.Sprintf("%d", f.LineInfo[pc])
		}
		fmt.Printf("\t%d\t[%s]\t0x%08x\n", pc+1, line, c)
	}
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
	if len(f.UpValueNames) > 0 {
		return f.UpValueNames[idx]
	}
	return "-"
}
