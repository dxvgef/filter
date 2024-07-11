package filter

import (
	"encoding/base64"
	"html"
	"net/url"
	"strings"
)

// Trim 删除每个元素左右的指定字符串
func (strSliceType *StringSliceType) Trim(sub string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = strings.Trim(strSliceType.value[k], sub)
	}
	return strSliceType
}

// TrimSpace 删除每个元素左右的空格
func (strSliceType *StringSliceType) TrimSpace() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = strings.TrimSpace(strSliceType.value[k])
	}
	return strSliceType
}

// TrimLeft 删除每个元素左边指定的字符串
func (strSliceType *StringSliceType) TrimLeft(sub string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = strings.TrimLeft(strSliceType.value[k], sub)
	}
	return strSliceType
}

// TrimRight 删除每个元素右边指定的字符串
func (strSliceType *StringSliceType) TrimRight(sub string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = strings.TrimRight(strSliceType.value[k], sub)
	}
	return strSliceType
}

// TrimPrefix 删除每个元素指定的前缀字符串
func (strSliceType *StringSliceType) TrimPrefix(sub string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = strings.TrimPrefix(strSliceType.value[k], sub)
	}
	return strSliceType
}

// TrimSuffix 删除每个元素指定的后缀字符串
func (strSliceType *StringSliceType) TrimSuffix(sub string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = strings.TrimSuffix(strSliceType.value[k], sub)
	}
	return strSliceType
}

// DeleteEmpty 删除空字符串的元素
func (strSliceType *StringSliceType) DeleteEmpty() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	var newValue []string
	for k := range strSliceType.value {
		if strSliceType.value[k] != "" {
			newValue = append(newValue, strSliceType.value[k])
		}
	}
	strSliceType.value = newValue
	return strSliceType
}

// RemoveSpace 删除每个元素中的所有空格
func (strSliceType *StringSliceType) RemoveSpace() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = strings.ReplaceAll(strSliceType.value[k], " ", "")
	}
	return strSliceType
}

// Base64StdEncode 对每个元素进行Base64 std编码
func (strSliceType *StringSliceType) Base64StdEncode() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = base64.StdEncoding.EncodeToString(strToBytes(strSliceType.value[k]))
	}
	return strSliceType
}

// Base64StdDecode 对每个元素进行Base64 std解码
func (strSliceType *StringSliceType) Base64StdDecode(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		bytes, err := base64.StdEncoding.DecodeString(strSliceType.value[k])
		if err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
		strSliceType.value[k] = bytesToStr(bytes)
	}
	return strSliceType
}

// Base64RawStdEncode 对每个元素进行Base64 RawStd编码
func (strSliceType *StringSliceType) Base64RawStdEncode() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = base64.RawStdEncoding.EncodeToString(strToBytes(strSliceType.value[k]))
	}
	return strSliceType
}

// Base64RawStdDecode 对每个元素进行Base64 RawStd解码
func (strSliceType *StringSliceType) Base64RawStdDecode(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		bytes, err := base64.RawStdEncoding.DecodeString(strSliceType.value[k])
		if err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
		strSliceType.value[k] = bytesToStr(bytes)
	}
	return strSliceType
}

// Base64URLEncode base64 URL编码
func (strSliceType *StringSliceType) Base64URLEncode() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = base64.URLEncoding.EncodeToString(strToBytes(strSliceType.value[k]))
	}
	return strSliceType
}

// Base64URLDecode base64 URL解码
func (strSliceType *StringSliceType) Base64URLDecode(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		bytes, err := base64.URLEncoding.DecodeString(strSliceType.value[k])
		if err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
		strSliceType.value[k] = bytesToStr(bytes)
	}
	return strSliceType
}

// Base64RawURLEncode base64 RawURL编码
func (strSliceType *StringSliceType) Base64RawURLEncode() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = base64.RawURLEncoding.EncodeToString(strToBytes(strSliceType.value[k]))
	}
	return strSliceType
}

// Base64RawURLDecode base64 RawURL解码
func (strSliceType *StringSliceType) Base64RawURLDecode(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		bytes, err := base64.RawURLEncoding.DecodeString(strSliceType.value[k])
		if err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
		strSliceType.value[k] = bytesToStr(bytes)
	}
	return strSliceType
}

// HTMLEscape html.EscapeString
func (strSliceType *StringSliceType) HTMLEscape() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = html.EscapeString(strSliceType.value[k])
	}
	return strSliceType
}

// HTMLUnescape html.UnescapeString
func (strSliceType *StringSliceType) HTMLUnescape() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = html.UnescapeString(strSliceType.value[k])
	}
	return strSliceType
}

// URLPathEscape 与url.PathEscape相同
func (strSliceType *StringSliceType) URLPathEscape() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = url.PathEscape(strSliceType.value[k])
	}
	return strSliceType
}

// URLPathUnescape url.PathUnescape
func (strSliceType *StringSliceType) URLPathUnescape(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		value, err := url.PathUnescape(strSliceType.value[k])
		if err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
		strSliceType.value[k] = value
	}
	return strSliceType
}

// URLQueryEscape 与url.QueryEscape相同
func (strSliceType *StringSliceType) URLQueryEscape() *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		strSliceType.value[k] = url.QueryEscape(strSliceType.value[k])
	}
	return strSliceType
}

// URLQueryUnescape 与url.QueryUnescape相同
func (strSliceType *StringSliceType) URLQueryUnescape(customError ...string) *StringSliceType {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		return strSliceType
	}
	for k := range strSliceType.value {
		value, err := url.QueryUnescape(strSliceType.value[k])
		if err != nil {
			strSliceType.err = wrapError(strSliceType.name, customError...)
			return strSliceType
		}
		strSliceType.value[k] = value
	}
	return strSliceType
}
