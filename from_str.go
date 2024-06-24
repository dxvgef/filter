package filter

// StringType 字符串类型
type StringType struct {
	name         string // 字段名称，用于拼接错误消息
	rawValue     string // 原始值
	currentValue string // 当前值
	err          error  // 错误
}

// FromString 输入string类型的值
func FromString(value string, name ...string) *StringType {
	str := &StringType{
		rawValue:     value,
		currentValue: value,
	}
	if len(name) > 0 {
		str.name = name[0]
	}
	return str
}

// Require 不能为零值
func (strType *StringType) Require(customError ...string) *StringType {
	if strType.currentValue == "" {
		strType.err = wrapError(strType.name, customError...)
		return strType
	}
	return strType
}

// Result 获得过滤结果
func (strType *StringType) Result() (string, error) {
	return strType.currentValue, strType.err
}

// Value 获得当前参数值
func (strType *StringType) Value() string {
	return strType.currentValue
}

// Error 获得错误信息
func (strType *StringType) Error() error {
	return strType.err
}

/*
CustomStringFunc 自定义字符串处理函数
用*Str.Value()获得当前参数值
用*StringType.Error()获得当前错误信息
出参是处理后的参数值和错误信息
*/
type CustomStringFunc func(*StringType) (string, error)

// Custom 自定义string处理方法
func (strType *StringType) Custom(f CustomStringFunc) *StringType {
	// 如果已经存在错误，则不进行处理，直接抛出错误
	if strType.err != nil {
		return strType
	}
	// 执行自定义函数
	value, err := f(strType)
	if err != nil {
		// 如果自定义函数返回了错误，则抛出错误
		strType.err = wrapError(strType.name, err.Error())
		return strType
	}
	// 否则获取自定义函数的返回值
	strType.currentValue = value
	return strType
}

// // Abort 是否要继续处理
// func (strType *StringType) Abort() error {
// 	return strType.err
// }
