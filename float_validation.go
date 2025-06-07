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

// Max 不能大于指定的值
func (floatType *FloatType) Max(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value > value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// Min 不能小于指定的值
func (floatType *FloatType) Min(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value < value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// Range 只能是指定范围的值
func (floatType *FloatType) Range(minValue, maxValue float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value < minValue || floatType.value > maxValue {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// Enum 只能是数组中的值
func (floatType *FloatType) Enum(values []float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if !slices.Contains(values, floatType.value) {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// Block 不能是指定的值
func (floatType *FloatType) Block(values []float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if slices.Contains(values, floatType.value) {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}
