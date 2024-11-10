package filter

import "slices"

// CountLessThan 元素数量小于
func (intSliceType *IntegerSliceType) CountLessThan(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) > value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// CountGreaterThan 元素数量大于
func (intSliceType *IntegerSliceType) CountGreaterThan(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) < value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// CountEquals 元素数量等于
func (intSliceType *IntegerSliceType) CountEquals(value int, customError ...string) *IntegerSliceType {
	if intSliceType.err != nil {
		return intSliceType
	}
	if len(intSliceType.value) != value {
		intSliceType.err = wrapError(intSliceType.name, customError...)
	}
	return intSliceType
}

// CountNotEquals 元素数量不等于
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

// MinValue 每个元素值都不能小于
func (intSliceType *IntegerSliceType) MinValue(value int64, customError ...string) *IntegerSliceType {
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

// MaxValue 每个元素值都不能大于
func (intSliceType *IntegerSliceType) MaxValue(value int64, customError ...string) *IntegerSliceType {
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

// AllowedValues 元素只能是数组中的值
func (intSliceType *IntegerSliceType) AllowedValues(values []int64, customError ...string) *IntegerSliceType {
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

// DisallowedValues 元素不能有数组中的值
func (intSliceType *IntegerSliceType) DisallowedValues(values []int64, customError ...string) *IntegerSliceType {
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
