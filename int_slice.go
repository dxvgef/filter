package filter

// IntegerSliceType 整数切片类型
type IntegerSliceType struct {
	name  string  // 字段名称，用于拼接错误消息
	value []int64 // 原始值
	err   error   // 错误
}

// FromIntegerSlice 输入整数切片类型的值
func FromIntegerSlice(value []int64, name ...string) *IntegerSliceType {
	intSliceType := &IntegerSliceType{
		value: value,
	}
	if len(name) > 0 {
		intSliceType.name = name[0]
	}
	return intSliceType
}

/*
CustomIntegerSliceFunc 自定义整数切片类型处理函数
用StrType.[]Int64()获得当前参数值
用StrType.Error()获得当前错误信息
出参是处理后的参数值和错误信息
*/
type CustomIntegerSliceFunc func(*IntegerSliceType) ([]int64, error)

// Custom 自定义[]int64处理方法
func (intSliceType *IntegerSliceType) Custom(f CustomIntegerSliceFunc) *IntegerSliceType {
	// 如果已经存在错误，则不进行处理
	if intSliceType.err != nil {
		return intSliceType
	}

	// 执行自定义函数
	value, err := f(intSliceType)
	if err != nil {
		intSliceType.err = wrapError(intSliceType.name, err.Error())
		return intSliceType
	}

	// 获取自定义函数的返回值
	intSliceType.value = value
	return intSliceType
}
