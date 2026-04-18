package filter

import "slices"

// Contains 包含
func (floatSliceType *FloatSliceType) Contains(values float64, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}

	if !slices.Contains(floatSliceType.value, values) {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// NotContains 不包含
func (floatSliceType *FloatSliceType) NotContains(values float64, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}

	if slices.Contains(floatSliceType.value, values) {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// In 在列表中
func (floatSliceType *FloatSliceType) In(values []float64, customError ...string) *FloatSliceType {
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

// NotIn 不在列表中
func (floatSliceType *FloatSliceType) NotIn(values []float64, customError ...string) *FloatSliceType {
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

// Min 最小值
func (floatSliceType *FloatSliceType) Min(value float64, customError ...string) *FloatSliceType {
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

// Max 最大值
func (floatSliceType *FloatSliceType) Max(value float64, customError ...string) *FloatSliceType {
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

// Range 范围
func (floatSliceType *FloatSliceType) Range(minValue, maxValue float64, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	for k := range floatSliceType.value {
		if floatSliceType.value[k] < minValue || floatSliceType.value[k] > maxValue {
			floatSliceType.err = wrapError(floatSliceType.name, customError...)
			break
		}
	}
	return floatSliceType
}

// MaxCount 元素数量最大值
func (floatSliceType *FloatSliceType) MaxCount(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) > value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// MinCount 元素数量最小值
func (floatSliceType *FloatSliceType) MinCount(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) < value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// CountIs 元素数量等于
func (floatSliceType *FloatSliceType) CountIs(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) != value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// CountIsNot 元素数量不等于
func (floatSliceType *FloatSliceType) CountIsNot(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) == value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}
