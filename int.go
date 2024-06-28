package filter

// IntegerType 整数类型
type IntegerType struct {
	name  string // 字段名称，用于拼接错误消息
	value int64  // 原始值
	err   error  // 错误
}

// FromInteger 输入整数类型的值
func FromInteger(value int64, name ...string) *IntegerType {
	intType := &IntegerType{
		value: value,
	}
	if len(name) > 0 {
		intType.name = name[0]
	}
	return intType
}

// Require 不能为零值
func (intType *IntegerType) Require(customError ...string) *IntegerType {
	if intType.value == 0 {
		intType.err = wrapError(intType.name, customError...)
		return intType
	}
	return intType
}

/*
CustomStringFunc 自定义整数处理函数
用StrType.Int64()获得当前参数值
用StrType.Error()获得当前错误信息
出参是处理后的参数值和错误信息
*/
type CustomIntegerFunc func(*IntegerType) (int64, error)

// Custom 自定义int64处理方法
func (intType *IntegerType) Custom(f CustomIntegerFunc) *IntegerType {
	// 如果已经存在错误，则不进行处理
	if intType.err != nil {
		return intType
	}

	// 执行自定义函数
	value, err := f(intType)
	if err != nil {
		intType.err = wrapError(intType.name, err.Error())
		return intType
	}

	// 获取自定义函数的返回值
	intType.value = value
	return intType
}
