package filter

// Contains 存在
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

// NotContains 不存在
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

// CountIs 元素数量等于
func (boolSliceType *BooleanSliceType) CountIs(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) != value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}

// CountIsNot 元素数量不等于
func (boolSliceType *BooleanSliceType) CountIsNot(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) == value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}

// MaxCount 元素数量最大值
func (boolSliceType *BooleanSliceType) MaxCount(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) > value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}

// MinCount 元素数量最小值
func (boolSliceType *BooleanSliceType) MinCount(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) < value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}
