package filter

import (
	"testing"
)

// 测试 FromIntegerSlice()
func TestFromIntegerSlice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{1, 2}, "test"))
}

// 测试用途的自定义整数切片处理函数
func testCustomIntegerSliceFunc(s *IntegerSliceType) ([]int64, error) {
	if s.Error() != nil {
		return s.Value(), s.Error()
	}
	value := s.Value()
	value = append(value, 3)
	return value, nil
}

// 测试 FromIntegerSlice().Custom()
func TestFromIntegerSliceCustom(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{1, 2}).Custom(testCustomIntegerSliceFunc).Result())
}
