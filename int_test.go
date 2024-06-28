package filter

import (
	"testing"
)

// 测试 FromInteger().Require()
func TestFromIntegerRequire(t *testing.T) {
	err := FromInteger(1, "id").Require().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试用途的自定义整数处理函数
func testCustomIntFunc(s *IntegerType) (int64, error) {
	return s.Int64() + 1, nil
	// return s.Int64() + 1, errors.New("错误信息")
}

// 测试 FromInteger().Custom()
func TestFromIntCustom(t *testing.T) {
	str, err := FromInteger(1).Custom(testCustomIntFunc).Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(str)
}
