package filter

// StringType 字符串类型
type StringType struct {
	name  string // 字段名称，用于拼接错误消息
	value string // 原始值
	err   error  // 错误
}

// FromString 输入string类型的值
func FromString(value string, name ...string) *StringType {
	strType := &StringType{
		value: value,
	}
	if len(name) > 0 {
		strType.name = name[0]
	}
	return strType
}

// Require 不能为零值
func (strType *StringType) Require(customError ...string) *StringType {
	if strType.value == "" {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}
	return strType
}

/*
CustomStringFunc 自定义字符串处理函数
用StrType.Int64()获得当前参数值
用StrType.Error()获得当前错误信息
出参是处理后的参数值和错误信息
*/
type CustomStringFunc func(*StringType) (string, error)

// Custom 自定义string处理方法
func (strType *StringType) Custom(f CustomStringFunc) *StringType {
	// 如果已经存在错误，则不进行处理
	if strType.err != nil {
		return strType
	}

	// 执行自定义函数
	value, err := f(strType)
	if err != nil {
		strType.err = wrapError(strType.name, err.Error())
		return strType
	}

	// 获取自定义函数的返回值
	strType.value = value
	return strType
}
