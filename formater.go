package filter

import (
	"encoding/base64"
	"html"
	"net/url"
	"strings"
)

// base64 std编码
func (obj *Object) Base64StdEncode() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = base64.StdEncoding.EncodeToString(strToBytes(obj.rawValue))
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// base64 std解码
func (obj *Object) Base64StdDecode() *Object {
	if obj.rawValue == "" {
		return obj
	}
	bs, err := base64.StdEncoding.DecodeString(obj.rawValue)
	if err != nil {
		obj.err = err
		return obj
	}
	obj.rawValue = bytesToStr(bs)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// base64 RawStd编码
func (obj *Object) Base64RawStdEncode() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = base64.RawStdEncoding.EncodeToString(strToBytes(obj.rawValue))
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// base64 RawStd解码
func (obj *Object) Base64RawStdDecode() *Object {
	if obj.rawValue == "" {
		return obj
	}
	bs, err := base64.RawStdEncoding.DecodeString(obj.rawValue)
	if err != nil {
		obj.err = err
		return obj
	}
	obj.rawValue = bytesToStr(bs)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// base64 URL编码
func (obj *Object) Base64URLEncode() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = base64.URLEncoding.EncodeToString(strToBytes(obj.rawValue))
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// base64 URL解码
func (obj *Object) Base64URLDecode() *Object {
	if obj.rawValue == "" {
		return obj
	}
	bs, err := base64.URLEncoding.DecodeString(obj.rawValue)
	if err != nil {
		obj.err = err
		return obj
	}
	obj.rawValue = bytesToStr(bs)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// base64 RawURL编码
func (obj *Object) Base64RawURLEncode() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = base64.RawURLEncoding.EncodeToString(strToBytes(obj.rawValue))
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// base64 RawURL解码
func (obj *Object) Base64RawURLDecode() *Object {
	if obj.rawValue == "" {
		return obj
	}
	bs, err := base64.RawURLEncoding.DecodeString(obj.rawValue)
	if err != nil {
		obj.err = err
		return obj
	}
	obj.rawValue = bytesToStr(bs)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// 与html.UnescapeString相同
func (obj *Object) UnescapeHTML() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = html.UnescapeString(obj.rawValue)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// 与html.EscapeString相同
func (obj *Object) EscapeHTML() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = html.EscapeString(obj.rawValue)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// 与url.PathUnescape相同
func (obj *Object) UnescapeURLPath() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue, obj.err = url.PathUnescape(obj.rawValue)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// 与url.PathEscape相同
func (obj *Object) EscapeURLPath() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = url.PathEscape(obj.rawValue)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// 与url.QueryUnescape相同
func (obj *Object) UnescapeURLQuery() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue, obj.err = url.QueryUnescape(obj.rawValue)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// 与url.QueryEscape相同
func (obj *Object) EscapeURLQuery() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = url.QueryEscape(obj.rawValue)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// Trim 去除前后空格
func (obj *Object) Trim() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = strings.TrimSpace(obj.rawValue)
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// 去除所有的空格
func (obj *Object) RemoveSpace() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = strings.ReplaceAll(obj.rawValue, " ", "")
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// ReplaceAll 替换所有
func (obj *Object) ReplaceAll(old, new string) *Object {
	if obj.rawValue == "" {
		return obj
	}
	count := strings.Count(obj.rawValue, old)
	for i := 0; i < count; i++ {
		obj.rawValue = strings.ReplaceAll(obj.rawValue, old, new)
	}
	if obj.require && obj.rawValue == "" {
		return obj.required()
	}
	return obj
}

// SnakeCaseToCamelCase 蛇形转驼峰: hello_world => helloWorld
func (obj *Object) SnakeCaseToCamelCase() *Object {
	if obj.rawValue == "" {
		return obj
	}
	slice := strings.Split(obj.rawValue, "_")
	for i := range slice {
		if i > 0 {
			slice[i] = strings.Title(slice[i])
		}
	}
	obj.rawValue = strings.Join(slice, "")
	return obj
}

// SnakeCaseToPascalCase 蛇形转帕斯卡: hello_world => HelloWorld
func (obj *Object) SnakeCaseToPascalCase() *Object {
	if obj.rawValue == "" {
		return obj
	}
	slice := strings.Split(obj.rawValue, "_")
	for i := range slice {
		slice[i] = strings.Title(slice[i])
	}
	obj.rawValue = strings.Join(slice, "")
	return obj
}

// CamelCaseToSnakeCase 驼峰(含帕斯卡)转蛇形 helloWorld/HelloWorld => hello_world
func (obj *Object) CamelCaseToSnakeCase() *Object {
	if obj.rawValue == "" {
		return obj
	}
	strLen := len(obj.rawValue)
	result := make([]byte, 0, strLen*2)
	j := false
	for i := 0; i < strLen; i++ {
		char := obj.rawValue[i]
		if i > 0 && char >= 'A' && char <= 'Z' && j {
			result = append(result, '_')
		}
		if char != '_' {
			j = true
		}
		result = append(result, char)
	}
	obj.rawValue = strings.ToLower(string(result))
	return obj
}
