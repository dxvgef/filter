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
		{"", &StringType{value: "", err: nil}},
		{"testValue", &StringType{value: "testValue", err: nil}},
		{"testValue", &StringType{value: "TestValue", err: expectedErr}},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value).Equals(testCase.expected.value)
			if (result.err != nil && testCase.expected.err == nil) || (result.err == nil && testCase.expected.err != nil) {
				t.Errorf("期望：%v, 结果：%v", testCase.expected.err, result.err)
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
			result := FromString(testCase.value).NotEquals(testCase.input)
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
		{"", "", nil},
		{"", "a", expectedErr},
		{"testValue", "va", expectedErr},
		{"testValue", "Va", nil},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value).Contains(testCase.input)
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
		{"", "a", nil},
		{"testValue", "va", nil},
		{"testValue", "Va", expectedErr},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value).NotContains(testCase.input)
			if (result.err != nil && testCase.err == nil) || (result.err == nil && testCase.err != nil) {
				t.Errorf("期望：%v, 结果：%v", testCase.err, result.err)
			}
		})
	}
}
