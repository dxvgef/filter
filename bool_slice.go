package filter

// BooleanSliceType 布尔切片类型
type BooleanSliceType struct {
	name  string // 字段名称，用于拼接错误消息
	value []bool // 原始值
	err   error  // 错误
}

// FromBooleanSlice 输入布尔切片类型的值
func FromBooleanSlice(value []bool, name ...string) *BooleanSliceType {
	boolSliceType := &BooleanSliceType{
		value: value,
	}
	if len(name) > 0 {
		boolSliceType.name = name[0]
	}
	return boolSliceType
}

/*
CustomBooleanSliceFunc 自定义布尔切片类型处理函数
用StrType.Value()获得当前参数值
用StrType.Error()获得当前错误信息
出参是处理后的参数值和错误信息
*/
type CustomBooleanSliceFunc func(*BooleanSliceType) ([]bool, error)

// Custom 自定义[]bool处理方法
func (boolSliceType *BooleanSliceType) Custom(f CustomBooleanSliceFunc) *BooleanSliceType {
	// 如果已经存在错误，则不进行处理
	if boolSliceType.err != nil {
		return boolSliceType
	}

	// 执行自定义函数
	value, err := f(boolSliceType)
	if err != nil {
		boolSliceType.err = wrapError(boolSliceType.name, err.Error())
		return boolSliceType
	}

	// 获取自定义函数的返回值
	boolSliceType.value = value
	return boolSliceType
}
