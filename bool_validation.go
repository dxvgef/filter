package filter

// Is 等于
func (boolType *BooleanType) Is(value bool, customError ...string) *BooleanType {
	if boolType.err != nil {
		return boolType
	}
	if boolType.value != value {
		boolType.err = wrapError(boolType.name, customError...)
		return boolType
	}
	return boolType
}

// IsNot 不等于
func (boolType *BooleanType) IsNot(value bool, customError ...string) *BooleanType {
	if boolType.err != nil {
		return boolType
	}
	if boolType.value == value {
		boolType.err = wrapError(boolType.name, customError...)
		return boolType
	}
	return boolType
}
