package filter

import (
	"testing"
)

// 测试 FromBoolean()
func TestFromBoolean(t *testing.T) {
	t.Log(FromBoolean(false, "test"))
}

// 测试用途的自定义布尔处理函数
func testCustomBooleanFunc(s *BooleanType) (bool, error) {
	return !s.Value(), nil
	// return s.Int64() + 1, errors.New("错误信息")
}

// 测试 FromBoolean().Custom()
func TestFromBooleanCustom(t *testing.T) {
	t.Log(FromBoolean(false).Custom(testCustomBooleanFunc).Result())
}
