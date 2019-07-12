package filter

import (
	"math"
	"strconv"
)

// InStrings 存在于[]string
func (obj *Object) InStrings(slice []string, customError ...string) *Object {
	if obj.err != nil {
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

// InInt 检查int值在一个int slice中是否存在
func (obj *Object) InInt(i []int, customError ...string) *Object {
	if obj.err != nil {
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

// InInt64 检查int64值在一个int64 slice中是否存在
func (obj *Object) InInt64(i []int64, customError ...string) *Object {
	if obj.err != nil {
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

// InFloat32 检查float32值在一个float32 slice中是否存在
func (obj *Object) InFloat32(f []float32, p float64, customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	value, err := strconv.ParseFloat(obj.rawValue, 32)
	if err != nil {
		obj.err = obj.setError(err.Error(), customError...)
		return obj
	}
	value32 := float32(value)
	for k := range f {
		if floatEqual(float64(value32), float64(f[k]), p) == true {
			obj.err = obj.setError("不允许的值", customError...)
			return obj
		}
	}
	return obj
}

// InFloat64 检查float64值在一个float64 slice中是否存在
func (obj *Object) InFloat64(f []float64, p float64, customError ...string) *Object {
	if obj.err != nil {
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
		if floatEqual(obj.f64, f[k], p) == true {
			obj.err = obj.setError("不允许的值", customError...)
			return obj
		}
	}
	return obj
}

func floatEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}
