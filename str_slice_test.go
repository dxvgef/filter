package filter

import "testing"

// 测试 FromStringSlice()
func TestFromStringSlice(t *testing.T) {
	t.Log(FromStringSlice([]string{"aa", "bb"}, "test"))
}

// 测试用途的自定义[]string处理函数
func testCustomStringSliceFunc(s *StringSliceType) ([]string, error) {
	if s.Error() != nil {
		return s.Value(), s.Error()
	}
	value := s.Value()
	value = append(value, "cc")
	return value, nil
}

// 测试 FromStringSlice().Custom()
func TestFromStringSliceCustom(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).Custom(testCustomStringSliceFunc).Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}
