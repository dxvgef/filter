package filter

import "slices"

// Equals 等于
func (intType *IntegerType) Equals(value int64, customError ...string) *IntegerType {
	if intType.value != value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// NotEquals 不等于
func (intType *IntegerType) NotEquals(value int64, customError ...string) *IntegerType {
	if intType.value == value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// LessThan 小于
func (intType *IntegerType) LessThan(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value > value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// GreaterThan 大于
func (intType *IntegerType) GreaterThan(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value < value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// Range 范围
func (intType *IntegerType) Range(minValue, maxValue int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value < minValue || intType.value > maxValue {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// AllowedValues 只能是数组中的值
func (intType *IntegerType) AllowedValues(values []int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if !slices.Contains(values, intType.value) {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// DisallowedValues 不能是数组中的值
func (intType *IntegerType) DisallowedValues(values []int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if slices.Contains(values, intType.value) {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}
