package rule

import "unicode/utf8"

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
