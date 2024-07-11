package filter

import (
	"math"
	"reflect"
)

// Result 获得过滤结果
func (floatType *FloatType) Result() (float64, error) {
	return floatType.value, floatType.err
}

// Value 获得当前值
func (floatType *FloatType) Value() float64 {
	return floatType.value
}

// Error 获得错误信息
func (floatType *FloatType) Error() error {
	return floatType.err
}

// Set 使用反射赋值到变量
func (floatType *FloatType) Set(target interface{}, customError ...string) error {
	if floatType.err != nil {
		return floatType.err
	}
	targetValueOf, checkErr := setCheck(target)
	if checkErr != nil {
		floatType.err = wrapError(floatType.name, customError...)
		return floatType.err
	}

	// 开始赋值
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	switch targetTypeOf {
	case reflect.Float64, reflect.Float32:
		targetValueOf.Elem().SetFloat(floatType.value)
	default:
		floatType.err = wrapError(floatType.name, customError...)
	}
	return floatType.err
}

// Float32 转为float32类型
func (floatType *FloatType) Float32(customError ...string) (value float32, err error) {
	if floatType.err != nil {
		return 0, floatType.err
	}
	if floatType.value > math.MaxFloat32 {
		floatType.err = wrapError(floatType.name, customError...)
		return 0, floatType.err
	}
	return float32(floatType.value), nil
}

// DefaultFloat32 转为float32类型，出错则用默认值替代
func (floatType *FloatType) DefaultFloat32(def float32) float32 {
	value, err := floatType.Float32()
	if err != nil {
		return def
	}
	return value
}

// Float64 转为float64类型
func (floatType *FloatType) Float64() (float64, error) {
	return floatType.Result()
}

// DefaultFloat64 转为float64类型，出错则用默认值替代
func (floatType *FloatType) DefaultFloat64(def float64) float64 {
	value, err := floatType.Result()
	if err != nil {
		return def
	}
	return value
}
