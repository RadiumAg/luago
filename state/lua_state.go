package state

import "luago/binchunk"

type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype
	pc    int
}

func New() *luaState {
	return &luaState{
		stack: newLuaStack(30),
	}
}
