package binchunk

type Prototype struct {
	Source          string
	LineDefined     uint32
	LastLineDefined uint32
	NumParams       uint32
	IsVararg        byte
	MaxStackSize    byte
	Code            []uint32
	Constants       []interface{}
	Upvalues        []Upvalue
	Protos          []*Prototype
	LineInfo        []uint32
	LocVars         []LocVars
	UpValueNames    []string
}

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0 * 53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSZITE_SIZE      = 8
	LUA_INTEGER_SIZE = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

type binaryChunk struct {
	header
	sizeUpvalues byte
	mainFunc     *Prototype
}
