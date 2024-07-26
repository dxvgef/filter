package filter

import (
	"encoding/base64"
	"html"
	"net/url"
	"strings"
)

// ToUpper 字母转为大写
func (strType *StringType) ToUpper() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.ToUpper(strType.value)
	return strType
}

// ToLower 字母转为小写
func (strType *StringType) ToLower() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.ToLower(strType.value)
	return strType
}

// Trim 删除左右的指定字符串
func (strType *StringType) Trim(sub string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.Trim(strType.value, sub)
	return strType
}

// TrimSpace 删除左右所有的空格
func (strType *StringType) TrimSpace() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.TrimSpace(strType.value)
	return strType
}

// TrimLeft 删除左边指定的字符串
func (strType *StringType) TrimLeft(sub string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.TrimLeft(strType.value, sub)
	return strType
}

// TrimRight 删除右边指定的字符串
func (strType *StringType) TrimRight(sub string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.TrimRight(strType.value, sub)
	return strType
}

// TrimPrefix 删除指定的前缀字符串
func (strType *StringType) TrimPrefix(sub string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.TrimPrefix(strType.value, sub)
	return strType
}

// TrimSuffix 删除指定的后缀字符串
func (strType *StringType) TrimSuffix(sub string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.TrimSuffix(strType.value, sub)
	return strType
}

// Replace 替换指定的字符串，可指定替换次数
func (strType *StringType) Replace(oldStr, newStr string, n int) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.Replace(strType.value, oldStr, newStr, n)
	return strType
}

// ReplaceAll 替换所有指定的字符串
func (strType *StringType) ReplaceAll(oldStr, newStr string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.ReplaceAll(strType.value, oldStr, newStr)
	return strType
}

// RemoveSpace 删除所有空格
func (strType *StringType) RemoveSpace() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = strings.ReplaceAll(strType.value, " ", "")
	return strType
}

// Base64StdEncode Base64 std编码
func (strType *StringType) Base64StdEncode() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}
	strType.value = base64.StdEncoding.EncodeToString(strToBytes(strType.value))
	return strType
}

// Base64StdDecode Base64 std解码
func (strType *StringType) Base64StdDecode(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	bytes, err := base64.StdEncoding.DecodeString(strType.value)
	if err != nil {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}

	strType.value = bytesToStr(bytes)
	return strType
}

// base64 RawStd编码
func (strType *StringType) Base64RawStdEncode() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	strType.value = base64.RawStdEncoding.EncodeToString(strToBytes(strType.value))
	return strType
}

// base64 RawStd解码
func (strType *StringType) Base64RawStdDecode(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	bytes, err := base64.RawStdEncoding.DecodeString(strType.value)
	if err != nil {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}

	strType.value = bytesToStr(bytes)
	return strType
}

// base64 URL编码
func (strType *StringType) Base64URLEncode() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	strType.value = base64.URLEncoding.EncodeToString(strToBytes(strType.value))
	return strType
}

// base64 URL解码
func (strType *StringType) Base64URLDecode(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	bytes, err := base64.URLEncoding.DecodeString(strType.value)
	if err != nil {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}

	strType.value = bytesToStr(bytes)
	return strType
}

// base64 RawURL编码
func (strType *StringType) Base64RawURLEncode() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	strType.value = base64.RawURLEncoding.EncodeToString(strToBytes(strType.value))
	return strType
}

// base64 RawURL解码
func (strType *StringType) Base64RawURLDecode(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	bytes, err := base64.RawURLEncoding.DecodeString(strType.value)
	if err != nil {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}

	strType.value = bytesToStr(bytes)
	return strType
}

// html.EscapeString
func (strType *StringType) HTMLEscape() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	strType.value = html.EscapeString(strType.value)
	return strType
}

// html.UnescapeString
func (strType *StringType) HTMLUnescape() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	strType.value = html.UnescapeString(strType.value)
	return strType
}

// 与url.PathEscape相同
func (strType *StringType) URLPathEscape() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	strType.value = url.PathEscape(strType.value)
	return strType
}

// url.PathUnescape
func (strType *StringType) URLPathUnescape(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	value, err := url.PathUnescape(strType.value)
	if err != nil {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}

	strType.value = value
	return strType
}

// URLQueryEscape 与url.QueryEscape相同
func (strType *StringType) URLQueryEscape() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	strType.value = url.QueryEscape(strType.value)
	return strType
}

// 与url.QueryUnescape相同
func (strType *StringType) URLQueryUnescape(customError ...string) *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	value, err := url.QueryUnescape(strType.value)
	if err != nil {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}

	strType.value = value
	return strType
}
