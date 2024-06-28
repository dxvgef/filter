package filter

// Min 不能最小值
func (intType *IntegerType) Min(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value < value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// Max 不能最大值
func (intType *IntegerType) Max(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value > value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// Equal 必须等于值
func (intType *IntegerType) Equal(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value != value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// NotEqual 不能等于值
func (intType *IntegerType) NotEqual(value int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value == value {
		intType.err = wrapError(intType.name, customError...)
	}
	return intType
}

// AllowedValues 只允许存在指定的值
func (intType *IntegerType) AllowedValues(values []int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	for k := range values {
		if intType.value == values[k] {
			return intType
		}
	}
	intType.err = wrapError(intType.name, customError...)
	return intType
}

// DeniedValues 禁止存在指定的值
func (intType *IntegerType) DeniedValues(values []int64, customError ...string) *IntegerType {
	if intType.err != nil {
		return intType
	}
	for k := range values {
		if intType.value == values[k] {
			intType.err = wrapError(intType.name, customError...)
			return intType
		}
	}
	return intType
}
