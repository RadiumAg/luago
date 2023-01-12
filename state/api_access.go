package state

import (
	"fmt"
	. "luago/api"
)

func (self *luaState) TypeName(tp LuaType) string {
	switch tp {
	case LUA_TNONE:
		return "no value"

	case LUA_TNIL:
		return "nil"

	case LUA_TNUMBER:
		return "number"

	case LUA_TSTRING:
		return "string"

	case LUA_TTABLE:
		return "table"

	case LUA_TFUNCTION:
		return "function"

	case LUA_TTHEAD:
		return "thread"

	default:
		return "userdata"

	}
}

func (self *luaState) Type(idx int) LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

func (self *luaState) IsNone(idx int) bool {
	return self.Type(idx) == LUA_TNONE
}

func (self *luaState) IsNoneOrNil(idx int) bool {
	return self.Type(idx) <= LUA_TNIL
}

func (self *luaState) IsBoolean(idx int) bool {
	return self.Type(idx) == LUA_TBOOLEAN
}

func (self *luaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

func (self *luaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

func (self *luaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (self *luaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return covertToBoolean(val)
}

func covertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}

func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}

func (self *luaState) ToInteger(idx int) int64 {
	i, _ := self.ToIntegerX(idx)
	return i
}

func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

func (self *luaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case string:
		return x, true

	case int64, float64:
		s := fmt.Sprintf("%w", x)
		self.stack.set(idx, s)
		return s, true

	default:
		return "", false
	}
}

func (self *luaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}