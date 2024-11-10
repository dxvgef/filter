package filter

import (
	"net"
	"net/mail"
	"net/url"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Contains 每个元素值都包含了指定的字符串
func (strSliceType *StringSliceType) AllContains(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// Contains 每个元素值都没有包含指定的字符串
func (strSliceType *StringSliceType) AllNotContains(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if strings.Contains(strSliceType.value[k], sub) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// AnyContains 有元素值包含了指定的字符串
func (strSliceType *StringSliceType) AnyContains(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if strings.Contains(strSliceType.value[k], sub) {
			return strSliceType
		}
	}
	strSliceType.err = wrapError(strSliceType.name, customError...)
	return strSliceType
}

// AnyNotContains 有元素值没有包含指定的字符串
func (strSliceType *StringSliceType) AnyNotContains(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// AllowedValues 元素值只能有数组中的值
func (strSliceType *StringSliceType) AllowedValues(values []string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if !slices.Contains(values, strSliceType.value[k]) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// DisallowedValues 元素值不能有数组中的值
func (strSliceType *StringSliceType) DisallowedValues(values []string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		for i := range values {
			if values[i] == strSliceType.value[k] {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// AllowedChars 元素值只能有数组中的字符
func (strSliceType *StringSliceType) AllowedChars(values []rune, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, r := range strSliceType.value[k] {
			if !slices.Contains(values, r) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}

	return strSliceType
}

// DisallowedChars 元素值不能有数组中的字符
func (strSliceType *StringSliceType) DisallowedChars(values []rune, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, r := range strSliceType.value[k] {
			if slices.Contains(values, r) {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}
	return strSliceType
}

// AllowedSymbols 元素值如果有符号，只能有数组中的符号
func (strSliceType *StringSliceType) AllowedSymbols(values []rune, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, r := range strSliceType.value[k] {
			// 如果是符号
			if unicode.IsPunct(r) {
				if !slices.Contains(values, r) {
					strSliceType.err = wrapError(strSliceType.name, customError...)
					return strSliceType
				}
			}
		}
	}

	return strSliceType
}

// DisallowedSymbols 元素值如果有符号，不能有数组中的符号
func (strSliceType *StringSliceType) DisallowedSymbols(values []rune, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		for _, r := range strSliceType.value[k] {
			// 如果是符号
			if unicode.IsPunct(r) {
				if slices.Contains(values, r) {
					strSliceType.err = wrapError(strSliceType.name, customError...)
					return strSliceType
				}
			}
		}
	}
	return strSliceType
}

// HasLetter 每个元素值都包含了字母(不区分大小写)
func (strSliceType *StringSliceType) HasLetter(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// HasLower 每个元素值都包含了小写字母
func (strSliceType *StringSliceType) HasLower(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// HasUpper 每个元素值都包含了大写字母
func (strSliceType *StringSliceType) HasUpper(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// HasNumber 每个元素值都包含了数字
func (strSliceType *StringSliceType) HasNumber(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// HasSymbol 每个元素值都包含了符号
func (strSliceType *StringSliceType) HasSymbol(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// HasPrefix 每个元素值都包含了指定的前缀字符串
func (strSliceType *StringSliceType) HasPrefix(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// HasSuffix 每个元素值都包含了指定的后缀字符串
func (strSliceType *StringSliceType) HasSuffix(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// CountLessThan 元素的数量小于
func (strSliceType *StringSliceType) CountLessThan(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	if len(strSliceType.value) > value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
	}
	return strSliceType
}

// CountGreaterThan 元素的数量大于
func (strSliceType *StringSliceType) CountGreaterThan(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	if len(strSliceType.value) > value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	return strSliceType
}

// CountEquals 元素的数量等于
func (strSliceType *StringSliceType) CountEquals(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	if len(strSliceType.value) != value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	return strSliceType
}

// CountNotEquals 元素的数量不等于
func (strSliceType *StringSliceType) CountNotEquals(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	if len(strSliceType.value) == value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	return strSliceType
}

// LengthEquals 每个元素值的长度都等于
func (strSliceType *StringSliceType) LengthEquals(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// LengthNotEquals 每个元素值的长度都不等于
func (strSliceType *StringSliceType) LengthNotEquals(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if len(strSliceType.value[k]) == value {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// LengthLessThan 每个元素值的长度都小于
func (strSliceType *StringSliceType) LengthLessThan(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if len(strSliceType.value[k]) > value {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// LengthGreaterThan 每个元素值的长度都大于
func (strSliceType *StringSliceType) LengthGreaterThan(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if len(strSliceType.value[k]) < value {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// UTF8LengthEquals 每个元素值的UTF8编码的长度都等于
func (strSliceType *StringSliceType) UTF8LengthEquals(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// UTF8LengthNotEquals 每个元素值的UTF8编码的长度都不等于
func (strSliceType *StringSliceType) UTF8LengthNotEquals(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if utf8.RuneCountInString(strSliceType.value[k]) == value {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// UTF8LengthLessThan 每个元素值的UTF8编码的长度都小于
func (strSliceType *StringSliceType) UTF8LengthLessThan(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if utf8.RuneCountInString(strSliceType.value[k]) > value {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// UTF8LengthGreaterThan 每个元素的UTF8编码的长度都大于
func (strSliceType *StringSliceType) UTF8LengthGreaterThan(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if utf8.RuneCountInString(strSliceType.value[k]) < value {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// IsLower 每个元素值都是小写字母
func (strSliceType *StringSliceType) IsLower(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsUpper 每个元素值都是大写字母
func (strSliceType *StringSliceType) IsUpper(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsLetter 每个元素值都是大小写字母
func (strSliceType *StringSliceType) IsLetter(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsLowerOrNumber 每个元素值都是小写字母或数字
func (strSliceType *StringSliceType) IsLowerOrNumber(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsUpperOrNumber 每个元素值都是大写字母或数字
func (strSliceType *StringSliceType) IsUpperOrNumber(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsLetterOrNumber 每个元素值都是字母或数字
func (strSliceType *StringSliceType) IsLetterOrNumber(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsChinese 每个元素值都是汉字
func (strSliceType *StringSliceType) IsChinese(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsMail 每个元素值都是电邮地址
func (strSliceType *StringSliceType) IsMail(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsIPv4 每个元素值都是IPv4地址
func (strSliceType *StringSliceType) IsIPv4(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		ip := net.ParseIP(strSliceType.value[k])
		if ip == nil || ip.To4() == nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsIPv6 每个元素值都是IPv6地址
func (strSliceType *StringSliceType) IsIPv6(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		ip := net.ParseIP(strSliceType.value[k])
		if ip == nil || ip.To16() == nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsIP 每个元素值都是IPv4或IPv6地址
func (strSliceType *StringSliceType) IsIP(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsTCPAddr 每个元素值都是 IP:Port 格式
func (strSliceType *StringSliceType) IsTCPAddr(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsMAC 每个元素值都是MAC地址
func (strSliceType *StringSliceType) IsMAC(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsSQLObject 每个元素值都是有效的SQL对象名
func (strSliceType *StringSliceType) IsSQLObject(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsUUID 每个元素值都是UUID格式
func (strSliceType *StringSliceType) IsUUID(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsULID 每个元素值都是ULID格式
func (strSliceType *StringSliceType) IsULID(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
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

// IsURL 每个元素值都是有效的URL
func (strSliceType *StringSliceType) IsURL(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if _, err := url.ParseRequestURI(strSliceType.value[k]); err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// IsChineseIDCard 每个元素值都是中国大陆身份证
func (strSliceType *StringSliceType) IsChineseIDCard(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if !isChineseIDCard(strSliceType.value[k]) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}
