package filter

import (
	"reflect"
)

// Result 获得过滤结果
func (strSliceType *StringSliceType) Result() ([]string, error) {
	return strSliceType.value, strSliceType.err
}

// Value 获得当前参数值
func (strSliceType *StringSliceType) Value() []string {
	return strSliceType.value
}

// DefaultValue 如果校验失败则返回默认值
func (strSliceType *StringSliceType) DefaultValue(def []string) []string {
	if strSliceType.err != nil || len(strSliceType.value) == 0 {
		strSliceType.err = nil
		return def
	}
	return strSliceType.value
}

// Error 获得错误信息
func (strSliceType *StringSliceType) Error() error {
	return strSliceType.err
}

// Set 使用反射赋值到变量
// todo:未完成
func (strSliceType *StringSliceType) Set(target interface{}, customError ...string) error {
	if strSliceType.err != nil {
		return strSliceType.err
	}
	targetValueOf, checkErr := setCheck(target)
	if checkErr != nil {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType.err
	}

	// 开始赋值
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	if targetTypeOf != reflect.String {
		strSliceType.err = wrapError(strSliceType.name, customError...)
		return strSliceType.err
	}
	// targetValueOf.Elem().SetString(strSliceType.value)
	return strSliceType.err
}
