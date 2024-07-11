package filter

import "testing"

// 测试 FromString().Value()
func TestFromStringValue(t *testing.T) {
	t.Log(FromString("   111   ").TrimSpace().Value())
}

// 测试 FromString().DefaultValue()
func TestFromStringDefaultValue(t *testing.T) {
	t.Log(FromString("").Require().DefaultValue("ok"))
}

// 测试 FromString().Error()
func TestFromStringError(t *testing.T) {
	t.Log(FromString(" ").Require().Error())
	t.Log(FromString("").Require().Error())
}

// 测试 FromString().Result()
func TestFromStringResult(t *testing.T) {
	t.Log(FromString("      ").Require().Result())
	t.Log(FromString("      ").TrimSpace().Require().Result())
}

// 测试 FromString().Set()
func TestFromStringSet(t *testing.T) {
	var value string
	t.Log(FromString("abc").Set(&value))
	t.Log(value)
}
