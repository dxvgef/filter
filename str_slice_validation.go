package filter

import (
	"net"
	"net/mail"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Require 不能为零值
func (strSliceType *StringSliceType) Require(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}
	if len(strSliceType.value) == 0 {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	if strSliceType.value[0] == "" {
		strSliceType.err = wrapError(strSliceType.name, customError...)
	}
	return strSliceType
}

// MinCount 切片元素的数量不能小于指定值
func (strSliceType *StringSliceType) MinCount(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}
	if len(strSliceType.value) < value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	return strSliceType
}

// MaxCount 切片元素的数量不能大于指定值
func (strSliceType *StringSliceType) MaxCount(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}
	if len(strSliceType.value) > value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	return strSliceType
}

// EqualCount 切片元素的数量必须是指定值
func (strSliceType *StringSliceType) EqualCount(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}
	if len(strSliceType.value) != value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	return strSliceType
}

// NotEqualCount 切片元素的数量不能是指定值
func (strSliceType *StringSliceType) NotEqualCount(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}
	if len(strSliceType.value) == value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	return strSliceType
}

// Length 指定的长度
func (strSliceType *StringSliceType) Length(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if len(strSliceType.value[k]) != value {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// UTF8Length 按UTF8编码指定长度
func (strSliceType *StringSliceType) UTF8Length(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if utf8.RuneCountInString(strSliceType.value[k]) != value {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// MinLength 最小长度
func (strSliceType *StringSliceType) MinLength(min int, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if len(strSliceType.value[k]) < min {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// UTF8MinLength 按UTF8编码校验最小长度
func (strSliceType *StringSliceType) UTF8MinLength(min int, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if utf8.RuneCountInString(strSliceType.value[k]) < min {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// MaxLength 最大长度
func (strSliceType *StringSliceType) MaxLength(max int, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if len(strSliceType.value[k]) > max {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// UTF8MaxLength 按UTF8编码校验最大长度
func (strSliceType *StringSliceType) UTF8MaxLength(max int, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if utf8.RuneCountInString(strSliceType.value[k]) > max {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// IsLower 小写字母
func (strSliceType *StringSliceType) IsLower(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		for _, v := range strSliceType.value[k] {
			if !unicode.IsLower(v) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// IsUpper 大写字母
func (strSliceType *StringSliceType) IsUpper(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, v := range strSliceType.value[k] {
			if !unicode.IsUpper(v) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// IsLetter 大小写字母
func (strSliceType *StringSliceType) IsLetter(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, v := range strSliceType.value[k] {
			if !unicode.IsLetter(v) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// IsLowerOrNumber 小写字母或数字
func (strSliceType *StringSliceType) IsLowerOrNumber(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		for _, v := range strSliceType.value[k] {
			if !unicode.IsLower(v) && !unicode.IsDigit(v) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// IsUpperOrNumber 大写字母或数字
func (strSliceType *StringSliceType) IsUpperOrNumber(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		for _, v := range strSliceType.value[k] {
			if !unicode.IsUpper(v) && !unicode.IsDigit(v) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// IsLetterOrNumber 字母或数字
func (strSliceType *StringSliceType) IsLetterOrNumber(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		for _, v := range strSliceType.value[k] {
			if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// IsChinese 汉字
func (strSliceType *StringSliceType) IsChinese(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		for _, v := range strSliceType.value[k] {
			if !unicode.Is(unicode.Scripts["Han"], v) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// IsMail 电邮地址
func (strSliceType *StringSliceType) IsMail(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		if _, err := mail.ParseAddress(strSliceType.value[k]); err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsIP IPv4/v6地址
func (strSliceType *StringSliceType) IsIP(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		if net.ParseIP(strSliceType.value[k]) == nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsTCPAddr IP:Port
func (strSliceType *StringSliceType) IsTCPAddr(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		if _, err := net.ResolveTCPAddr("tcp", strSliceType.value[k]); err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsMAC MAC地址
func (strSliceType *StringSliceType) IsMAC(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		if _, err := net.ParseMAC(strSliceType.value[k]); err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsSQLObject SQL对象（库名、表名、字段名）
func (strSliceType *StringSliceType) IsSQLObject(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		if !isSQLObject(strSliceType.value[k]) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsUUID UUID格式
func (strSliceType *StringSliceType) IsUUID(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		if !isUUID(strSliceType.value[k]) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsULID ULID格式
func (strSliceType *StringSliceType) IsULID(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if len(strSliceType.value[k]) != 26 {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
		validChars := "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
		for _, char := range strSliceType.value[k] {
			if !strings.ContainsRune(validChars, unicode.ToUpper(char)) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}

	return strSliceType
}

// AllowedChars 仅允许字符串中包启allowValues内的字符
func (strSliceType *StringSliceType) AllowedChars(allowValues []rune, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, r := range strSliceType.value[k] {
			if !slices.Contains(allowValues, r) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}

	return strSliceType
}

// AllowedSymbols 如果有符号，只允许存在指定的符号
func (strSliceType *StringSliceType) AllowedSymbols(allowValues []rune, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, r := range strSliceType.value[k] {
			// 如果是符号
			if unicode.IsPunct(r) {
				if !slices.Contains(allowValues, r) {
					strSliceType.err = wrapError(strSliceType.name, customError...)
					return strSliceType
				}
			}
		}
	}

	return strSliceType
}

// AllowedStrings 只允许存在指定的字符串（枚举）
func (strSliceType *StringSliceType) AllowedStrings(allowedValues []string, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if !slices.Contains(allowedValues, strSliceType.value[k]) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// DeniedChars 阻止deniedValues中的值
func (strSliceType *StringSliceType) DeniedChars(deniedValues []rune, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, r := range strSliceType.value[k] {
			if slices.Contains(deniedValues, r) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// DeniedSymbols 如果有符号，阻止存在指定的符号
func (strSliceType *StringSliceType) DeniedSymbols(deniedValues []rune, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, r := range strSliceType.value[k] {
			// 如果是符号
			if unicode.IsPunct(r) {
				if slices.Contains(deniedValues, r) {
					strSliceType.err = wrapError(strSliceType.name, customError...)
					return strSliceType
				}
			}
		}
	}
	return strSliceType
}

// DeniedStrings 禁止存在指定的字符串
func (strSliceType *StringSliceType) DeniedStrings(deniedValues []string, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		for i := range deniedValues {
			if deniedValues[i] == strSliceType.value[k] {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// 包含了字母(不区分大小写)
func (strSliceType *StringSliceType) HasLetter(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		has := false
		for _, v := range strSliceType.value[k] {
			if unicode.IsLetter(v) {
				has = true
				break // 停止遍历当前字符串元素
			}
		}
		if !has {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
	}
	return strSliceType
}

// HasLower 包含了小写字母
func (strSliceType *StringSliceType) HasLower(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		has := false
		for _, v := range strSliceType.value[k] {
			if unicode.IsLower(v) {
				has = true
				break // 停止遍历当前字符串元素
			}
		}
		if !has {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
	}
	return strSliceType
}

// HasUpper 包含了大写字母
func (strSliceType *StringSliceType) HasUpper(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		has := false
		for _, v := range strSliceType.value[k] {
			if unicode.IsUpper(v) {
				has = true
				break // 停止遍历当前字符串元素
			}
		}
		if !has {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
	}
	return strSliceType
}

// HasNumber 包含了数字
func (strSliceType *StringSliceType) HasNumber(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		has := false
		for _, v := range strSliceType.value[k] {
			if unicode.IsDigit(v) {
				has = true
				break // 停止遍历当前字符串元素
			}
		}
		if !has {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
	}
	return strSliceType
}

// HasSymbol 包含了符号
func (strSliceType *StringSliceType) HasSymbol(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		has := false
		for _, v := range strSliceType.value[k] {
			if !unicode.IsDigit(v) && !unicode.IsLetter(v) && !unicode.Is(unicode.Han, v) {
				has = true
				break // 停止遍历当前字符串元素
			}
		}
		if !has {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
	}
	return strSliceType
}

// Contains 必须包含指定的字符串
func (strSliceType *StringSliceType) Contains(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		if !strings.Contains(strSliceType.value[k], sub) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// HasPrefix 包含了指定的前缀字符串
func (strSliceType *StringSliceType) HasPrefix(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if !strings.HasPrefix(strSliceType.value[k], sub) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// HasSuffix 包含了指定的后缀字符串
func (strSliceType *StringSliceType) HasSuffix(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}

	for k := range strSliceType.value {
		if !strings.HasSuffix(strSliceType.value[k], sub) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}
