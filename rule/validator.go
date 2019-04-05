package rule

import (
	"encoding/json"
	"regexp"
)

// 验证器
type Validator func(string) bool

// Error 为验证器自定义错误消息
func (v *Rule) Error(message string) Rule {
	return NewValidator(v.Validate, message)
}

// NewValidator 创建新验证器
func NewValidator(validator Validator, message string) Rule {
	return Rule{
		Validate: validator,
		Message:  message,
	}
}

// 不能为空
func must(value string) bool {
	if value == "" {
		return false
	}
	return true
}

// 电邮地址
func email(value string) bool {
	rule := `^[a-z0-9]+([\-_\.][a-z0-9]+)*@([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,4}$`
	return regexp.MustCompile(rule).MatchString(value)
}

// ip地址
func ip(str string) bool {
	rule := "^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$"
	return regexp.MustCompile(rule).MatchString(str)
}

// 小写字母
func lower(str string) bool {
	rule := "^[a-z]*$"
	return regexp.MustCompile(rule).MatchString(str)
}

// 大写字母
func upper(str string) bool {
	rule := `^[a-z0-9A-Z\p{Han}]+(_[a-z0-9A-Z\p{Han}]+)*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 大小写字母
func letter(str string) bool {
	rule := `^[a-zA-Z]*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 整数(含0)
func intRule(str string) bool {
	rule := `^(-?[1-9]\d*|0)$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 正整数(含0)
func uintRule(str string) bool {
	rule := `^([1-9]\d*|0)$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 负整数
func nint(str string) bool {
	rule := `^-[1-9]\d*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 小数(含0)
func floatRule(str string) bool {
	if str == "0" {
		return true
	}
	rule := `^(?:[1-9][0-9]*(?:\.[0-9]+)?|0\.[0-9]+)$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 正小数(含0)
func pfloat(str string) bool {
	rule := `^\d+(\.\d+)?$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 负小数
func nfloat(str string) bool {
	rule := `^-([1-9]\d*\.\d*|0\.\d*[1-9]\d*)$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 小写字母和数字
func lowerAndDigit(str string) bool {
	rule := `^[a-z0-9]*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 大写字母和数字
func upperAndDigit(str string) bool {
	rule := `^[A-Z0-9]*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 字母和数字
func letterAndDigit(str string) bool {
	rule := `^[a-zA-Z0-9]*$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 中国大陆地区座机号码
func chinaTel(str string) bool {
	rule := `^(\(\d{3,4}\)|\d{3,4}-)?\d{7,8}$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 中国大陆地区手机号码
func chinaMobile(str string) bool {
	rule := `^1(3|4|5|7|8)\d{9}$`
	return regexp.MustCompile(rule).MatchString(str)
}

// 中文
func chinese(str string) bool {
	rule := `^[\p{Han}]+$`
	return regexp.MustCompile(rule).MatchString(str)
}

// JSON类型
func jsonRule(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// 密码，必须包含大小写字母和数字
func password(str string) bool {
	//rule := `([a-z]+?)([A-Z]+?)(\d+?)`
	rule := `([a-z]+?)`
	result1 := regexp.MustCompile(rule).MatchString(str)
	rule = `([A-Z]+?)`
	result2 := regexp.MustCompile(rule).MatchString(str)
	rule = `(\d+?)`
	result3 := regexp.MustCompile(rule).MatchString(str)
	if !result1 || !result2 || !result3 {
		return false
	}
	return true
}

