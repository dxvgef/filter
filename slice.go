package filter

import (
	"math"
	"strconv"
)

// InStrings 检查string是否存在于[]string中
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

// InInts 检查int是否存在一个[]int中
func (obj *Object) InInts(i []int, customError ...string) *Object {
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

// InInts32 检查int32是否存在一个[]int32中
func (obj *Object) InInts32(i []int32, customError ...string) *Object {
	if obj.err != nil {
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

// InInts64 检查int64是否存在一个[]int64中
func (obj *Object) InInts64(i []int64, customError ...string) *Object {
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

// InFloats32 检查float32是否丰在一个[]float32中
func (obj *Object) InFloats32(f []float32, p float64, customError ...string) *Object {
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

// InFloats64 检查float64是否存在[]float64中
func (obj *Object) InFloats64(f []float64, p float64, customError ...string) *Object {
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
