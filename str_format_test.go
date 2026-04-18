package filter

import (
	"errors"
	"testing"
)

// 字母转为大写
func TestStringType_ToUpper(t *testing.T) {
	cases := []struct {
		name     string
		value    string
		hasError bool
		expected string
	}{
		{name: "empty", value: "", hasError: true, expected: ""},
		{name: "大写", value: "testValue", hasError: false, expected: "TESTVALUE"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			result := FromStr(testCase.value, testCase.name).ToUpper()
			if (result.Error() != nil) != testCase.hasError {
				t.Errorf("期望错误：%v, 实际结果：%v", testCase.hasError, result.Error())
			}
			if !testCase.hasError && result.Value() != testCase.expected {
				t.Errorf("期望值：%s, 实际值：%s", testCase.expected, result.Value())
			}
		})
	}
}

// 字母转为小写
func TestStringType_ToLower(t *testing.T) {
	cases := []struct {
		name     string
		value    string
		hasError bool
		expected string
	}{
		{name: "empty", value: "", hasError: true, expected: ""},
		{name: "ToLower", value: "TestValue", hasError: false, expected: "testvalue"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			result := FromStr(testCase.value, testCase.name).ToLower()
			if (result.Error() != nil) != testCase.hasError {
				t.Errorf("期望错误：%v, 实际结果：%v", testCase.hasError, result.Error())
			}
			if !testCase.hasError && result.Value() != testCase.expected {
				t.Errorf("期望值：%s, 实际值：%s", testCase.expected, result.Value())
			}
		})
	}
}

// 删除左右的指定字符串
func TestStringType_Trim(t *testing.T) {
	cases := []struct {
		value    string
		hasError bool
		expected string
	}{
		{value: "", hasError: true, expected: ""},
		{value: "oTestValue", hasError: false, expected: "TestValue"},
		{value: "oooTestValue", hasError: false, expected: "TestValue"},
		{value: "okoTestValue", hasError: false, expected: "koTestValue"},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromStr(testCase.value).Trim("o")
			if (result.Error() != nil) != testCase.hasError {
				t.Errorf("期望错误：%v, 实际结果：%v", testCase.hasError, result.Error())
			}
			if !testCase.hasError && result.Value() != testCase.expected {
				t.Errorf("期望值：%s, 实际值：%s", testCase.expected, result.Value())
			}
		})
	}
}

// 删除所有空格
func TestStringType_RemoveSpace(t *testing.T) {
	cases := []struct {
		value    string
		hasError bool
		expected string
	}{
		{value: "", hasError: true, expected: ""},
		{value: "    test      Value        ", hasError: false, expected: "testValue"},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromStr(testCase.value).RemoveSpace()
			if (result.Error() != nil) != testCase.hasError {
				t.Errorf("期望错误：%v, 实际结果：%v", testCase.hasError, result.Error())
			}
			if !testCase.hasError && result.Value() != testCase.expected {
				t.Errorf("期望值：%s, 实际值：%s", testCase.expected, result.Value())
			}
		})
	}
}

// 转为蛇形命名
func TestStringType_ToSnakeCase(t *testing.T) {
	cases := []struct {
		value    string
		expected *StringType
	}{
		{value: "aaBB", expected: &StringType{value: "aa_bb", err: nil}},
		{value: "aaBb", expected: &StringType{value: "aa_bb", err: nil}},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromStr(testCase.value).ToSnakeCase()
			if result.Value() != testCase.expected.value || !errors.Is(result.Error(), testCase.expected.err) {
				t.Errorf("期望：%v, 结果：%v", testCase.expected, result)
			}
		})
	}
}

// 转为驼峰命名
func TestStringType_ToCamelCase(t *testing.T) {
	cases := []struct {
		value    string
		expected *StringType
	}{
		{value: "aa_BB", expected: &StringType{value: "aaBB", err: nil}},
		{value: "aa_Bb", expected: &StringType{value: "aaBb", err: nil}},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromStr(testCase.value).ToCamelCase()
			if result.Value() != testCase.expected.value || !errors.Is(result.Error(), testCase.expected.err) {
				t.Errorf("期望：%v, 结果：%v", testCase.expected, result)
			}
		})
	}
}
