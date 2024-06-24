package filter

import (
	"strings"
)

// ToUpper 字母转为大写
func (strType *StringType) ToUpper() *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.ToUpper(strType.currentValue)
	return strType
}

// ToLower 字母转为小写
func (strType *StringType) ToLower() *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.ToLower(strType.currentValue)
	return strType
}

// Trim 删除左右的指定字符串
func (strType *StringType) Trim(sub string) *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.Trim(strType.currentValue, sub)
	return strType
}

// TrimSpace 删除左右的空格，是 Trim() 函数的封装
func (strType *StringType) TrimSpace() *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.TrimSpace(strType.currentValue)
	return strType
}

// TrimLeft 删除左边指定的字符串
func (strType *StringType) TrimLeft(sub string) *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.TrimLeft(strType.currentValue, sub)
	return strType
}

// TrimRight 删除右边指定的字符串
func (strType *StringType) TrimRight(sub string) *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.TrimRight(strType.currentValue, sub)
	return strType
}

// TrimPrefix 删除指定的前缀字符串
func (strType *StringType) TrimPrefix(sub string) *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.TrimPrefix(strType.currentValue, sub)
	return strType
}

// TrimSuffix 删除指定的后缀字符串
func (strType *StringType) TrimSuffix(sub string) *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.TrimSuffix(strType.currentValue, sub)
	return strType
}

// Replace 替换指定的字符串，可指定替换次数
func (strType *StringType) Replace(old, new string, n int) *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.Replace(strType.currentValue, old, new, n)
	return strType
}

// ReplaceAll 替换所有指定的字符串
func (strType *StringType) ReplaceAll(old, new string) *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.ReplaceAll(strType.currentValue, old, new)
	return strType
}

// RemoveSpace 删除所有空格
func (strType *StringType) RemoveSpace() *StringType {
	if strType.err != nil || strType.currentValue == "" {
		return strType
	}
	strType.currentValue = strings.ReplaceAll(strType.currentValue, " ", "")
	return strType
}
