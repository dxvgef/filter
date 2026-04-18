package filter

import "slices"

// MaxCount 元素数量最大值
func (intSliceType *IntegerSliceType) MaxCount(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) > value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// MinCount 元素数量最小值
func (intSliceType *IntegerSliceType) MinCount(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) < value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// CountIs 元素数量等于
func (intSliceType *IntegerSliceType) CountIs(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) != value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// CountIsNot 元素数量不等于
func (intSliceType *IntegerSliceType) CountIsNot(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) == value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// Contains 元素中包含
func (intSliceType *IntegerSliceType) Contains(values int64, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}

	if !slices.Contains(intSliceType.value, values) {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// NotContains 元素中不包含
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

// In 在列表中
func (intSliceType *IntegerSliceType) In(values []int64, customError ...string) *IntegerSliceType {
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

// NotIn 不在列表中
func (intSliceType *IntegerSliceType) NotIn(values []int64, customError ...string) *IntegerSliceType {
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
