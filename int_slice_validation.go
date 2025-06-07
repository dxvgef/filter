package filter

import "slices"

// CountMax 元素数量不能大于指定的值
func (intSliceType *IntegerSliceType) CountMax(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) > value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// CountMin 元素数量不能小于指定的值
func (intSliceType *IntegerSliceType) CountMin(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) < value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// CountEquals 元素数量必须等于指定的值
func (intSliceType *IntegerSliceType) CountEquals(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) != value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// CountNotEquals 元素数量不能等于指定的值
func (intSliceType *IntegerSliceType) CountNotEquals(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) == value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// Contains 元素中存在指定的值
func (intSliceType *IntegerSliceType) Contains(values int64, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}

	if !slices.Contains(intSliceType.value, values) {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// NotContains 元素中不存在指定的值
func (intSliceType *IntegerSliceType) NotContains(values int64, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}

	if slices.Contains(intSliceType.value, values) {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// Min 每个元素值都不能小于
func (intSliceType *IntegerSliceType) Min(value int64, customError ...string) *IntegerSliceType {
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

// Max 每个元素值都不能大于
func (intSliceType *IntegerSliceType) Max(value int64, customError ...string) *IntegerSliceType {
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

// Range 每个元素值的范围
func (intSliceType *IntegerSliceType) Range(minValue, maxValue int64, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	for k := range intSliceType.value {
		if intSliceType.value[k] < minValue || intSliceType.value[k] > maxValue {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			break
		}
	}
	return intSliceType
}

// Enum 元素只能是数组中的值
func (intSliceType *IntegerSliceType) Enum(values []int64, customError ...string) *IntegerSliceType {
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

// Block 元素不能有数组中的值
func (intSliceType *IntegerSliceType) Block(values []int64, customError ...string) *IntegerSliceType {
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
