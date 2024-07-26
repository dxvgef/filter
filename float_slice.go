package filter

// FloatSliceType 整数切片类型
type FloatSliceType struct {
	name  string    // 字段名称，用于拼接错误消息
	value []float64 // 原始值
	err   error     // 错误
}

// FromFloatSlice 输入整数类型的值
func FromFloatSlice(value []float64, name ...string) *FloatSliceType {
	floatSliceType := &FloatSliceType{
		value: value,
	}
	if len(name) > 0 {
		floatSliceType.name = name[0]
	}
	return floatSliceType
}

/*
CustomFloatSliceFunc 自定义浮点切片处理函数
用FloatSliceType.Value()获得当前参数值
用FloatSliceType.Error()获得当前错误信息
出参是处理后的参数值和错误信息
*/
type CustomFloatSliceFunc func(*FloatSliceType) ([]float64, error)

// Custom 自定义[]float处理方法
func (floatSliceType *FloatSliceType) Custom(f CustomFloatSliceFunc) *FloatSliceType {
	// 如果已经存在错误，则不进行处理
	if floatSliceType.err != nil {
		return floatSliceType
	}

	// 执行自定义函数
	value, err := f(floatSliceType)
	if err != nil {
		floatSliceType.err = wrapError(floatSliceType.name, err.Error())
		return floatSliceType
	}

	// 获取自定义函数的返回值
	floatSliceType.value = value
	return floatSliceType
}
