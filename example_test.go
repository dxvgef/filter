package filter

import (
	"testing"

	"github.com/dxvgef/filter/rule"
	"strings"
)

// TestValidation 测试普通验证器
func TestValidation(t *testing.T) {
	value := "dxvgef@outlook.com"
	err := Result(value, rule.Email)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestLengthValidation 测试长度验证器
func TestLengthValidation(t *testing.T) {
	value := "ab"
	err := Result(value,
		rule.MinLength(3),
		rule.MaxLength(8),
	)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestIntRangeValidation 测试整数值范围验证器
func TestIntRangeValidation(t *testing.T) {
	value := "9"
	err := Result(value,
		rule.MinInteger(3),
		rule.MaxInteger(8),
	)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestFloatRangeValidation 测试浮点数值范围验证器
func TestFloatRangeValidation(t *testing.T) {
	value := "3.4"
	err := Result(value,
		rule.MinFloat(3.5),
		rule.MaxFloat(8.5),
	)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestTrim 测试修剪器
func TestTrim(t *testing.T) {
	value := "    127.0.0.1      "
	err := Result(value, rule.Trim, rule.IP.Error("请填写正确的IP地址"))
	if err != nil {
		t.Log(err.Error())
	}
}

// TestCustomValidation 测试自定义验证器
func TestCustomValidation(t *testing.T) {
	value := "    127.0.0.1      "
	var validator = func(value string) bool {
		return true
	}
	customRule := rule.NewValidator(validator, "默认消息")
	err := Result(value, rule.Trim, customRule)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestCustomTrimmer 测试自定义修剪器
func TestCustomTrimmer(t *testing.T) {
	value := "abc"
	var trimmer = func(value string) string {
		return value[0:1]
	}
	customRule := rule.NewTrimmer(trimmer)
	err := Result(value, rule.Trim, customRule)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestHasLower 测试含有至少有一位小写字母
func TestHasLower(t *testing.T) {
	value := "outlook.com"
	err := Result(value, rule.HasLower)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestHasUpper 测试含有至少有一位大写字母
func TestHasUpper(t *testing.T) {
	value := "outLook"
	err := Result(value, rule.HasUpper)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestHasDigit 测试含有至少有一位数字
func TestHasDigit(t *testing.T) {
	value := "outLook2"
	err := Result(value, rule.HasDigit)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestHasSpecialLetter 测试含有至少有一个特殊字符
func TestHasSpecialChar(t *testing.T) {
	value := "outLook.com"
	err := Result(value, rule.HasSpecialChar)
	if err != nil {
		t.Log(err.Error())
	}
}

// TestStringConverterOpt 测试字符日转换优化器
func TestStringConverterOpt(t *testing.T) {
	value := "Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http"
	err := Result(value, rule.NewStringOpt(rule.TrimStringAndSnake))
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(value) //todo 结果值没变,bug
}


// TestStringValueInByStringValidate 测试包含字符串
func TestStringValueInByStringValidate(t *testing.T){
	value := "Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http"
	split := strings.Split(value, " ")
	err := Result("Go,", rule.StringInValidate(split))
	if err != nil {
		t.Log(err.Error())
	}

}