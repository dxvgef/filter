package filter

import "testing"

// 测试 FromString().ToInteger()
func TestFromStringToInteger(t *testing.T) {
	value, err := FromString("   111   ", "id").TrimSpace().Require().ToInteger().Require().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}
