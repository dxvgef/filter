package rule

import (
	"strings"
	"unicode"
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
	for _, v := range str {
		if unicode.IsLower(v) == true {
			return true
		}
	}
	return false
}

// 是否存在大写字母
func hasUpper(str string) bool {
	for _, v := range str {
		if unicode.IsUpper(v) == true {
			return true
		}
	}
	return false
}

// 是否存在数字
func hasDigit(str string) bool {
	for _, v := range str {
		if unicode.IsDigit(v) == true {
			return true
		}
	}
	return false
}

// 是否存在除字母、数字、汉字之外的其它特殊字符
func hasSpecialChar(str string) bool {
	for _, v := range str {
		if unicode.IsDigit(v) == false || unicode.IsLetter(v) == false || unicode.Is(unicode.Han, v) == false {
			return true
		}
	}
	return false
}
