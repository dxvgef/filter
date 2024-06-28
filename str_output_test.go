package filter

import "testing"

// 测试 FromString().Result()
func TestFromStringResult(t *testing.T) {
	result, err := FromString("  ok  ", "username").TrimSpace().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

// 测试 FromString().Int64()
func TestFromStringValue(t *testing.T) {
	result := FromString("  ok  ", "username").TrimSpace().Value()
	t.Log(result)
}

// 测试 FromString().DefaultValue()
func TestFromStringDefaultValue(t *testing.T) {
	result := FromString("    ", "username").TrimSpace().Require().DefaultValue("ok")
	t.Log(result)
}

// 测试 FromString().Error()
func TestFromStringError(t *testing.T) {
	err := FromString("    ", "username").TrimSpace().Require().Error()
	t.Error(err)
}

// 测试 FromString().Set()
func TestFromStringSet(t *testing.T) {
	var username string
	err := FromString("  ok  ", "username").TrimSpace().Require().Set(&username)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(username)
}
