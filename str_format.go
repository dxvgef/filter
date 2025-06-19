package filter

import (
	"encoding/base64"
	"html"
	"net/url"
	"strings"
	"unicode"
)

// ToUpper 字母转为大写
func (strType *StringType) ToUpper() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.ToUpper(strType.value)
	return strType
}

// ToLower 字母转为小写
func (strType *StringType) ToLower() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.ToLower(strType.value)
	return strType
}

// Trim 删除左右的指定字符串
func (strType *StringType) Trim(sub string) *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.Trim(strType.value, sub)
	return strType
}

// TrimSpace 删除左右所有的空格
func (strType *StringType) TrimSpace() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.TrimSpace(strType.value)
	return strType
}

// TrimLeft 删除左边所有指定的字符串
func (strType *StringType) TrimLeft(sub string) *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.TrimLeft(strType.value, sub)
	return strType
}

// TrimRight 删除右边所有指定的字符串
func (strType *StringType) TrimRight(sub string) *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.TrimRight(strType.value, sub)
	return strType
}

// TrimPrefix 删除指定的前缀字符串
func (strType *StringType) TrimPrefix(sub string) *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.TrimPrefix(strType.value, sub)
	return strType
}

// TrimSuffix 删除指定的后缀字符串
func (strType *StringType) TrimSuffix(sub string) *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.TrimSuffix(strType.value, sub)
	return strType
}

// Replace 替换指定的字符串，可指定替换次数
func (strType *StringType) Replace(oldStr, newStr string, n int) *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.Replace(strType.value, oldStr, newStr, n)
	return strType
}

// ReplaceAll 替换所有指定的字符串
func (strType *StringType) ReplaceAll(oldStr, newStr string) *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.ReplaceAll(strType.value, oldStr, newStr)
	return strType
}

// RemoveSpace 删除所有空格
func (strType *StringType) RemoveSpace() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = strings.ReplaceAll(strType.value, " ", "")
	return strType
}

// Base64StdEncode 使用base64.StdEncoding编码
func (strType *StringType) Base64StdEncode() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = base64.StdEncoding.EncodeToString(strToBytes(strType.value))
	return strType
}

// Base64StdDecode 使用base64.StdEncoding解码
func (strType *StringType) Base64StdDecode(customError ...string) *StringType {
	if strType.err != nil {
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

// Base64RawStdEncode 使用base64.RawStdEncoding编码
func (strType *StringType) Base64RawStdEncode() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = base64.RawStdEncoding.EncodeToString(strToBytes(strType.value))
	return strType
}

// Base64RawStdDecode 使用base64.RawStdEncoding解码
func (strType *StringType) Base64RawStdDecode(customError ...string) *StringType {
	if strType.err != nil {
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

// Base64URLEncode 使用 base64.URLEncoding 编码
func (strType *StringType) Base64URLEncode() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = base64.URLEncoding.EncodeToString(strToBytes(strType.value))
	return strType
}

// Base64URLDecode 使用 base64.URLEncoding 解码
func (strType *StringType) Base64URLDecode(customError ...string) *StringType {
	if strType.err != nil {
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

// Base64RawURLEncode 使用 base64.RawURLEncoding 编码
func (strType *StringType) Base64RawURLEncode() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = base64.RawURLEncoding.EncodeToString(strToBytes(strType.value))
	return strType
}

// Base64RawURLDecode 使用 base64.RawURLEncoding 解码
func (strType *StringType) Base64RawURLDecode(customError ...string) *StringType {
	if strType.err != nil {
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

// HTMLEscape 使用 html.EscapeString 编码
func (strType *StringType) HTMLEscape() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = html.EscapeString(strType.value)
	return strType
}

// HTMLUnescape 使用 html.UnescapeString 解码
func (strType *StringType) HTMLUnescape() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = html.UnescapeString(strType.value)
	return strType
}

// URLPathEscape 使用 url.PathEscape 编码
func (strType *StringType) URLPathEscape() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = url.PathEscape(strType.value)
	return strType
}

// URLPathUnescape 使用 url.PathUnescape 解码
func (strType *StringType) URLPathUnescape(customError ...string) *StringType {
	if strType.err != nil {
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

// URLQueryEscape 使用 url.QueryEscape 编码
func (strType *StringType) URLQueryEscape() *StringType {
	if strType.err != nil {
		return strType
	}

	strType.value = url.QueryEscape(strType.value)
	return strType
}

// URLQueryUnescape 使用 url.QueryUnescape 解码
func (strType *StringType) URLQueryUnescape(customError ...string) *StringType {
	if strType.err != nil {
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

func toSnakeCase(str string) string {
	if str == "" {
		return ""
	}

	var result []rune
	for i, r := range str {
		if unicode.IsUpper(r) {
			if i > 0 {
				prev := rune(str[i-1])
				if prev != '_' && !unicode.IsUpper(prev) && !unicode.IsDigit(prev) {
					result = append(result, '_')
				}
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// ToSnakeCase 转换为蛇形命名 aa_bb
func (strType *StringType) ToSnakeCase() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	strType.value = toSnakeCase(strType.value)
	return strType
}

func toCamelCase(str string) string {
	if str == "" {
		return ""
	}

	// 按下划线分割字符串
	parts := strings.Split(str, "_")
	if len(parts) == 0 {
		return ""
	}

	var result strings.Builder

	// 处理首段的第一个字符：强制小写
	firstPart := parts[0]
	if len(firstPart) > 0 {
		result.WriteRune(unicode.ToLower(rune(firstPart[0])))
		result.WriteString(firstPart[1:])
	}

	// 拼接其余部分（保留原有大小写）
	for _, part := range parts[1:] {
		result.WriteString(part)
	}

	return result.String()
}

// toCamelCase 将单个蛇形命名字符串转为驼峰命名 aaBB
func (strType *StringType) ToCamelCase() *StringType {
	if strType.err != nil || strType.value == "" {
		return strType
	}

	strType.value = toCamelCase(strType.value)
	return strType
}
