package filter

// Has 必须存在指定的值
func (boolSliceType *BooleanSliceType) Has(value bool, customError ...string) *BooleanSliceType {
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

// Count 必须存在指定数量的元素
func (boolSliceType *BooleanSliceType) Count(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) != value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}

// MinCount 元素的数量不能小于指定值
func (boolSliceType *BooleanSliceType) MinCount(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) < value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}

// MaxCount 元素的数量不能大于指定值
func (boolSliceType *BooleanSliceType) MaxCount(value int, customError ...string) *BooleanSliceType {
	if boolSliceType.err != nil {
		return boolSliceType
	}
	if len(boolSliceType.value) > value {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
	}
	return boolSliceType
}
