package filter

import "slices"

// Require 不能为零值
func (intType *IntegerType) Require(customError ...string) *IntegerType {
	if intType.value == 0 {
		intType.err = wrapError(intType.name, customError...)
		return intType
	}
	return intType
}

// MinValue 不能小于最小值
func (intType *IntegerType) MinValue(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value < value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// MaxValue 不能大于最大值
func (intType *IntegerType) MaxValue(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value > value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// AllowedValues 只允许存在指定的值
func (intType *IntegerType) AllowedValues(values []int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if !slices.Contains(values, intType.value) {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// DeniedValues 禁止存在指定的值
func (intType *IntegerType) DeniedValues(values []int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if slices.Contains(values, intType.value) {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}
