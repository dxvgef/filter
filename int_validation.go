package filter

import "slices"

// Is 等于
func (intType *IntegerType) Is(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value != value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// IsNot 不等于
func (intType *IntegerType) IsNot(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value == value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// Max 最大值
func (intType *IntegerType) Max(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value > value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// Min 最小值
func (intType *IntegerType) Min(value int64, customError ...string) *IntegerType {
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

// In 在列表中
func (intType *IntegerType) In(values []int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if !slices.Contains(values, intType.value) {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// NotIn 不在列表中
func (intType *IntegerType) NotIn(values []int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if slices.Contains(values, intType.value) {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}
