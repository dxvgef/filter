package filter

import "slices"

// Require 长度不能为0
func (intSliceType *IntegerSliceType) Require(customError ...string) *IntegerSliceType {
	if len(intSliceType.value) == 0 {
		intSliceType.err = wrapError(intSliceType.name, customError...)
		return intSliceType
	}
	intSliceType.err = wrapError(intSliceType.name, customError...)
	return intSliceType
}

// MinCount 元素数量不能小于指定值
func (intSliceType *IntegerSliceType) MinCount(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) < value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// MaxCount 元素数量不能大于指定值
func (intSliceType *IntegerSliceType) MaxCount(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) > value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// EqualCount 元素数量必须等于指定值
func (intSliceType *IntegerSliceType) EqualCount(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) != value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// NotEqualCount 元素数量不能等于指定值
func (intSliceType *IntegerSliceType) NotEqualCount(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) == value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// MinValue 切片中不能存在小于指定值的元素
func (intSliceType *IntegerSliceType) MinValue(value int64, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	for k := range intSliceType.value {
		if intSliceType.value[k] < value {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			break
		}
	}
	return intSliceType
}

// MaxValue 切片中不能存在大于指定值的元素
func (intSliceType *IntegerSliceType) MaxValue(value int64, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	for k := range intSliceType.value {
		if intSliceType.value[k] > value {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			break
		}
	}
	return intSliceType
}

// AllowedValues 只允许存在指定的值
func (intSliceType *IntegerSliceType) AllowedValues(values []int64, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	for k := range intSliceType.value {
		if !slices.Contains(values, intSliceType.value[k]) {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			break
		}
	}
	return intSliceType
}

// DeniedValues 禁止存在指定的值
func (intSliceType *IntegerSliceType) DeniedValues(values []int64, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	for k := range intSliceType.value {
		if slices.Contains(values, intSliceType.value[k]) {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			break
		}
	}
	return intSliceType
}
