package filter

import "testing"

// 测试 FromString()
func TestFromString(t *testing.T) {
	t.Log(FromString("abc", "test"))
}

// 测试用途的自定义string处理函数
func testCustomStringFunc(s *StringType) (string, error) {
	if s.Error() != nil {
		return s.Value(), s.Error()
	}
	return s.Value() + "def", nil
}

// 测试 FromString().Custom()
func TestFromStringCustom(t *testing.T) {
	str, err := FromString("abc").Custom(testCustomStringFunc).Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(str)
}
