package filter

import (
	"errors"
	"testing"
)

func TestFromString(t *testing.T) {
	// 测试从字符串初始化
	cases := []struct {
		input    string
		name     string
		expected *StringType
	}{
		{"testValue", "", &StringType{value: "testValue", name: ""}},
		{"testValue", "testName", &StringType{value: "testValue", name: "testName"}},
	}

	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			result := FromString(tt.input, tt.name)
			if result.value != tt.expected.value || result.name != tt.expected.name {
				t.Errorf("期望：%v, 结果：%v", tt.expected, result)
			}
		})
	}
}

func TestCustom(t *testing.T) {
	// 自定义函数，处理字符串
	customFunc := func(strType *StringType) (string, error) {
		return strType.value + "_modified", nil
	}

	// 测试正常情况
	strType := FromString("testValue")
	result := strType.Custom(customFunc)
	if result.err != nil {
		t.Errorf("期望没有错误, 错误：%v", result.err)
	}
	if result.value != "testValue_modified" {
		t.Errorf("期望值：'testValue_modified', 结果值：'%v'", result.value)
	}

	// 测试自定义函数返回错误
	customFuncWithError := func(strType *StringType) (string, error) {
		return strType.value + "_modified", errors.New("testError")
	}
	strTypeWithError := FromString("testValue")
	resultWithError := strTypeWithError.Custom(customFuncWithError)
	if resultWithError.err == nil {
		t.Errorf("期望有错误，但结果为 nil")
	}
	if resultWithError.value != "testValue" {
		t.Errorf("期望值：'testValue'，结果：'%v'", resultWithError.value)
	}
	if resultWithError.Error().Error() != "testError" {
		t.Errorf("期望值：'testError'，结果：'%v'", resultWithError.Error().Error())
	}
}
