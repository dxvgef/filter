package filter

import "slices"

// Equals 等于
func (floatType *FloatType) Equals(value float64, customError ...string) *FloatType {
	if floatType.value != value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// NotEquals 不等于
func (floatType *FloatType) NotEquals(value float64, customError ...string) *FloatType {
	if floatType.value == value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// LessThan 小于
func (floatType *FloatType) LessThan(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value > value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// GreaterThan 大于
func (floatType *FloatType) GreaterThan(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value < value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// AllowedValues 只能是数组中的值
func (floatType *FloatType) AllowedValues(values []float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if !slices.Contains(values, floatType.value) {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// DisallowedValues 不能是数组中的值
func (floatType *FloatType) DisallowedValues(values []float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if slices.Contains(values, floatType.value) {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}
