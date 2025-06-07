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

// Enum 元素中只能使用列表中的值
func (floatSliceType *FloatSliceType) Enum(values []float64, customError ...string) *FloatSliceType {
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

// Block 元素中不能存在列表中的值
func (floatSliceType *FloatSliceType) Block(values []float64, customError ...string) *FloatSliceType {
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

// Min 元素中不能出现小于指定的值
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

// Max 元素中不能出现大于指定的值
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

// Range 元素中只能使用指定范围的值
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

// CountMax 元素数量不能大于指定的值
func (floatSliceType *FloatSliceType) CountMax(value int, customError ...string) *FloatSliceType {
	if floatSliceType.err != nil {
		return floatSliceType
	}
	if len(floatSliceType.value) > value {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType
}

// CountMin 元素数量不能小于指定的值
func (floatSliceType *FloatSliceType) CountMin(value int, customError ...string) *FloatSliceType {
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
