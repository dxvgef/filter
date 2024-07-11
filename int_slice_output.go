package filter

import (
	"math"
	"reflect"
)

// Result 获得过滤结果
func (intSliceType *IntegerSliceType) Result() ([]int64, error) {
	return intSliceType.value, intSliceType.err
}

// Value 获得当前值
func (intSliceType *IntegerSliceType) Value() []int64 {
	return intSliceType.value
}

// Error 获得错误信息
func (intSliceType *IntegerSliceType) Error() error {
	return intSliceType.err
}

// Set 使用反射赋值到变量
// todo 未完成
func (intSliceType *IntegerSliceType) Set(target interface{}, customError ...string) error {
	if intSliceType.err != nil {
		return intSliceType.err
	}
	targetValueOf, checkErr := setCheck(target)
	if checkErr != nil {
		intSliceType.err = wrapError(intSliceType.name, customError...)
		return intSliceType.err
	}

	// 开始赋值
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	if targetTypeOf != reflect.Int64 {
		intSliceType.err = wrapError(intSliceType.name, customError...)
		return intSliceType.err
	}
	// targetValueOf.Elem().SetInt(intSliceType.value)
	return intSliceType.err
}

// Int8Slice 转为[]int8类型
func (intSliceType *IntegerSliceType) Int8Slice(customError ...string) ([]int8, error) {
	if intSliceType.err != nil {
		return nil, intSliceType.err
	}
	value := make([]int8, len(intSliceType.value))
	for k := range intSliceType.value {
		if intSliceType.value[k] < math.MinInt8 || intSliceType.value[k] > math.MaxInt8 {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			return nil, intSliceType.err
		}
		value[k] = int8(intSliceType.value[k])
	}
	return value, nil
}

// DefaultInt8Slice 转为[]int8类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultInt8Slice(def []int8) []int8 {
	if intSliceType.err != nil {
		return def
	}
	value, err := intSliceType.Int8Slice()
	if err != nil {
		return def
	}
	return value
}

// Int16Slice 转为[]int16类型
func (intSliceType *IntegerSliceType) Int16Slice(customError ...string) ([]int16, error) {
	if intSliceType.err != nil {
		return nil, intSliceType.err
	}
	value := make([]int16, len(intSliceType.value))
	for k := range intSliceType.value {
		if intSliceType.value[k] < math.MinInt16 || intSliceType.value[k] > math.MaxInt16 {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			return nil, intSliceType.err
		}
		value[k] = int16(intSliceType.value[k])
	}

	return value, nil
}

// DefaultInt16Slice 转为[]int16类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultInt16Slice(def []int16) []int16 {
	if intSliceType.err != nil {
		return def
	}
	value, err := intSliceType.Int16Slice()
	if err != nil {
		return def
	}
	return value
}

// Int32Slice 转为[]int32类型
func (intSliceType *IntegerSliceType) Int32Slice(customError ...string) ([]int32, error) {
	if intSliceType.err != nil {
		return nil, intSliceType.err
	}
	value := make([]int32, len(intSliceType.value))
	for k := range intSliceType.value {
		if intSliceType.value[k] < math.MinInt32 || intSliceType.value[k] > math.MaxInt32 {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			return nil, intSliceType.err
		}
		value[k] = int32(intSliceType.value[k])
	}

	return value, nil
}

// DefaultInt32Slice 转为[]int32类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultInt32Slice(def []int32) []int32 {
	if intSliceType.err != nil {
		return def
	}
	value, err := intSliceType.Int32Slice()
	if err != nil {
		return def
	}
	return value
}

// Int64Slice 转为[]int64类型
func (intSliceType *IntegerSliceType) Int64Slice() ([]int64, error) {
	return intSliceType.Result()
}

// DefaultInt64Slice 转为[]int64类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultInt64Slice(def []int64) []int64 {
	value, err := intSliceType.Result()
	if err != nil {
		return def
	}
	return value
}

// IntSlice 转为[]int类型
func (intSliceType *IntegerSliceType) IntSlice(customError ...string) ([]int, error) {
	if intSliceType.err != nil {
		return nil, intSliceType.err
	}
	value := make([]int, len(intSliceType.value))
	for k := range intSliceType.value {
		if intSliceType.value[k] < math.MinInt || intSliceType.value[k] > math.MaxInt {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			return nil, intSliceType.err
		}
		value[k] = int(intSliceType.value[k])
	}
	return value, nil
}

// DefaultIntSlice 转为[]int类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultIntSlice(def []int) []int {
	if intSliceType.err != nil {
		return def
	}
	value, err := intSliceType.IntSlice()
	if err != nil {
		return def
	}
	return value
}

// Uint8Slice 转为[]uint8类型
func (intSliceType *IntegerSliceType) Uint8Slice(customError ...string) ([]uint8, error) {
	if intSliceType.err != nil {
		return nil, intSliceType.err
	}
	value := make([]uint8, len(intSliceType.value))
	for k := range intSliceType.value {
		if intSliceType.value[k] < 0 || intSliceType.value[k] > math.MaxUint8 {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			return nil, intSliceType.err
		}
		value[k] = uint8(intSliceType.value[k])
	}
	return value, nil
}

// DefaultUint8Slice 转为[]uint8类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultUint8Slice(def []uint8) []uint8 {
	if intSliceType.err != nil {
		return def
	}
	value, err := intSliceType.Uint8Slice()
	if err != nil {
		return def
	}
	return value
}

// Uint16Slice 转为[]uint16类型
func (intSliceType *IntegerSliceType) Uint16Slice(customError ...string) ([]uint16, error) {
	if intSliceType.err != nil {
		return nil, intSliceType.err
	}
	value := make([]uint16, len(intSliceType.value))
	for k := range intSliceType.value {
		if intSliceType.value[k] < 0 || intSliceType.value[k] > math.MaxUint16 {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			return nil, intSliceType.err
		}
		value[k] = uint16(intSliceType.value[k])
	}
	return value, nil
}

// DefaultUint16Slice 转为[]uint16类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultUint16Slice(def []uint16) []uint16 {
	if intSliceType.err != nil {
		return def
	}
	value, err := intSliceType.Uint16Slice()
	if err != nil {
		return def
	}
	return value
}

// Uint32Slice 转为[]uint32类型
func (intSliceType *IntegerSliceType) Uint32Slice(customError ...string) ([]uint32, error) {
	if intSliceType.err != nil {
		return nil, intSliceType.err
	}
	value := make([]uint32, len(intSliceType.value))
	for k := range intSliceType.value {
		if intSliceType.value[k] < 0 || intSliceType.value[k] > math.MaxUint32 {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			return nil, intSliceType.err
		}
		value[k] = uint32(intSliceType.value[k])
	}
	return value, nil
}

// DefaultUint32Slice 转为[]uint32类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultUint32Slice(def []uint32) []uint32 {
	if intSliceType.err != nil {
		return def
	}
	value, err := intSliceType.Uint32Slice()
	if err != nil {
		return def
	}
	return value
}

// Uint64Slice 转为[]uint64类型
func (intSliceType *IntegerSliceType) Uint64Slice(customError ...string) ([]uint64, error) {
	if intSliceType.err != nil {
		return nil, intSliceType.err
	}
	value := make([]uint64, len(intSliceType.value))
	for k := range intSliceType.value {
		if intSliceType.value[k] < 0 || intSliceType.value[k] > math.MaxUint32 {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			return nil, intSliceType.err
		}
		value[k] = uint64(intSliceType.value[k])
	}
	return value, nil
}

// DefaultUint64Slice 转为[]uint64类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultUint64Slice(def []uint64) []uint64 {
	value, err := intSliceType.Uint64Slice()
	if err != nil {
		return def
	}
	return value
}

// UintSlice 转为[]uint类型
func (intSliceType *IntegerSliceType) UintSlice(customError ...string) ([]uint, error) {
	if intSliceType.err != nil {
		return nil, intSliceType.err
	}
	value := make([]uint, len(intSliceType.value))
	for k := range intSliceType.value {
		if intSliceType.value[k] < 0 || intSliceType.value[k] > math.MaxInt {
			intSliceType.err = wrapError(intSliceType.name, customError...)
			return nil, intSliceType.err
		}
		value[k] = uint(intSliceType.value[k])
	}
	return value, nil
}

// DefaultUintSlice 转为[]uint类型，出错则用默认值替代
func (intSliceType *IntegerSliceType) DefaultUintSlice(def []uint) []uint {
	if intSliceType.err != nil {
		return def
	}
	value, err := intSliceType.UintSlice()
	if err != nil {
		return def
	}
	return value
}
