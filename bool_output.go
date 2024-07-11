package filter

import (
	"reflect"
)

// Result 获得过滤结果
func (boolType *BooleanType) Result() (bool, error) {
	return boolType.value, boolType.err
}

// Value 获得当前值
func (boolType *BooleanType) Value() bool {
	return boolType.value
}

// DefaultValue 获得当前值，如果出错则输出默认值
func (boolType *BooleanType) DefaultValue(def bool) bool {
	if boolType.err != nil {
		return def
	}
	return boolType.value
}

// Error 获得错误信息
func (boolType *BooleanType) Error() error {
	return boolType.err
}

// Set 使用反射赋值到变量
func (boolType *BooleanType) Set(target interface{}, customError ...string) error {
	if boolType.err != nil {
		return boolType.err
	}
	targetValueOf, checkErr := setCheck(target)
	if checkErr != nil {
		boolType.err = wrapError(boolType.name, customError...)
		return boolType.err
	}

	// 开始赋值
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	if targetTypeOf != reflect.Bool {
		boolType.err = wrapError(boolType.name, customError...)
		return boolType.err
	}
	targetValueOf.Elem().SetBool(boolType.value)
	return boolType.err
}
