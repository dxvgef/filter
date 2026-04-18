package filter

import "slices"

// Is 等于
func (floatType *FloatType) Is(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value != value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// IsNot 不等于
func (floatType *FloatType) IsNot(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value == value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// Max 最大值
func (floatType *FloatType) Max(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value > value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// Min 最小值
func (floatType *FloatType) Min(value float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value < value {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// Range 范围
func (floatType *FloatType) Range(minValue, maxValue float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if floatType.value < minValue || floatType.value > maxValue {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// In 在列表中
func (floatType *FloatType) In(values []float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if !slices.Contains(values, floatType.value) {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}

// NotIn 不在列表中
func (floatType *FloatType) NotIn(values []float64, customError ...string) *FloatType {
	if floatType.err != nil {
		return floatType
	}
	if slices.Contains(values, floatType.value) {
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType
}
