package filter

import (
	"strconv"
	"strings"
)

// Separator 指定slice类型的分隔符
func (obj *Object) Separator(sep string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	obj.sep = sep
	return obj
}

// ------------------------------------ In ---------------------------------------

// 值是否存在于指定字符串中
func (obj *Object) InString(allow string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value := strings.Split(obj.rawValue, "")
	for k := range value {
		if !strings.Contains(allow, value[k]) {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	return obj
}

// ------------------------------------ Enum ---------------------------------------
// EnumString 仅允许[]string中的值
func (obj *Object) EnumString(slice []string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	for k := range slice {
		if slice[k] == obj.rawValue {
			return obj
		}
	}
	obj.err = obj.setError(customError...)
	return obj
}

// EnumInt 仅允许[]int中的值
func (obj *Object) EnumInt(i []int, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value, err := strconv.Atoi(obj.rawValue)
	if err != nil {
		obj.err = obj.setError(customError...)
		return obj
	}
	for k := range i {
		if value == i[k] {
			return obj
		}
	}
	obj.err = obj.setError(customError...)
	return obj
}

// EnumInt32 仅允许[]int32中的值
func (obj *Object) EnumInt32(i []int32, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value64, err := strconv.ParseInt(obj.rawValue, 10, 32)
	if err != nil {
		obj.err = obj.setError(customError...)
		return obj
	}
	value32 := int32(value64)
	for k := range i {
		if value32 == i[k] {
			return obj
		}
	}
	obj.err = obj.setError(customError...)
	return obj
}

// EnumInt64 仅允许[]int64中的值
func (obj *Object) EnumInt64(i []int64, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if obj.i64 == 0 {
		var err error
		obj.i64, err = strconv.ParseInt(obj.rawValue, 10, 64)
		if err != nil {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	for k := range i {
		if obj.i64 == i[k] {
			return obj
		}
	}
	obj.err = obj.setError(customError...)
	return obj
}

// EnumFloat32 仅允许[]float32中的值
func (obj *Object) EnumFloat32(f []float32, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value, err := strconv.ParseFloat(obj.rawValue, 32)
	if err != nil {
		obj.err = obj.setError(customError...)
		return obj
	}
	value32 := float32(value)
	for k := range f {
		if value32 == f[k] {
			return obj
		}
	}
	obj.err = obj.setError(customError...)
	return obj
}

// EnumFloat64 仅允许[]float64中的值
func (obj *Object) EnumFloat64(f []float64, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if obj.f64 == 0 {
		var err error
		obj.f64, err = strconv.ParseFloat(obj.rawValue, 64)
		if err != nil {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	for k := range f {
		if obj.f64 == f[k] {
			return obj
		}
	}
	obj.err = obj.setError(customError...)
	return obj
}

// ------------------------------------ Deny ---------------------------------------

// DenyString 阻止存在于[]string中的值
func (obj *Object) DenyString(slice []string, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	for k := range slice {
		if slice[k] == obj.rawValue {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	return obj
}

// DenyInt 阻止存在于[]int中的值
func (obj *Object) DenyInt(i []int, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value, err := strconv.Atoi(obj.rawValue)
	if err != nil {
		obj.err = obj.setError(customError...)
		return obj
	}
	for k := range i {
		if value == i[k] {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	return obj
}

// DenyInt32 阻止存在于[]int32中的值
func (obj *Object) DenyInt32(i []int32, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value64, err := strconv.ParseInt(obj.rawValue, 10, 32)
	if err != nil {
		obj.err = obj.setError(customError...)
		return obj
	}
	value32 := int32(value64)
	for k := range i {
		if value32 == i[k] {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	return obj
}

// DenyInt64 阻止存在于[]int64中的值
func (obj *Object) DenyInt64(i []int64, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if obj.i64 == 0 {
		var err error
		obj.i64, err = strconv.ParseInt(obj.rawValue, 10, 64)
		if err != nil {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	for k := range i {
		if obj.i64 == i[k] {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	return obj
}

// DenyFloat32 阻止存在于[]float32中的值
func (obj *Object) DenyFloat32(f []float32, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	value, err := strconv.ParseFloat(obj.rawValue, 32)
	if err != nil {
		obj.err = obj.setError(customError...)
		return obj
	}
	value32 := float32(value)
	for k := range f {
		if value32 == f[k] {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	return obj
}

// DenyFloats64 阻止存在于[]float64中的值
func (obj *Object) DenyInFloat64(f []float64, customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	if obj.f64 == 0 {
		var err error
		obj.f64, err = strconv.ParseFloat(obj.rawValue, 64)
		if err != nil {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	for k := range f {
		if obj.f64 == f[k] {
			obj.err = obj.setError(customError...)
			return obj
		}
	}
	return obj
}
