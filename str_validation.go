package filter

import (
	"encoding/json"
	"net"
	"net/mail"
	"net/url"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Equals 等于
func (strType *StringType) Equals(value string, customError ...string) *StringType {
	if strType.value != value {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// NotEquals 不等于
func (strType *StringType) NotEquals(value string, customError ...string) *StringType {
	if strType.value == value {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// Contains 包含了指定的字符串
func (strType *StringType) Contains(sub string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if !strings.Contains(strType.value, sub) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// NotContains 没有包含指定的字符串
func (strType *StringType) NotContains(sub string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if strings.Contains(strType.value, sub) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// LengthEquals 长度等于
func (strType *StringType) LengthEquals(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if len(strType.value) != value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// LengthNotEquals 长度不等于
func (strType *StringType) LengthNotEquals(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if len(strType.value) == value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// LengthLessThan 长度小于
func (strType *StringType) LengthLessThan(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if len(strType.value) > value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// LengthGreaterThan 长度大于
func (strType *StringType) LengthGreaterThan(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if len(strType.value) < value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8LengthEquals UTF8编码长度等于
func (strType *StringType) UTF8LengthEquals(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if utf8.RuneCountInString(strType.value) != value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8LengthNotEquals UTF8编码长度不等于
func (strType *StringType) UTF8LengthNotEquals(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if utf8.RuneCountInString(strType.value) == value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8LengthLessThan UTF8编码长度小于
func (strType *StringType) UTF8LengthLessThan(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if utf8.RuneCountInString(strType.value) > value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8LengthGreaterThan UTF8编码长度大于
func (strType *StringType) UTF8LengthGreaterThan(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if utf8.RuneCountInString(strType.value) < value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// AllowedValues 只能是数组中的值
func (strType *StringType) AllowedValues(values []string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if !slices.Contains(values, strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// DisallowedValues 不能是数组中的值
func (strType *StringType) DisallowedValues(values []string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if slices.Contains(values, strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// AllowedChars 只能是数组中的字符
func (strType *StringType) AllowedChars(values []rune, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, r := range strType.value {
		if !slices.Contains(values, r) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}

	return strType
}

// InRuneSet 不能是数组中的字符
func (strType *StringType) DisallowedChars(values []rune, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, r := range strType.value {
		if slices.Contains(values, r) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}

	return strType
}

// AllowedSymbols 如果有符号，只能是数组中的符号
func (strType *StringType) AllowedSymbols(values []rune, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, r := range strType.value {
		// 如果是符号
		if unicode.IsPunct(r) && !slices.Contains(values, r) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}
	return strType
}

// DisallowedSymbols 如果有符号，不能是数组中的符号
func (strType *StringType) DisallowedSymbols(values []rune, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, r := range strType.value {
		// 如果是符号
		if unicode.IsPunct(r) {
			if slices.Contains(values, r) {
				strType.err = wrapError(strType.name, customError...)
				break
			}
		}
	}
	return strType
}

// 包含了字母(不区分大小写)
func (strType *StringType) HasLetter(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if unicode.IsLetter(v) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}
	return strType
}

// HasLower 包含小写字母
func (strType *StringType) HasLower(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if unicode.IsLower(v) {
			return strType
		}
	}

	strType.err = wrapError(strType.name, customError...)
	return strType
}

// HasUpper 包含了大写字母
func (strType *StringType) HasUpper(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if unicode.IsUpper(v) {
			return strType
		}
	}

	strType.err = wrapError(strType.name, customError...)
	return strType
}

// HasNumber 包含了数字
func (strType *StringType) HasNumber(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if unicode.IsDigit(v) {
			return strType
		}
	}

	strType.err = wrapError(strType.name, customError...)
	return strType
}

// HasSymbol 包含了符号
func (strType *StringType) HasSymbol(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) && !unicode.Is(unicode.Han, v) {
			return strType
		}
	}

	strType.err = wrapError(strType.name, customError...)
	return strType
}

// HasPrefix 包含了指定的前缀字符串
func (strType *StringType) HasPrefix(sub string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if strings.HasPrefix(strType.value, sub) {
		return strType
	}
	strType.err = wrapError(strType.name, customError...)
	return strType
}

// HasSuffix 包含了指定的后缀字符串
func (strType *StringType) HasSuffix(sub string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if strings.HasSuffix(strType.value, sub) {
		return strType
	}
	strType.err = wrapError(strType.name, customError...)
	return strType
}

// IsLower 是小写字母
func (strType *StringType) IsLower(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsLower(v) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}

	return strType
}

// IsUpper 是大写字母
func (strType *StringType) IsUpper(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsUpper(v) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}

	return strType
}

// IsLetter 是大小写字母
func (strType *StringType) IsLetter(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsLetter(v) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}
	return strType
}

// IsLowerOrNumber 是小写字母或数字
func (strType *StringType) IsLowerOrNumber(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsLower(v) && !unicode.IsDigit(v) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}
	return strType
}

// IsUpperOrNumber 是大写字母或数字
func (strType *StringType) IsUpperOrNumber(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsUpper(v) && !unicode.IsDigit(v) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}
	return strType
}

// IsLetterOrNumber 是大小写字母或数字
func (strType *StringType) IsLetterOrNumber(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}
	return strType
}

// IsChinese 是汉字
func (strType *StringType) IsChinese(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.Is(unicode.Scripts["Han"], v) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}
	return strType
}

// IsMail 是电邮地址
func (strType *StringType) IsMail(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if _, err := mail.ParseAddress(strType.value); err != nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsIPv4 是IPv4地址
func (strType *StringType) IsIPv4(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}
	ip := net.ParseIP(strType.value)
	if ip == nil || ip.To4() == nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsIPv6 是IPv6地址
func (strType *StringType) IsIPv6(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}
	ip := net.ParseIP(strType.value)
	if ip == nil || ip.To16() == nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsIP 是IPv4或IPv6地址
func (strType *StringType) IsIP(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if net.ParseIP(strType.value) == nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsTCPAddr 是 IP:Port 格式
func (strType *StringType) IsTCPAddr(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if _, err := net.ResolveTCPAddr("tcp", strType.value); err != nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsMAC 是MAC地址
func (strType *StringType) IsMAC(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if _, err := net.ParseMAC(strType.value); err != nil {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// IsJSON 是JSON格式
func (strType *StringType) IsJSON(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	var js json.RawMessage
	if json.Unmarshal([]byte(strType.value), &js) != nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsSQLObject 是有效的SQL对象名（可用于库、表、字段等名称）
func (strType *StringType) IsSQLObject(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if !isSQLObject(strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsURL 是有效的URL
func (strType *StringType) IsURL(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if _, err := url.ParseRequestURI(strType.value); err != nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsUUID UUID格式
func (strType *StringType) IsUUID(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if !isUUID(strType.value) {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}
	return strType
}

// IsULID ULID格式
func (strType *StringType) IsULID(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if len(strType.value) != 26 {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}

	validChars := "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

	for _, char := range strType.value {
		if !strings.ContainsRune(validChars, unicode.ToUpper(char)) {
			strType.err = wrapError(strType.name, customError...)
			break
		}
	}

	return strType
}

// IsChineseIDCard 中国大陆地区身份证号码
func (strType *StringType) IsChineseIDCard(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if !isChineseIDCard(strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}
