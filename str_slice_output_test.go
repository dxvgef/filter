package filter

import "testing"

// 测试 FromString().Value()
func TestFromStringSliceValue(t *testing.T) {
	t.Log(FromStringSlice([]string{"   111   ", "   bbb   "}).TrimSpace().Value())
}

// 测试 FromStringSlice().DefaultValue()
func TestFromStringSliceDefaultValue(t *testing.T) {
	t.Log(FromStringSlice([]string{"      ", "      "}).TrimSpace().Require().DefaultValue([]string{"aa", "bb"}))
}

// 测试 FromStringSlice().Error()
func TestFromStringSliceError(t *testing.T) {
	t.Log(FromStringSlice([]string{"      ", "   bb   "}).TrimSpace().Require().Error())
}

// 测试 FromStringSlice().Result()
func TestFromStringSliceResult(t *testing.T) {
	t.Log(FromStringSlice([]string{"   AA   ", "   bb   "}).TrimSpace().Require().Result())
}

// 测试 FromStringSlice().Set()
func TestFromStringSliceSet(t *testing.T) {
	var value []string
	t.Log(FromStringSlice([]string{"AA", "bb"}).Set(&value))
	t.Log(value)
}
