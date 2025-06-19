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
		expected *StringType
	}{
		{name: "empty", value: "", expected: &StringType{value: "", name: "empty", err: nil}},
		{name: "大写", value: "testValue", expected: &StringType{value: "TESTVALUE", name: "大写", err: nil}},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value, testCase.name).ToUpper()
			if result.Value() != testCase.expected.value || !errors.Is(result.Error(), testCase.expected.err) {
				t.Errorf("期望：%v, 结果：%v", testCase.expected, result)
			}
		})
	}
}

// 字母转为小写
func TestStringType_ToLower(t *testing.T) {
	cases := []struct {
		name     string
		value    string
		expected *StringType
	}{
		{name: "empty", value: "", expected: &StringType{value: "", name: "empty", err: nil}},
		{name: "ToLower", value: "TestValue", expected: &StringType{value: "testvalue", name: "ToLower", err: nil}},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value, testCase.name).ToLower()
			if result.Value() != testCase.expected.value || !errors.Is(result.Error(), testCase.expected.err) {
				t.Errorf("期望：%v, 结果：%v", testCase.expected, result)
			}
		})
	}
}

// 删除左右的指定字符串
func TestStringType_Trim(t *testing.T) {
	cases := []struct {
		value    string
		expected *StringType
	}{
		{value: "", expected: &StringType{value: "", err: nil}},
		{value: "oTestValue", expected: &StringType{value: "TestValue", err: nil}},
		{value: "oooTestValue", expected: &StringType{value: "TestValue", err: nil}},
		{value: "okoTestValue", expected: &StringType{value: "koTestValue", err: nil}},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value).Trim("o")
			if result.Value() != testCase.expected.value || !errors.Is(result.Error(), testCase.expected.err) {
				t.Errorf("期望：%v, 结果：%v", testCase.expected, result)
			}
		})
	}
}

// 删除所有空格
func TestStringType_RemoveSpace(t *testing.T) {
	cases := []struct {
		value    string
		expected *StringType
	}{
		{value: "", expected: &StringType{value: "", err: nil}},
		{value: "    test      Value        ", expected: &StringType{value: "testValue", err: nil}},
	}

	for _, testCase := range cases {
		t.Run(testCase.value, func(t *testing.T) {
			result := FromString(testCase.value).RemoveSpace()
			if result.Value() != testCase.expected.value || !errors.Is(result.Error(), testCase.expected.err) {
				t.Errorf("期望：%v, 结果：%v", testCase.expected, result)
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
			result := FromString(testCase.value).ToSnakeCase()
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
			result := FromString(testCase.value).ToCamelCase()
			if result.Value() != testCase.expected.value || !errors.Is(result.Error(), testCase.expected.err) {
				t.Errorf("期望：%v, 结果：%v", testCase.expected, result)
			}
		})
	}
}
