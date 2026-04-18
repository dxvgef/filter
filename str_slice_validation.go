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

// AllContain 所有元素都包含
func (strSliceType *StringSliceType) AllContain(sub string, customError ...string) *StringSliceType {
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

// NotContain 没有元素包含
func (strSliceType *StringSliceType) NotContain(sub string, customError ...string) *StringSliceType {
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

// AnyContain 任意元素包含
func (strSliceType *StringSliceType) AnyContain(sub string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for i := range strSliceType.value {
		if strings.Contains(strSliceType.value[i], sub) {
			return strSliceType
		}
	}
	strSliceType.err = wrapError(strSliceType.name, customError...)
	return strSliceType
}

// AllIn 所有元素都包含在列表中
func (strSliceType *StringSliceType) AllIn(values []string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for i := range strSliceType.value {
		if !slices.Contains(values, strSliceType.value[i]) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// AnyIn 任意元素在列表中
func (strSliceType *StringSliceType) AnyIn(values []string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}
	strSliceType.err = wrapError(strSliceType.name, customError...)
	for i := range strSliceType.value {
		if slices.Contains(values, strSliceType.value[i]) {
			strSliceType.err = nil
			break
		}
	}
	return strSliceType
}

// NotIn 没有元素在列表中
func (strSliceType *StringSliceType) NotIn(values []string, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for i := range strSliceType.value {
		if slices.Contains(values, strSliceType.value[i]) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// AllowChars 只允许使用的字符
func (strSliceType *StringSliceType) AllowChars(values []rune, customError ...string) *StringSliceType {
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

// BlockChars 不允许使用的字符
func (strSliceType *StringSliceType) BlockChars(values []rune, customError ...string) *StringSliceType {
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

// AllowSymbols 只允许使用的符号
func (strSliceType *StringSliceType) AllowSymbols(values []rune, customError ...string) *StringSliceType {
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

// BlockSymbols 不允许使用的符号
func (strSliceType *StringSliceType) BlockSymbols(values []rune, customError ...string) *StringSliceType {
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

// HasLetter 包含字母(不区分大小写)
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

// HasLower 包含小写字母
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

// HasUpper 包含大写字母
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

// HasNumber 包含数字
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

// HasSymbol 包含符号
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

// HasPrefix 包含指定的前缀
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

// HasSuffix 包含指定的后缀
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

// MaxCount 元素数量最大值
func (strSliceType *StringSliceType) MaxCount(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	if len(strSliceType.value) > value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
	}
	return strSliceType
}

// MinCount 元素数量最小值
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

// CountIs 元素数量等于
func (strSliceType *StringSliceType) CountIs(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	if len(strSliceType.value) != value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	return strSliceType
}

// CountIsNot 元素数量不等于
func (strSliceType *StringSliceType) CountIsNot(value int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	if len(strSliceType.value) == value {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType
	}
	return strSliceType
}

// LengthIs 元素长度等于
func (strSliceType *StringSliceType) LengthIs(value int, customError ...string) *StringSliceType {
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

// LengthIsNot 元素长度不等于
func (strSliceType *StringSliceType) LengthIsNot(value int, customError ...string) *StringSliceType {
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

// MaxLength 元素长度最大值
func (strSliceType *StringSliceType) MaxLength(value int, customError ...string) *StringSliceType {
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

// MinLength 元素长度最小值
func (strSliceType *StringSliceType) MinLength(value int, customError ...string) *StringSliceType {
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

// LengthRange 元素长度范围
func (strSliceType *StringSliceType) LengthRange(minValue, maxValue int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		l := len(strSliceType.value[k])
		if l < minValue || l > maxValue {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// UTF8LengthIs 元素的UTF8编码长度等于
func (strSliceType *StringSliceType) UTF8LengthIs(value int, customError ...string) *StringSliceType {
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

// UTF8LengthIsNot 元素的UTF8编码长度不等于
func (strSliceType *StringSliceType) UTF8LengthIsNot(value int, customError ...string) *StringSliceType {
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

// MaxUTF8Length UTF8编码，元素长度最大值
func (strSliceType *StringSliceType) MaxUTF8Length(value int, customError ...string) *StringSliceType {
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

// MinUTF8Length UTF8编码，字符串长度最小值
func (strSliceType *StringSliceType) MinUTF8Length(value int, customError ...string) *StringSliceType {
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

// UTF8LengthRange UTF8编码，元素长度范围
func (strSliceType *StringSliceType) UTF8LengthRange(minValue, maxValue int, customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		l := utf8.RuneCountInString(strSliceType.value[k])
		if l < minValue || l > maxValue {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// IsLower 是小写字母
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

// IsUpper 是大写字母
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

// IsLetter 是字母
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

// IsLowerOrNumber 是小写字母或数字
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

// IsUpperOrNumber 是大写字母或数字
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

// IsLetterOrNumber 是字母或数字
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

// IsChinese 是汉字
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

// IsMail 是电邮地址
func (strSliceType *StringSliceType) IsMail(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		_, err := mail.ParseAddress(strSliceType.value[k])
		if err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsIPv4 是IPv4地址
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

// IsIPv6 是IPv6地址
func (strSliceType *StringSliceType) IsIPv6(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		ip := net.ParseIP(strSliceType.value[k])
		if ip == nil || ip.To4() != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsIP 是IPv4或IPv6地址
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

// IsTCPAddr 是 IP:Port 格式
func (strSliceType *StringSliceType) IsTCPAddr(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if !isTCPAddr(strSliceType.value[k]) {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsMAC 是MAC地址
func (strSliceType *StringSliceType) IsMAC(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		_, err := net.ParseMAC(strSliceType.value[k])
		if err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}
	return strSliceType
}

// IsSQLObject 是有效的SQL对象名
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

// IsUUID 是UUID格式
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

// IsULID 是ULID格式
func (strSliceType *StringSliceType) IsULID(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		if len(strSliceType.value[k]) != 26 {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
		for i := 0; i < 26; i++ {
			c := strSliceType.value[k][i]
			if c >= 'a' && c <= 'z' {
				c -= 32
			}
			if !ulidValid[c] {
				strSliceType.err = wrapError(strSliceType.name, customError...)
				return strSliceType
			}
		}
	}

	return strSliceType
}

// IsURL 是有效的URL
func (strSliceType *StringSliceType) IsURL(customError ...string) *StringSliceType {
	if strSliceType.err != nil {
		return strSliceType
	}

	for k := range strSliceType.value {
		_, err := url.ParseRequestURI(strSliceType.value[k])
		if err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			break
		}
	}

	return strSliceType
}

// IsChineseIDCard 是中国大陆身份证
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
