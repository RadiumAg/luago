package state

import "luago/binchunk"

type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype
	pc    int

	/** get functions (Lua -> stack) **/

}

func New(stackSize int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(30),
		proto: proto,
		pc:    0,
	}
}
