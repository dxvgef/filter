package filter

import (
	"strconv"
	"unicode/utf8"
)

// MinLength 最小长度
func (obj *Object) MinLength(min int, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if len(obj.rawValue) < min {
		obj.err = obj.setError("长度不能小于"+strconv.Itoa(min)+"个字符", customError...)
	}
	return obj
}

// MinUTF8Length UTF8编码最小长度
func (obj *Object) MinUTF8Length(min int, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if utf8.RuneCountInString(obj.rawValue) < min {
		obj.err = obj.setError("长度不能小于"+strconv.Itoa(min)+"个字符", customError...)
	}
	return obj
}

// MaxLength 最大长度
func (obj *Object) MaxLength(max int, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if len(obj.rawValue) > max {
		obj.err = obj.setError("长度不能大于"+strconv.Itoa(max)+"个字符", customError...)
	}
	return obj
}

// MaxUTF8Length UTF8编码最大长度
func (obj *Object) MaxUTF8Length(max int, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if utf8.RuneCountInString(obj.rawValue) > max {
		obj.err = obj.setError("长度不能大于"+strconv.Itoa(max)+"个字符", customError...)
	}
	return obj
}
