package vm

import "luago/api"

const (
	IABC = iota
	IABx
	IAsBx
	IAx
)

const (
	OP_MOVE = iota
	OP_LOADK
	OP_LOADKX
	OP_LOADBOOL
	OP_LOADNIL
	OP_GETUPVAL
	OP_GETTABUP
	OP_GETTABLE
	OP_SETTABUP
	OP_SETUPVAL
	OP_SETTABLE
	OP_NEWTABLE
	OP_SELF
	OP_ADD
	OP_SUB
	OP_MUL
	OP_MOD
	OP_POW
	OP_DIV
	OP_IDIV
	OP_BAND
	OP_BOR
	OP_BXOR
	OP_SHL
	OP_SHR
	OP_UNM
	OP_BNOT
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_JMP
	OP_EQ
	OP_LT
	OP_LE
	OP_TEST
	OP_TESTSET
	OP_CALL
	OP_TAILCALL
	OP_RETURN
	OP_FORLOOP
	OP_FORPREP
	OP_TFORCALL
	OP_TFORLOOP
	OP_SETLIST
	OP_CLOSURE
	OP_VARARG
	OP_EXTRAARG
)

const (
	OpArgN = iota // argument is not used
	OpArgU        // argument is used
	OpArgR        // argument is a register or a jump offset
	OpArgK        // argument is a constant or register/constant
)

type opcode struct {
	testFlag byte
	setAFlag byte
	argBMode byte
	argCMode byte
	opMode   byte
	name     string
	action   func(i Instruction, vm api.LuaVM)
}

var opcodes = []opcode{
	// T  A     B        C      mode     name
	{0, 1, OpArgR, OpArgN, IABC, "MOVE     ", move},
	{0, 1, OpArgK, OpArgN, IABx, "LOADK    ", loadK},
	{0, 1, OpArgN, OpArgN, IABx, "LOADKX  ", loadKx},
	{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL", loadBool},
	{0, 1, OpArgU, OpArgN, IABC, "LOADNIL ", loadNil},
	{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL"},
	{0, 1, OpArgU, OpArgK, IABC, "GETTABUP"},
	{0, 1, OpArgR, OpArgK, IABC, "GETTABLE"},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABUP"},
	{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL"},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABLE"},
	{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE"},
	{0, 1, OpArgR, OpArgK, IABC, "SELF     "},
	{0, 1, OpArgK, OpArgK, IABC, "ADD      ", add},
	{0, 1, OpArgK, OpArgK, IABC, "SUB      ", sub},
	{0, 1, OpArgK, OpArgK, IABC, "MUL      ", mul},
	{0, 1, OpArgK, OpArgK, IABC, "MOD      ", mod},
	{0, 1, OpArgK, OpArgK, IABC, "POW      ", pow},
	{0, 1, OpArgK, OpArgK, IABC, "DIV      ", div},
	{0, 1, OpArgK, OpArgK, IABC, "IDIV     ", idiv},
	{0, 1, OpArgK, OpArgK, IABC, "BAND     ", band},
	{0, 1, OpArgK, OpArgK, IABC, "BOR      ", bor},
	{0, 1, OpArgK, OpArgK, IABC, "BXOR     ", bxor},
	{0, 1, OpArgK, OpArgK, IABC, "SHL      ", shl},
	{0, 1, OpArgK, OpArgK, IABC, "SHR      ", shr},
	{0, 1, OpArgR, OpArgN, IABC, "UNM      ", unm},
	{0, 1, OpArgR, OpArgN, IABC, "BNOT     ", bnot},
	{0, 1, OpArgR, OpArgN, IABC, "NOT      ", not},
	{0, 1, OpArgR, OpArgN, IABC, "LEN      ", len},
	{0, 1, OpArgR, OpArgR, IABC, "CONCAT  ", concat},
	{0, 0, OpArgR, OpArgN, IAsBx, "JMP      "},
	{1, 0, OpArgK, OpArgK, IABC, "EQ       ", eq},
	{1, 0, OpArgK, OpArgK, IABC, "LT       ", lt},
	{1, 0, OpArgK, OpArgK, IABC, "LE       ", le},
	{1, 0, OpArgN, OpArgU, IABC, "TEST     ", test},
	{1, 1, OpArgR, OpArgU, IABC, "TESTSET ", testSet},
	{0, 1, OpArgU, OpArgU, IABC, "CALL     "},
	{0, 1, OpArgU, OpArgU, IABC, "TAILCALL"},
	{0, 0, OpArgU, OpArgN, IABC, "RETURN  "},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP ", forLoop},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP ", forPrep},
	{0, 0, OpArgN, OpArgU, IABC, "TFORCALL"},
	{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP"},
	{0, 0, OpArgU, OpArgU, IABC, "SETLIST "},
	{0, 1, OpArgU, OpArgN, IABx, "CLOSURE ", close},
	{0, 1, OpArgU, OpArgN, IABC, "VARARG  "},
	{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG"},
}
