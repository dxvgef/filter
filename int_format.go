package filter // RemoveSpace 删除所有空格

// Replace 如果值等于oldValue，则替换成newValue
func (intType *IntegerType) Replace(oldValue int64, newValue int64) *IntegerType {
	if intType.err != nil {
		return intType
	}
	if intType.value == oldValue {
		intType.value = newValue
	}
	return intType
}
