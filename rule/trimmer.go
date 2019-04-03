package rule

import "strings"

// 修剪器
type Trimmer func(string) string

// NewTrimmer 创建新修剪器
func NewTrimmer(trimmer Trimmer) Rule {
	return Rule{
		Trim: trimmer,
	}
}

// 去除前后的空格
func trim(value string) string {
	return strings.TrimSpace(value)
}
