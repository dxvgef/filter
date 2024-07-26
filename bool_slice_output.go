package filter

import "reflect"

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
func (boolSliceType *BooleanSliceType) Set(target interface{}, customError ...string) error {
	if boolSliceType.err != nil {
		return boolSliceType.err
	}
	targetValueOf, checkErr := setCheck(target)
	if checkErr != nil {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
		return boolSliceType.err
	}

	// 开始赋值
	sliceType := targetValueOf.Elem().Type().String()
	if sliceType != "[]bool" {
		boolSliceType.err = wrapError(boolSliceType.name, customError...)
		return boolSliceType.err
	}
	targetValueOf.Elem().Set(reflect.ValueOf(boolSliceType.Value()))
	return boolSliceType.err
}
