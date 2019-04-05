package rule

import "unicode/utf8"

// 长度验证器
type LengthValidator func(string, int) bool

// NewMinLengthValidator 创建字符串最小长度验证器
func NewMinLengthValidator(min int, validator LengthValidator, message string) Rule {
	return Rule{
		Message:        message,
		LengthValidate: validator,
		MinLength:      min,
	}
}

// NewMaxLengthValidator 创建字符串最大长度验证器
func NewMaxLengthValidator(max int, validator LengthValidator, message string) Rule {
	return Rule{
		Message:        message,
		LengthValidate: validator,
		MaxLength:      max,
	}
}

// 验证字符串最小长度
func minLength(value string, min int) bool {
	l := utf8.RuneCountInString(value)
	if l < min {
		return false
	}
	return true
}

// 验证字符串最大长度
func maxLength(value string, max int) bool {
	l := utf8.RuneCountInString(value)
	if l > max {
		return false
	}
	return true
}
