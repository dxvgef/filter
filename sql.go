package filter

import (
	"strings"
	"unicode"
)

func isSQLobject(value string) bool {
	for _, v := range value {
		if unicode.IsLetter(v) == false && unicode.IsDigit(v) == false && v != '-' && v != '_' {
			return false
		}
	}
	return true
}

// IsSQLobject SQL对象（库名、表名、字段名）
func (obj *Object) IsSQLobject(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	if isSQLobject(obj.rawValue) == false {
		obj.err = obj.setError("必须是字母、数字或-_符号", customError...)
		return obj
	}
	return obj
}

// IsSQLobjects SQL对象集合（库名、表名、字段名）
func (obj *Object) IsSQLobjects(sep string, customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	values := strings.Split(obj.rawValue, sep)
	if values[0] == "" || values[0] == " " {
		obj.err = obj.setError("必须是字母、数字或-_符号", customError...)
		return obj
	}
	for _, v := range values {
		if isSQLobject(v) == false {
			obj.err = obj.setError("必须是字母、数字或-_符号", customError...)
			return obj
		}
	}
	obj.sep = sep
	return obj
}
