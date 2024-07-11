package filter

// Equal 必须等于指定的值
func (boolType *BooleanType) Equal(value bool, customError ...string) *BooleanType {
	if boolType.err != nil {
		return boolType
	}
	if boolType.value != value {
		boolType.err = wrapError(boolType.name, customError...)
		return boolType
	}
	return boolType
}

// NotEqual 不能等于指定的值
func (boolType *BooleanType) NotEqual(value bool, customError ...string) *BooleanType {
	if boolType.err != nil {
		return boolType
	}
	if boolType.value == value {
		boolType.err = wrapError(boolType.name, customError...)
		return boolType
	}
	return boolType
}
