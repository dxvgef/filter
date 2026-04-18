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

// Is 等于
func (strType *StringType) Is(value string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}
	if strType.value != value {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsNot 不等于
func (strType *StringType) IsNot(value string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}
	if strType.value == value {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// Contains 包含
func (strType *StringType) Contains(sub string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if !strings.Contains(strType.value, sub) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// NotContains 没有包含
func (strType *StringType) NotContains(sub string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if strings.Contains(strType.value, sub) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// LenIs 长度是
func (strType *StringType) LenIs(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if len(strType.value) != value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// LenIsNot 长度不是
func (strType *StringType) LenIsNot(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if len(strType.value) == value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// MaxLen 长度最大值
func (strType *StringType) MaxLen(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if len(strType.value) > value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// MinLen 长度最小值
func (strType *StringType) MinLen(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if len(strType.value) < value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// LenRange 长度范围
func (strType *StringType) LenRange(minValue, maxValue int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}
	l := len(strType.value)
	if l < minValue || l > maxValue {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8LenIs UTF8编码长度等于
func (strType *StringType) UTF8LenIs(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if utf8.RuneCountInString(strType.value) != value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8LenIsNot UTF8编码长度不等于
func (strType *StringType) UTF8LenIsNot(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if utf8.RuneCountInString(strType.value) == value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// MaxUTF8Len UTF8编码长度最大值
func (strType *StringType) MaxUTF8Len(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if utf8.RuneCountInString(strType.value) > value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// MinUTF8Len UTF8编码长度最小值
func (strType *StringType) MinUTF8Len(value int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if utf8.RuneCountInString(strType.value) < value {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// UTF8LenRange UTF8编码长度范围
func (strType *StringType) UTF8LenRange(minValue, maxValue int, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}
	l := utf8.RuneCountInString(strType.value)
	if l < minValue || l > maxValue {
		strType.err = wrapError(strType.name, customError...)
	}

	return strType
}

// In 在列表中
func (strType *StringType) In(values []string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if !slices.Contains(values, strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// NotIn 不在列表中
func (strType *StringType) NotIn(values []string, customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	if slices.Contains(values, strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// AllowChars 只允许使用的字符
func (strType *StringType) AllowChars(values []rune, customError ...string) *StringType {
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

// BlockChars 不允许使用的字符
func (strType *StringType) BlockChars(values []rune, customError ...string) *StringType {
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

// AllowSymbols 只允许使用的符号
func (strType *StringType) AllowSymbols(values []rune, customError ...string) *StringType {
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

// BlockSymbols 不允许使用的符号
func (strType *StringType) BlockSymbols(values []rune, customError ...string) *StringType {
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

// HasLetter 包含字母(不区分大小写)
func (strType *StringType) HasLetter(customError ...string) *StringType {
	if strType.err != nil {
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

// HasUpper 包含大写字母
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

// HasNumber 包含数字
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

// HasSymbol 包含符号
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

// HasPrefix 包含指定的前缀字符串
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

// HasSuffix 包含指定的后缀字符串
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

// IsLetter 是字母
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

	_, err := mail.ParseAddress(strType.value)
	if err != nil {
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
	if ip == nil || ip.To4() != nil {
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

	if !isTCPAddr(strType.value) {
		strType.err = wrapError(strType.name, customError...)
	}
	return strType
}

// IsMAC 是MAC地址
func (strType *StringType) IsMAC(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	_, err := net.ParseMAC(strType.value)
	if err != nil {
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

// IsSQLOperator 是有效的 SQL 运算符
func (strType *StringType) IsSQLOperator(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	v := strings.ToUpper(strType.value)
	switch v {
	case "=", "<>", "!=", ">", ">=", "<", "<=", "IN", "NOT IN", "BETWEEN", "LIKE", "NOT LIKE", "IS NULL", "IS NOT NULL":
		return strType
	}

	strType.err = wrapError(strType.name, customError...)
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

	_, err := url.ParseRequestURI(strType.value)
	if err != nil {
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

	for i := 0; i < 26; i++ {
		c := strType.value[i]
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		if !ulidValid[c] {
			strType.err = wrapError(strType.name, customError...)
			return strType
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

// IsHexColor 是Hex颜色值
func (strType *StringType) IsHexColor(customError ...string) *StringType {
	if strType.err != nil {
		return strType
	}

	color := strings.TrimPrefix(strType.value, "#")

	Len := len(color)

	// 检查长度是否符合规则
	if Len != 3 && Len != 6 && Len != 8 && Len != 4 {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}

	// 遍历字符串检查每个字符是否为有效的十六进制字符
	for _, char := range color {
		if !isHexChar(char) {
			strType.err = wrapError(strType.name, customError...)
			return strType
		}
	}

	return strType
}
