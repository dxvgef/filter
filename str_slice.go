package filter

// StringSliceType 字符串切片类型
type StringSliceType struct {
	name  string   // 字段名称，用于拼接错误消息
	value []string // 原始值
	err   error    // 错误
}

// FromStringSlice 输入[]string类型的值
func FromStringSlice(value []string, name ...string) *StringSliceType {
	strType := &StringSliceType{
		value: value,
	}
	if len(name) > 0 {
		strType.name = name[0]
	}
	return strType
}

/*
CustomStringSliceFunc 自定义字符串切片处理函数
用StrType.Value()获得当前参数值
用StrType.Error()获得当前错误信息
出参是处理后的参数值和错误信息
*/
type CustomStringSliceFunc func(*StringSliceType) ([]string, error)

// Custom 自定义[]string处理方法
func (strSliceType *StringSliceType) Custom(f CustomStringSliceFunc) *StringSliceType {
	// 如果已经存在错误，则不进行处理
	if strSliceType.err != nil {
		return strSliceType
	}

	// 执行自定义函数
	value, err := f(strSliceType)
	if err != nil {
		strSliceType.err = wrapError(strSliceType.name, err.Error())
		return strSliceType
	}

	// 获取自定义函数的返回值
	strSliceType.value = value
	return strSliceType
}
