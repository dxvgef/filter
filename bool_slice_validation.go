package filter

// Contains 存在指定值的元素
func (boolSliceType *BooleanSliceType) Contains(value bool, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	for k := range boolSliceType.value {
		if value == boolSliceType.value[k] {
			return boolSliceType
		}
	}
	boolSliceType.err = wrapError(boolSliceType.name, customError...)
	return boolSliceType
}

// NotContains 不存在指定值的元素
func (boolSliceType *BooleanSliceType) NotContains(value bool, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}

	for k := range boolSliceType.value {
		if value == boolSliceType.value[k] {
			boolSliceType.err = wrapError(boolSliceType.name, customError...)
			break
		}
	}
	return boolSliceType
}

// CountEquals 元素数量等于
func (boolSliceType *BooleanSliceType) CountEquals(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) != value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}

// CountNotEquals 元素数量不等于
func (boolSliceType *BooleanSliceType) CountNotEquals(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) == value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}

// CountMax 元素数量不能大于
func (boolSliceType *BooleanSliceType) CountMax(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) > value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}

// CountMin 元素数量不能小于
func (boolSliceType *BooleanSliceType) CountMin(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) < value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}
