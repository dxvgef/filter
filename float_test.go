package filter

import (
	"testing"
)

// 测试 FromFloat()
func TestFromFloat(t *testing.T) {
	t.Log(FromFloat(1, "test"))
}

// 测试用途的自定义整数处理函数
func testCustomFloatFunc(s *FloatType) (float64, error) {
	return s.Value() + 1, nil
	// return s.Value() + 1, errors.New("错误信息")
}

// 测试 FromFloat().Custom()
func TestFromFloatCustom(t *testing.T) {
	t.Log(FromFloat(1).Custom(testCustomFloatFunc).Result())
}
