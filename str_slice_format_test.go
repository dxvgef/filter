package filter

import (
	"errors"
	"slices"
	"testing"
)

// 转为蛇形命名
func TestStringSliceType_ToSnakeCase(t *testing.T) {
	cases := []struct {
		value    []string
		expected *StringSliceType
	}{
		{value: []string{"aaBB", "aaBb"}, expected: &StringSliceType{value: []string{"aa_bb", "aa_bb"}, err: nil}},
	}

	for _, testCase := range cases {
		t.Run("TestStringSliceType_ToSnakeCase", func(t *testing.T) {
			result := FromStringSlice(testCase.value).ToSnakeCase()
			if !slices.Equal(result.Value(), testCase.expected.value) || !errors.Is(result.Error(), testCase.expected.err) {
				t.Errorf("期望：%v, 结果：%v", testCase.expected, result)
			}
		})
	}
}
