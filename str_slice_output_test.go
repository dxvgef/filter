package filter

import "testing"

// 测试 FromString().Value()
func TestFromStringSliceValue(t *testing.T) {
	value := FromStringSlice([]string{"   111   ", "   bbb   "}).TrimSpace().Value()
	t.Log(value)
}

// 测试 FromStringSlice().DefaultValue()
func TestFromStringSliceDefaultValue(t *testing.T) {
	value := FromStringSlice([]string{"      ", "      "}).TrimSpace().Require().DefaultValue([]string{"aa", "bb"})
	t.Log(value)
}

// 测试 FromStringSlice().Error()
func TestFromStringSliceError(t *testing.T) {
	err := FromStringSlice([]string{"      ", "   bb   "}).TrimSpace().Require().Error()
	if err != nil {
		t.Error(err)
	}
}

// 测试 FromStringSlice().Result()
func TestFromStringSliceResult(t *testing.T) {
	value, err := FromStringSlice([]string{"   AA   ", "   bb   "}).TrimSpace().Require().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}
