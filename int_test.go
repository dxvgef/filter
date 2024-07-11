package filter

import (
	"testing"
)

// 测试 FromInteger()
func TestFromInteger(t *testing.T) {
	t.Log(FromInteger(1, "test"))
}

// 测试用途的自定义整数处理函数
func testCustomIntegerFunc(s *IntegerType) (int64, error) {
	return s.Value() + 1, nil
	// return s.Int64() + 1, errors.New("错误信息")
}

// 测试 FromInteger().Custom()
func TestFromIntegerCustom(t *testing.T) {
	t.Log(FromInteger(1).Custom(testCustomIntegerFunc).Result())
}
