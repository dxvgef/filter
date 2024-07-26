package filter

import "slices"

// Require 不能为零值
func (floatSliceType *FloatSliceType) Require(customError ...string) *FloatSliceType {
	if len(floatSliceType.value) == 0 {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
		return floatSliceType
	}
	for k := range floatSliceType.value {
		if floatSliceType.value[k] != 0 {
			return floatSliceType
		}
	}
	floatSliceType.err = wrapError(floatSliceType.name, customError...)
	return floatSliceType
}

// MinCount 元素数量不能小于指定值
func (floatSliceType *FloatSliceType) MinCount(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) < value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// MaxCount 元素数量不能大于指定值
func (floatSliceType *FloatSliceType) MaxCount(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) > value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// EqualCount 元素数量必须等于指定值
func (floatSliceType *FloatSliceType) EqualCount(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) != value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// NotEqualCount 元素数量不能等于指定值
func (floatSliceType *FloatSliceType) NotEqualCount(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) == value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// MinValue 切片中不能存在小于指定值的元素
func (floatSliceType *FloatSliceType) MinValue(value float64, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	for k := range floatSliceType.value {
		if floatSliceType.value[k] < value {
			floatSliceType.err = wrapError(floatSliceType.name, customError...)
			break
		}
	}
	return floatSliceType
}

// MaxValue 切片中不能存在大于指定值的元素
func (floatSliceType *FloatSliceType) MaxValue(value float64, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	for k := range floatSliceType.value {
		if floatSliceType.value[k] > value {
			floatSliceType.err = wrapError(floatSliceType.name, customError...)
			break
		}
	}
	return floatSliceType
}

// AllowedValues 只允许存在指定的值
func (floatSliceType *FloatSliceType) AllowedValues(values []float64, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	for k := range floatSliceType.value {
		if !slices.Contains(values, floatSliceType.value[k]) {
			floatSliceType.err = wrapError(floatSliceType.name, customError...)
			break
		}
	}
	return floatSliceType
}

// DeniedValues 禁止存在指定的值
func (floatSliceType *FloatSliceType) DeniedValues(values []float64, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	for k := range floatSliceType.value {
		if slices.Contains(values, floatSliceType.value[k]) {
			floatSliceType.err = wrapError(floatSliceType.name, customError...)
			break
		}
	}
	return floatSliceType
}
