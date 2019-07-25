package filter

import (
	"strconv"
)

// Separator 指定slice类型的分隔符
func (obj *Object) Separator(sep string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	obj.sep = sep
	return obj
}

// DenyStrings 阻止存在于[]string中的值
func (obj *Object) DenyStrings(slice []string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	for k := range slice {
		if slice[k] == obj.rawValue {
			obj.err = obj.setError("不允许的值", customError...)
			return obj
		}
	}
	return obj
}

// DenyInts 阻止存在于[]int中的值
func (obj *Object) DenyInts(i []int, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value, err := strconv.Atoi(obj.rawValue)
	if err != nil {
		obj.err = obj.setError(err.Error(), customError...)
		return obj
	}
	for k := range i {
		if value == i[k] {
			obj.err = obj.setError("不允许的值", customError...)
			return obj
		}
	}
	return obj
}

// DenyInts32 阻止存在于[]int32中的值
func (obj *Object) DenyInts32(i []int32, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value64, err := strconv.ParseInt(obj.rawValue, 10, 32)
	if err != nil {
		obj.err = obj.setError(err.Error(), customError...)
		return obj
	}
	value32 := int32(value64)
	for k := range i {
		if value32 == i[k] {
			obj.err = obj.setError("不允许的值", customError...)
			return obj
		}
	}
	return obj
}

// DenyInts64 阻止存在于[]int64中的值
func (obj *Object) DenyInts64(i []int64, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if obj.i64 == 0 {
		var err error
		obj.i64, err = strconv.ParseInt(obj.rawValue, 10, 64)
		if err != nil {
			obj.err = obj.setError(err.Error(), customError...)
			return obj
		}
	}
	for k := range i {
		if obj.i64 == i[k] {
			obj.err = obj.setError("不允许的值", customError...)
			return obj
		}
	}
	return obj
}

// DenyFloats32 阻止存在于[]float32中的值
func (obj *Object) DenyFloats32(f []float32, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value, err := strconv.ParseFloat(obj.rawValue, 32)
	if err != nil {
		obj.err = obj.setError(err.Error(), customError...)
		return obj
	}
	value32 := float32(value)
	for k := range f {
		if value32 == f[k] {
			obj.err = obj.setError("不允许的值", customError...)
			return obj
		}
	}
	return obj
}

// DenyFloats64 阻止存在于[]float64中的值
func (obj *Object) DenyInFloats64(f []float64, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if obj.f64 == 0 {
		var err error
		obj.f64, err = strconv.ParseFloat(obj.rawValue, 64)
		if err != nil {
			obj.err = obj.setError(err.Error(), customError...)
			return obj
		}
	}
	for k := range f {
		if obj.f64 == f[k] {
			obj.err = obj.setError("不允许的值", customError...)
			return obj
		}
	}
	return obj
}
