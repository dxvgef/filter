package filter

import (
	"reflect"
)

// Result 获得过滤结果
func (strType *StringType) Result() (string, error) {
	return strType.value, strType.err
}

// Value 获得当前参数值
func (strType *StringType) Value() string {
	return strType.value
}

// DefaultValue 如果校验失败则返回默认值
func (strType *StringType) DefaultValue(def string) string {
	if strType.err != nil || strType.value == "" {
		strType.err = nil
		return def
	}
	return strType.value
}

// Error 获得错误信息
func (strType *StringType) Error() error {
	return strType.err
}

// Set 使用反射赋值到变量
func (strType *StringType) Set(target interface{}, customError ...string) error {
	if strType.err != nil {
		return strType.err
	}
	targetValueOf, checkErr := setCheck(target)
	if checkErr != nil {
		strType.err = wrapError(strType.name, customError...)
		return strType.err
	}

	// 开始赋值
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	if targetTypeOf != reflect.String {
		strType.err = wrapError(strType.name, customError...)
		return strType.err
	}
	targetValueOf.Elem().SetString(strType.value)
	return strType.err
}
