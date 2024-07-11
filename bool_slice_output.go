package filter

import (
	"reflect"
)

// Result 获得过滤结果
func (boolSliceType *BooleanSliceType) Result() ([]bool, error) {
	return boolSliceType.value, boolSliceType.err
}

// Value 获得当前值
func (boolSliceType *BooleanSliceType) Value() []bool {
	return boolSliceType.value
}

// DefaultValue 获得当前值，如果出错则输出默认值
func (boolSiceType *BooleanSliceType) DefaultValue(def []bool) []bool {
	if boolSiceType.err != nil {
		return def
	}
	return boolSiceType.value
}

// Error 获得错误信息
func (boolSiceType *BooleanSliceType) Error() error {
	return boolSiceType.err
}

// Set 使用反射赋值到变量
// todo 未完成
func (boolSiceType *BooleanSliceType) Set(target interface{}, customError ...string) error {
	if boolSiceType.err != nil {
		return boolSiceType.err
	}
	targetValueOf, checkErr := setCheck(target)
	if checkErr != nil {
		boolSiceType.err = wrapError(boolSiceType.name, customError...)
		return boolSiceType.err
	}

	// 开始赋值
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	if targetTypeOf != reflect.Bool {
		boolSiceType.err = wrapError(boolSiceType.name, customError...)
		return boolSiceType.err
	}
	// targetValueOf.Elem().Set(boolSiceType.value)
	return boolSiceType.err
}
