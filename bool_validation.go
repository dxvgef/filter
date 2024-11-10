package filter

// Equals 等于指定的值
func (boolType *BooleanType) Equals(value bool, customError ...string) *BooleanType {
	if boolType.err != nil {
		return boolType
	}
	if boolType.value != value {
		boolType.err = wrapError(boolType.name, customError...)
		return boolType
	}
	return boolType
}

// NotEquals 不等于指定的值
func (boolType *BooleanType) NotEquals(value bool, customError ...string) *BooleanType {
	if boolType.err != nil {
		return boolType
	}
	if boolType.value == value {
		boolType.err = wrapError(boolType.name, customError...)
		return boolType
	}
	return boolType
}
