package rule

import "strings"

// inValidator 值必须存在于slice的验证器
type inValidator func(string, []string) bool

// NewInValidate 创建值必须存在于slice的验证器
func NewInValidate(slice []string, validator func(string, []string) bool, message string) Rule {
	return Rule{
		Message:    message,
		InValidate: validator,
		InValues:   slice,
	}
}

// 值必须存在于slice
func inSlice(str string, slice []string) bool {
	for k := range slice {
		if slice[k] == str {
			return true
		}
	}
	return false
}

// ------------------------

// ContainValidator 必须存在指定字符串的验证器
type ContainValidator func(string, string) bool

// NewContainValidate 创建必须存在指定字符串的验证器
func NewContainValidate(sub string, validator ContainValidator, message string) Rule {
	return Rule{
		Message:         message,
		ContainValidate: validator,
		ContainString:   sub,
	}
}

// 必须存在指定字符串
func contains(str string, sub string) bool {
	return strings.Contains(str, sub)
}
