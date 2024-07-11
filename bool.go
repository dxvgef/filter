package filter

// BooleanType 布尔类型
type BooleanType struct {
	name  string // 字段名称，用于拼接错误消息
	value bool   // 原始值
	err   error  // 错误
}

// FromBoolean 输入布尔类型的值
func FromBoolean(value bool, name ...string) *BooleanType {
	boolType := &BooleanType{
		value: value,
	}
	if len(name) > 0 {
		boolType.name = name[0]
	}
	return boolType
}

/*
CustomBooleanFunc 自定义布尔类型处理函数
用StrType.Value()获得当前参数值
用StrType.Error()获得当前错误信息
出参是处理后的参数值和错误信息
*/
type CustomBooleanFunc func(*BooleanType) (bool, error)

// Custom 自定义bool处理方法
func (boolType *BooleanType) Custom(f CustomBooleanFunc) *BooleanType {
	// 如果已经存在错误，则不进行处理
	if boolType.err != nil {
		return boolType
	}

	// 执行自定义函数
	value, err := f(boolType)
	if err != nil {
		boolType.err = wrapError(boolType.name, err.Error())
		return boolType
	}

	// 获取自定义函数的返回值
	boolType.value = value
	return boolType
}
