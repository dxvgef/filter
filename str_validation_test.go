package filter

import (
	"errors"
	"testing"
)

// 等于
func TestStringType_Equals(t *testing.T) {
	expectedErr := errors.New("未通过")
	cases := []struct {
		value    string
		expected *StringType
	}{
		{"", &StringType{value: "", err: expectedErr}},
		{"testValue", &StringType{value: "testValue", err: nil}},
		{"testValue", &StringType{value: "TestValue", err: expectedErr}},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value).Is(testCase.expected.value)
			if (result.err != nil && testCase.expected.err == nil) || (result.err == nil && testCase.expected.err != nil) {
				t.Errorf("输入: %s, 期望错误：%v, 实际错误：%v", testCase.value, testCase.expected.err, result.err)
			}
		})
	}
}

func TestStringType_NotEquals(t *testing.T) {
	expectedErr := errors.New("未通过")
	cases := []struct {
		value string
		input string
		err   error
	}{
		{"", "", expectedErr},
		{"testValue", "testValue", expectedErr},
		{"testValue", "TestValue", nil},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value).IsNot(testCase.input)
			if (result.err != nil && testCase.err == nil) || (result.err == nil && testCase.err != nil) {
				t.Errorf("期望：%v, 结果：%v", testCase.err, result.err)
			}
		})
	}
}

func TestStringType_Contains(t *testing.T) {
	expectedErr := errors.New("未通过")
	cases := []struct {
		value string
		input string
		err   error
	}{
		{"", "", expectedErr},
		{"", "a", expectedErr},
		{"testValue", "va", expectedErr},
		{"testValue", "Va", nil},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value).Has(testCase.input)
			if (result.err != nil && testCase.err == nil) || (result.err == nil && testCase.err != nil) {
				t.Errorf("期望：%v, 结果：%v", testCase.err, result.err)
			}
		})
	}
}

func TestStringType_NotContains(t *testing.T) {
	expectedErr := errors.New("未通过")
	cases := []struct {
		value string
		input string
		err   error
	}{
		{"", "", expectedErr},
		{"", "a", expectedErr},
		{"testValue", "va", nil},
		{"testValue", "Va", expectedErr},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value).HasNot(testCase.input)
			if (result.err != nil && testCase.err == nil) || (result.err == nil && testCase.err != nil) {
				t.Errorf("期望：%v, 结果：%v", testCase.err, result.err)
			}
		})
	}
}
