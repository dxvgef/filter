package filter

import (
	"testing"
)

// 测试 FromBooleanSlice()
func TestFromBooleanSlice(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, false}, "test"))
}

// 测试用途的自定义[]bool处理函数
func testCustomBooleanSliceFunc(s *BooleanSliceType) ([]bool, error) {
	value := s.Value()
	value = append(value, true)
	return value, nil
}

// 测试 FromBooleanSlice().Custom()
func TestFromBooleanSliceCustom(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, false}).Custom(testCustomBooleanSliceFunc).Result())
}
