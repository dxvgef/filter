package filter

import (
	"fmt"
	"strconv"
)

// MinInteger 最小整数值
func (obj *Object) MinInteger(min int64, customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	if obj.i64 == 0 {
		var err error
		obj.i64, err = strconv.ParseInt(obj.rawValue, 10, 64)
		if err != nil {
			obj.err = err
			return obj
		}
	}
	if obj.i64 < min {
		obj.err = obj.setError("不能小于"+strconv.FormatInt(min, 10), customError...)
		return obj
	}
	return obj
}

// MinInteger 最大整数值
func (obj *Object) MaxInteger(max int64, customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	if obj.i64 == 0 {
		var err error
		obj.i64, err = strconv.ParseInt(obj.rawValue, 10, 64)
		if err != nil {
			obj.err = err
			return obj
		}
	}
	if obj.i64 > max {
		obj.err = obj.setError("不能大于"+strconv.FormatInt(max, 10), customError...)
		return obj
	}
	return obj
}

// MinFloat 最小浮点值
func (obj *Object) MinFloat(min float64, customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	if obj.f64 == 0 {
		var err error
		obj.f64, err = strconv.ParseFloat(obj.rawValue, 64)
		if err != nil {
			obj.err = err
			return obj
		}
	}

	if obj.f64 < min {
		obj.err = obj.setError("不能小于"+fmt.Sprint(min), customError...)
		return obj
	}
	return obj
}

// MinFloat 最大浮点值
func (obj *Object) MaxFloat(max float64, customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	if obj.f64 == 0 {
		var err error
		obj.f64, err = strconv.ParseFloat(obj.rawValue, 64)
		if err != nil {
			obj.err = err
			return obj
		}
	}
	if obj.f64 > max {
		obj.err = obj.setError("不能大于"+fmt.Sprint(max), customError...)
		return obj
	}
	return obj
}
