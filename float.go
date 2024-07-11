package filter

// FloatType 整数类型
type FloatType struct {
	name  string  // 字段名称，用于拼接错误消息
	value float64 // 原始值
	err   error   // 错误
}

// FromFloat 输入整数类型的值
func FromFloat(value float64, name ...string) *FloatType {
	floatType := &FloatType{
		value: value,
	}
	if len(name) > 0 {
		floatType.name = name[0]
	}
	return floatType
}

/*
CustomFloatFunc 自定义浮点处理函数
用FloatType.Value()获得当前参数值
用FloatType.Error()获得当前错误信息
出参是处理后的参数值和错误信息
*/
type CustomFloatFunc func(*FloatType) (float64, error)

// Custom 自定义int64处理方法
func (floatType *FloatType) Custom(f CustomFloatFunc) *FloatType {
	// 如果已经存在错误，则不进行处理
	if floatType.err != nil {
		return floatType
	}

	// 执行自定义函数
	value, err := f(floatType)
	if err != nil {
		floatType.err = wrapError(floatType.name, err.Error())
		return floatType
	}

	// 获取自定义函数的返回值
	floatType.value = value
	return floatType
}
