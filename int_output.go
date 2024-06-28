package filter

import (
	"math"
	"reflect"
	"unsafe"
)

// Result 获得过滤结果
func (intType *IntegerType) Result() (int64, error) {
	return intType.value, intType.err
}

// DefaultValue 如果校验失败则返回默认值
func (intType *IntegerType) DefaultValue(def int64) int64 {
	if intType.err != nil || intType.value == 0 {
		intType.err = nil
		return def
	}
	return intType.value
}

// Error 获得错误信息
func (intType *IntegerType) Error() error {
	return intType.err
}

// Set 使用反射赋值到变量
func (intType *IntegerType) Set(target interface{}, customError ...string) error {
	if intType.err != nil {
		return intType.err
	}
	targetValueOf, checkErr := setCheck(target)
	if checkErr != nil {
		intType.err = wrapError(intType.name, customError...)
		return intType.err
	}

	// 开始赋值
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	if targetTypeOf != reflect.Int64 {
		intType.err = wrapError(intType.name, customError...)
		return intType.err
	}
	targetValueOf.Elem().SetInt(intType.value)
	return intType.err
}

// Int8 转为int8类型
func (intType *IntegerType) Int8(customError ...string) (value int8, err error) {
	if intType.err != nil {
		return 0, intType.err
	}
	if intType.value < math.MinInt8 || intType.value > math.MaxInt8 {
		intType.err = wrapError(intType.name, customError...)
		return 0, intType.err
	}
	return int8(intType.value), nil
}

// DefaultInt8 转为int8类型，出错则用默认值替代
func (intType *IntegerType) DefaultInt8(def int8) int8 {
	if intType.err != nil {
		return def
	}
	if intType.value < math.MinInt8 || intType.value > math.MaxInt8 {
		return def
	}
	return int8(intType.value)
}

// Int16 转为int16类型
func (intType *IntegerType) Int16(customError ...string) (value int16, err error) {
	if intType.err != nil {
		return 0, intType.err
	}
	if intType.value < math.MinInt16 || intType.value > math.MaxInt16 {
		intType.err = wrapError(intType.name, customError...)
		return 0, intType.err
	}
	return int16(intType.value), nil
}

// DefaultInt16 转为int16类型，出错则用默认值替代
func (intType *IntegerType) DefaultInt16(def int16) int16 {
	if intType.err != nil {
		return def
	}
	if intType.value < math.MinInt16 || intType.value > math.MaxInt16 {
		return def
	}
	return int16(intType.value)
}

// Int32 转为int32类型
func (intType *IntegerType) Int32(customError ...string) (value int32, err error) {
	if intType.err != nil {
		return 0, intType.err
	}
	if intType.value < math.MinInt32 || intType.value > math.MaxInt32 {
		intType.err = wrapError(intType.name, customError...)
		return 0, intType.err
	}
	return int32(intType.value), nil
}

// DefaultInt32 转为int32类型，出错则用默认值替代
func (intType *IntegerType) DefaultInt32(def int32) int32 {
	if intType.err != nil {
		return def
	}
	if intType.value < math.MinInt32 || intType.value > math.MaxInt32 {
		return def
	}
	return int32(intType.value)
}

// Int64 转为int64类型
func (intType *IntegerType) Int64() int64 {
	return intType.value
}

// DefaultInt64 转为int64类型，出错则用默认值替代
func (intType *IntegerType) DefaultInt64(def int64) int64 {
	if intType.err != nil {
		return def
	}
	return intType.value
}

// Int 转为int类型
func (intType *IntegerType) Int(customError ...string) (value int, err error) {
	if intType.err != nil {
		return 0, intType.err
	}
	if intType.value < math.MinInt || intType.value > math.MaxInt {
		intType.err = wrapError(intType.name, customError...)
		return 0, intType.err
	}
	return int(intType.value), nil
}

// DefaultInt 转为int类型，出错则用默认值替代
func (intType *IntegerType) DefaultInt(def int) int {
	if intType.err != nil {
		return def
	}
	if intType.value < math.MinInt || intType.value > math.MaxInt {
		return def
	}
	return int(intType.value)
}

// Uint8 转为uint8类型
func (intType *IntegerType) Uint8(customError ...string) (value uint8, err error) {
	if intType.err != nil {
		return 0, intType.err
	}
	if intType.value < 0 || intType.value > math.MaxUint8 {
		intType.err = wrapError(intType.name, customError...)
		return 0, intType.err
	}
	return uint8(intType.value), nil
}

// DefaultUint8 转为uint8类型，出错则用默认值替代
func (intType *IntegerType) DefaultUint8(def uint8) uint8 {
	if intType.err != nil {
		return def
	}
	if intType.value > math.MaxUint8 {
		return def
	}
	return uint8(intType.value)
}

// Uint16 转为uint16类型
func (intType *IntegerType) Uint16(customError ...string) (value uint16, err error) {
	if intType.err != nil {
		return 0, intType.err
	}
	if intType.value < 0 || intType.value > math.MaxUint16 {
		intType.err = wrapError(intType.name, customError...)
		return 0, intType.err
	}
	return uint16(intType.value), nil
}

// DefaultUint16 转为uint16类型，出错则用默认值替代
func (intType *IntegerType) DefaultUint16(def uint16) uint16 {
	if intType.err != nil {
		return def
	}
	if intType.value > math.MaxUint16 {
		return def
	}
	return uint16(intType.value)
}

// Uint32 转为uint32类型
func (intType *IntegerType) Uint32(customError ...string) (value uint32, err error) {
	if intType.err != nil {
		return 0, intType.err
	}
	if intType.value > math.MaxUint32 {
		intType.err = wrapError(intType.name, customError...)
		return 0, intType.err
	}
	return uint32(intType.value), nil
}

// DefaultUint32 转为uint32类型，出错则用默认值替代
func (intType *IntegerType) DefaultUint32(def uint32) uint32 {
	if intType.err != nil {
		return def
	}
	if intType.value > math.MaxUint32 {
		return def
	}
	return uint32(intType.value)
}

// DefaultUint64 转为uint64类型，出错则用默认值替代
func (intType *IntegerType) DefaultUint64(def uint64) uint64 {
	if intType.err != nil {
		return def
	}
	return uint64(intType.value)
}

// Uint 转为uint类型
func (intType *IntegerType) Uint(customError ...string) (value uint, err error) {
	if intType.err != nil {
		return 0, intType.err
	}
	if intType.value < 0 {
		intType.err = wrapError(intType.name, customError...)
		return 0, intType.err
	}
	return uint(intType.value), nil
}

// DefaultUint 转为uint类型，出错则用默认值替代
func (intType *IntegerType) DefaultUint(def uint) uint {
	if intType.err != nil {
		return def
	}
	if intType.value < 0 {
		return def
	}
	// 32位
	if unsafe.Sizeof(int(0)) == 4 && intType.value > math.MaxUint32 {
		return def
	}
	return uint(intType.value)
}
