package filter

import "slices"

// Contains 存在指定值的元素
func (floatSliceType *FloatSliceType) Contains(values float64, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}

	if !slices.Contains(floatSliceType.value, values) {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// NotContains 不存在指定值的元素
func (floatSliceType *FloatSliceType) NotContains(values float64, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}

	if slices.Contains(floatSliceType.value, values) {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// AllowedValues 每个元素都只能是列表中的值
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

// DisallowedValues 每个元素都不能是列表中的值
func (floatSliceType *FloatSliceType) DisallowedValues(values []float64, customError ...string) *FloatSliceType {
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

// MinValue 每个元素值都不能小于
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

// MaxValue 每个元素值都不能大于
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

// RangeValue 每个元素值的范围
func (floatSliceType *FloatSliceType) RangeValue(minValue, maxValue float64, customError ...string) *FloatSliceType {
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

// CountLessThan 元素数量小于
func (floatSliceType *FloatSliceType) CountLessThan(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) > value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// CountGreaterThan 元素数量大于
func (floatSliceType *FloatSliceType) CountGreaterThan(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) < value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// CountEquals 元素数量等于
func (floatSliceType *FloatSliceType) CountEquals(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) != value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// CountNotEquals 元素数量不等于
func (floatSliceType *FloatSliceType) CountNotEquals(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) == value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}
