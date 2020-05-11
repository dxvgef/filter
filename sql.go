package filter

import (
	"strings"
	"unicode"
)

func isSQLobject(value string) bool {

	for _, v := range value {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '-' && v != '_' {
			return false
		}
	}
	return true
}

// IsSQLobject SQL对象（库名、表名、字段名）
func (obj *Object) IsSQLobject(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if !isSQLobject(obj.rawValue) {
		obj.err = obj.setError(customError...)
		return obj
	}
	return obj
}

// IsSQLobjects SQL对象集合（库名、表名、字段名）
func (obj *Object) IsSQLobjects(sep string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	values := strings.Split(obj.rawValue, sep)
	if values[0] == "" || values[0] == " " {
		obj.err = obj.setError(customError...)
		return obj
	}
	for _, v := range values {
		if !isSQLobject(v) {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	obj.sep = sep
	return obj
}
