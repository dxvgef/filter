package filter

import "testing"

// 测试 FromString().Require()
func TestFromStringRequire(t *testing.T) {
	err := FromString("aa", "username").TrimSpace().Require().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试用途的自定义string处理函数
func testCustomStringFunc(s *StringType) (string, error) {
	return s.Value() + "def", nil
	// return s.Int64() + "def", errors.New("错误信息")
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
