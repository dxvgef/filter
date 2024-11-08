package filter

import (
	"math"
	"reflect"
)

// Result 获得过滤结果
func (floatSliceType *FloatSliceType) Result() ([]float64, error) {
	return floatSliceType.value, floatSliceType.err
}

// Value 获得当前值
func (floatSliceType *FloatSliceType) Value() []float64 {
	return floatSliceType.value
}

// Error 获得错误信息
func (floatSliceType *FloatSliceType) Error() error {
	return floatSliceType.err
}

// Set 使用反射赋值到变量
func (floatSliceType *FloatSliceType) Set(target interface{}, customError ...string) error {
	if floatSliceType.err != nil {
		return floatSliceType.err
	}
	targetValueOf, checkErr := setCheck(target)
	if checkErr != nil {
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
		return floatSliceType.err
	}

	// 开始赋值
	sliceType := targetValueOf.Elem().Type().String()
	switch sliceType {
	case "[]float64":
		targetValueOf.Elem().Set(reflect.ValueOf(floatSliceType.Value()))
	case "[]float32":
		value, err := floatSliceType.Float32Slice()
		if err != nil {
			floatSliceType.err = wrapError(floatSliceType.name, customError...)
			return floatSliceType.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	default:
		floatSliceType.err = wrapError(floatSliceType.name, customError...)
	}
	return floatSliceType.err
}

// Float32Slice 转为[]float32类型
func (floatSliceType *FloatSliceType) Float32Slice(customError ...string) ([]float32, error) {
	if floatSliceType.err != nil {
		return nil, floatSliceType.err
	}
	value := make([]float32, len(floatSliceType.value))
	for k := range floatSliceType.value {
		if floatSliceType.value[k] < -math.SmallestNonzeroFloat32 || floatSliceType.value[k] > math.MaxFloat32 {
			floatSliceType.err = wrapError(floatSliceType.name, customError...)
			return nil, floatSliceType.err
		}
		value[k] = float32(floatSliceType.value[k])
	}
	return value, nil
}

// DefaultFloat32Slice 转为[]float32类型，出错则用默认值替代
func (floatSliceType *FloatSliceType) DefaultFloat32Slice(def []float32) []float32 {
	if floatSliceType.err != nil {
		return def
	}
	value, err := floatSliceType.Float32Slice()
	if err != nil {
		return def
	}
	return value
}

// Float64Slice 转为[]float64类型
func (floatSliceType *FloatSliceType) Float64Slice() ([]float64, error) {
	return floatSliceType.Result()
}

// DefaultFloat64Slice 转为[]float64类型，出错则用默认值替代
func (floatSliceType *FloatSliceType) DefaultFloat64Slice(def []float64) []float64 {
	value, err := floatSliceType.Result()
	if err != nil {
		return def
	}
	return value
}
