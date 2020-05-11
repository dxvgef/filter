package filter

import (
	"strings"
	"unicode"
)

// 必须包含字母(不区分大小写)
func (obj *Object) MustHasLetter(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsLetter(v) {
			return obj
		}
	}

	obj.err = obj.setError(customError...)
	return obj
}

// 必须包含小写字母
func (obj *Object) MustHasLower(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsLower(v) {
			return obj
		}
	}

	obj.err = obj.setError(customError...)
	return obj
}

// 必须包含大写字母
func (obj *Object) MustHasUpper(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsUpper(v) {
			return obj
		}
	}

	obj.err = obj.setError(customError...)
	return obj
}

// 必须包含数字
func (obj *Object) MustHasDigit(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsDigit(v) {
			return obj
		}
	}

	obj.err = obj.setError(customError...)
	return obj
}

// 必须包含符号
func (obj *Object) MustHasSymbol(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}

	for _, v := range obj.rawValue {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) && !unicode.Is(unicode.Han, v) {
			return obj
		}
	}

	obj.err = obj.setError(customError...)
	return obj
}

// 必须包含特定的字符串
func (obj *Object) MustHasString(sub string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if !strings.Contains(obj.rawValue, sub) {
		obj.err = obj.setError(customError...)
		return obj
	}
	return obj
}

// 必须包含指定的前缀字符串
func (obj *Object) MustHasPrefix(sub string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if !strings.HasPrefix(obj.rawValue, sub) {
		obj.err = obj.setError(customError...)
		return obj
	}
	return obj
}

// 必须包含指定的后缀字符串
func (obj *Object) MustHasSuffix(sub string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if !strings.HasSuffix(obj.rawValue, sub) {
		obj.err = obj.setError(customError...)
		return obj
	}
	return obj
}
