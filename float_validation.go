package filter

import "slices"

// Require 不能为零值
func (floatType *FloatType) Require(customError ...string) *FloatType {
	if floatType.value == 0 {
		floatType.err = wrapError(floatType.name, customError...)
		return floatType
	}
	return floatType
}

// MinValue 不能小于最小值
func (floatType *FloatType) MinValue(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value < value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// MaxValue 不能大于最大值
func (floatType *FloatType) MaxValue(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value > value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// AllowedValues 只允许存在指定的值
func (floatType *FloatType) AllowedValues(values []float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if !slices.Contains(values, floatType.value) {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// DeniedValues 禁止存在指定的值
func (floatType *FloatType) DeniedValues(values []float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if slices.Contains(values, floatType.value) {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}
