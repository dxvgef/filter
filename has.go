package filter

import (
	"strings"
	"unicode"
)

// HasLetter 存在字母
func (obj *Object) HasLetter(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsLetter(v) {
			return obj
		}
	}

	obj.err = obj.setError("必须包含字母", customError...)
	return obj
}

// HasLower 存在小写字母
func (obj *Object) HasLower(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsLower(v) {
			return obj
		}
	}

	obj.err = obj.setError("必须包含小写字母", customError...)
	return obj
}

// HasLower 存在大写字母
func (obj *Object) HasUpper(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsUpper(v) {
			return obj
		}
	}

	obj.err = obj.setError("必须包含大写字母", customError...)
	return obj
}

// HasDigit 存在数字
func (obj *Object) HasDigit(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsDigit(v) {
			return obj
		}
	}

	obj.err = obj.setError("必须包含数字", customError...)
	return obj
}

// HasSymbol 存在符号
func (obj *Object) HasSymbol(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) && !unicode.Is(unicode.Han, v) {
			return obj
		}
	}

	obj.err = obj.setError("必须包含符号", customError...)
	return obj
}

// Contains 存在指定字符串
func (obj *Object) Contains(sub string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if !strings.Contains(obj.rawValue, sub) {
		obj.err = obj.setError("不允许的值", customError...)
		return obj
	}
	return obj
}

// HasPrefix 存在指定的前缀字符串
func (obj *Object) HasPrefix(sub string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if !strings.HasPrefix(obj.rawValue, sub) {
		obj.err = obj.setError("不允许的值", customError...)
		return obj
	}
	return obj
}
