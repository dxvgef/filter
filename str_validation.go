package filter

import (
	"encoding/json"
	"net"
	"net/mail"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Require 不能为零值
func (strType *StringType) Require(customError ...string) *StringType {
	if strType.value == "" {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}
	return strType
}

// IsLower 小写字母
func (strType *StringType) IsLower(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsLower(v) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}

	return strType
}

// IsUpper 大写字母
func (strType *StringType) IsUpper(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsUpper(v) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}

	return strType
}

// IsLetter 大小写字母
func (strType *StringType) IsLetter(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	for _, v := range strType.value {
		if !unicode.IsLetter(v) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}
	return strType
}

// IsLowerOrNumber 小写字母或数字
func (strType *StringType) IsLowerOrNumber(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	for _, v := range strType.value {
		if !unicode.IsLower(v) && !unicode.IsDigit(v) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}
	return strType
}

// IsUpperOrNumber 大写字母或数字
func (strType *StringType) IsUpperOrNumber(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	for _, v := range strType.value {
		if !unicode.IsUpper(v) && !unicode.IsDigit(v) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}
	return strType
}

// IsLetterOrNumber 字母或数字
func (strType *StringType) IsLetterOrNumber(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	for _, v := range strType.value {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}
	return strType
}

// IsChinese 汉字
func (strType *StringType) IsChinese(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	for _, v := range strType.value {
		if !unicode.Is(unicode.Scripts["Han"], v) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}
	return strType
}

// IsMail 电邮地址
func (strType *StringType) IsMail(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	_, err := mail.ParseAddress(strType.value)
	if err != nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsIP IPv4/v6地址
func (strType *StringType) IsIP(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	if net.ParseIP(strType.value) == nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsTCPAddr IP:Port
func (strType *StringType) IsTCPAddr(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	if _, err := net.ResolveTCPAddr("tcp", strType.value); err != nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsMAC MAC地址
func (strType *StringType) IsMAC(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	if _, err := net.ParseMAC(strType.value); err != nil {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}

	return strType
}

// IsJSON JSON格式
func (strType *StringType) IsJSON(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	var js json.RawMessage
	if json.Unmarshal([]byte(strType.value), &js) != nil {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}
	return strType
}

// IsChinaIDNumber 中国身份证号码
func (strType *StringType) IsChinaIDNumber(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	var idV int
	if strType.value[17:] == "X" {
		idV = 88
	} else {
		var err error
		if idV, err = strconv.Atoi(strType.value[17:]); err != nil {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}

	var verify int
	id := strType.value[:17]
	arr := make([]int, 17)
	for i := 0; i < 17; i++ {
		arr[i], strType.err = strconv.Atoi(string(id[i]))
		if strType.err != nil {
			return strType
		}
	}
	wi := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		res += arr[i] * wi[i]
	}
	verify = res % 11

	var temp int
	a18 := [11]int{1, 0, 88 /* 'X' */, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			break
		}
	}
	if temp != idV {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsSQLObject SQL对象（库名、表名、字段名）
func (strType *StringType) IsSQLObject(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	if !isSQLObject(strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsURL 是有效的URL
func (strType *StringType) IsURL(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	if _, err := url.ParseRequestURI(strType.value); err != nil {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsUUID UUID格式
func (strType *StringType) IsUUID(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
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
	if strType.err != nil || strType.value == "" {
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
			return strType
		}
	}

	return strType
}

// IsChineseIDCard 中国大陆地区身份证号码
func (strType *StringType) IsChineseIDCard(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	var idV int
	if strType.value[17:] == "X" {
		idV = 88
	} else {
		var err error
		if idV, err = strconv.Atoi(strType.value[17:]); err != nil {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}

	var verify int
	id := strType.value[:17]
	arr := make([]int, 17)
	for i := 0; i < 17; i++ {
		arr[i], strType.err = strconv.Atoi(string(id[i]))
		if strType.err != nil {
			return strType
		}
	}
	wi := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		res += arr[i] * wi[i]
	}
	verify = res % 11

	var temp int
	a18 := [11]int{1, 0, 88 /* 'X' */, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			break
		}
	}
	if temp != idV {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// Length 指定的长度
func (strType *StringType) Length(value int, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	if len(strType.value) != value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8Length 按UTF8编码指定长度
func (strType *StringType) UTF8Length(value int, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	if utf8.RuneCountInString(strType.value) != value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// MinLength 最小长度
func (strType *StringType) MinLength(min int, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	if len(strType.value) < min {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8MinLength 按UTF8编码校验最小长度
func (strType *StringType) UTF8MinLength(min int, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	if utf8.RuneCountInString(strType.value) < min {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// MaxLength 最大长度
func (strType *StringType) MaxLength(max int, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	if len(strType.value) > max {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8MaxLength 按UTF8编码校验最大长度
func (strType *StringType) UTF8MaxLength(max int, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	if utf8.RuneCountInString(strType.value) > max {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// AllowedChars 仅允许字符串中包启allowValues内的字符
func (strType *StringType) AllowedChars(allowValues []rune, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	for _, r := range strType.value {
		if !slices.Contains(allowValues, r) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}

	return strType
}

// AllowedSymbols 如果有符号，只允许存在指定的符号
func (strType *StringType) AllowedSymbols(allowValues []rune, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	for _, r := range strType.value {
		// 如果是符号
		if unicode.IsPunct(r) {
			if !slices.Contains(allowValues, r) {
				strType.err = wrapError(strType.name, customError...)
				return strType
			}
		}
	}
	return strType
}

// AllowedStrings 只允许存在指定的字符串（枚举）
func (strType *StringType) AllowedStrings(allowedValues []string, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	if !slices.Contains(allowedValues, strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// DeniedChars 阻止deniedValues中的值
func (strType *StringType) DeniedChars(deniedValues []rune, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	for _, r := range strType.value {
		if slices.Contains(deniedValues, r) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}
	return strType
}

// DeniedSymbols 如果有符号，阻止存在指定的符号
func (strType *StringType) DeniedSymbols(deniedValues []rune, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	for _, r := range strType.value {
		// 如果是符号
		if unicode.IsPunct(r) {
			if slices.Contains(deniedValues, r) {
				strType.err = wrapError(strType.name, customError...)
				return strType
			}
		}
	}
	return strType
}

// DeniedStrings 禁止存在指定的字符串
func (strType *StringType) DeniedStrings(deniedValues []string, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	if slices.Contains(deniedValues, strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// 包含了字母(不区分大小写)
func (strType *StringType) HasLetter(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	for _, v := range strType.value {
		if unicode.IsLetter(v) {
			return strType
		}
	}
	strType.err = wrapError(strType.name, customError...)
	return strType
}

// HasLower 包含了小写字母
func (strType *StringType) HasLower(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
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
	if strType.err != nil || strType.value == "" {
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
	if strType.err != nil || strType.value == "" {
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
	if strType.err != nil || strType.value == "" {
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

// Contains 必须包含指定的字符串
func (strType *StringType) Contains(sub string, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	if !strings.Contains(strType.value, sub) {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}
	return strType
}

// HasPrefix 包含了指定的前缀字符串
func (strType *StringType) HasPrefix(sub string, customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
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
	if strType.err != nil || strType.value == "" {
		return strType
	}
	if strings.HasSuffix(strType.value, sub) {
		return strType
	}
	strType.err = wrapError(strType.name, customError...)
	return strType
}
