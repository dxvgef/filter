package rule

import (
	"regexp"
	"strings"
)

// 值必须存在于slice
func inSlice(str string, slice []string) bool {
	for k := range slice {
		if slice[k] == str {
			return true
		}
	}
	return false
}

// 必须存在指定字符串
func contains(str string, sub string) bool {
	return strings.Contains(str, sub)
}

// 必须存在指定的前缀字符串
func hasPrefix(str string, sub string) bool {
	return strings.HasPrefix(str, sub)
}

// 是否存在小写字母
func hasLower(str string) bool {
	rule := `([a-z]+?)`
	return regexp.MustCompile(rule).MatchString(str)
}

// 是否存在大写字母
func hasUpper(str string) bool {
	rule := `([A-Z]+?)`
	return regexp.MustCompile(rule).MatchString(str)
}

// 是否存在数字
func hasDigit(str string) bool {
	rule := `(\d+?)`
	return regexp.MustCompile(rule).MatchString(str)
}

// 是否存在特殊字符
func hasSpecialChar(str string) bool {
	rule := `([~!@#$%^&*()_+\-=:;\"'/?<>,.\[\]{}|]+?)`
	return regexp.MustCompile(rule).MatchString(str)
}
