package filter

import (
	"testing"
)

// 测试 FromFloatSlice()
func TestFromFloatSlice(t *testing.T) {
	t.Log(FromFloatSlice([]float64{1, 2}, "test"))
}

// 测试用途的自定义整数切片处理函数
func testCustomFloatSliceFunc(s *FloatSliceType) ([]float64, error) {
	if s.Error() != nil {
		return s.Value(), s.Error()
	}
	value := s.Value()
	value = append(value, 3)
	return value, nil
}

// 测试 FromFloatSlice().Custom()
func TestFromFloatSliceCustom(t *testing.T) {
	t.Log(FromFloatSlice([]float64{1, 2}).Custom(testCustomFloatSliceFunc).Result())
}
