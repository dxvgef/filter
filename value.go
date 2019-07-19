package filter

import (
	"strconv"
	"strings"
)

// String 转为string类型
func (obj *Object) String() (string, error) {
	if obj.err != nil || obj.rawValue == "" {
		return "", obj.err
	}
	return obj.rawValue, nil
}

// DefaultString 转为string类型，如果出错则只返回默认值
func (obj *Object) DefaultString(defaultValue string) string {
	if obj.err != nil || obj.rawValue == "" {
		return defaultValue
	}
	return obj.rawValue
}

// Strings 转为[]string类型，sep为分隔符
func (obj *Object) Strings(sep string) ([]string, error) {
	if obj.err != nil || obj.rawValue == "" {
		return []string{}, obj.err
	}
	value := strings.Split(obj.rawValue, sep)
	if value[0] == "" {
		return []string{}, obj.err
	}
	return value, nil
}

// DefaultStrings 转为[]string类型，sep为分隔符，如果出错则只返回默认值
func (obj *Object) DefaultStrings(sep string, defaultValue []string) []string {
	if obj.err != nil {
		return defaultValue
	}
	value := strings.Split(obj.rawValue, sep)
	return value
}

// Int 转为int类型
func (obj *Object) Int() (int, error) {
	if obj.err != nil || obj.rawValue == "" {
		return 0, obj.err
	}
	value, err := strconv.Atoi(obj.rawValue)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// DefaultInt 转为int类型，如果出错则只返回默认值
func (obj *Object) DefaultInt(defaultValue int) int {
	if obj.err != nil {
		return defaultValue
	}
	value, err := strconv.Atoi(obj.rawValue)
	if err != nil {
		return defaultValue
	}
	return value
}

// Int8 转为int8类型
func (obj *Object) Int8() (int8, error) {
	if obj.err != nil || obj.rawValue == "" {
		return 0, obj.err
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(value), nil
}

// DefaultInt8 转为int8类型，如果出错则只返回默认值
func (obj *Object) DefaultInt8(defaultValue int8) int8 {
	if obj.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 8)
	if err != nil {
		return defaultValue
	}
	return int8(value)
}

// Int16 转为int16类型
func (obj *Object) Int16() (int16, error) {
	if obj.err != nil || obj.rawValue == "" {
		return 0, obj.err
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(value), nil
}

// DefaultInt16 转为int16类型，如果出错则只返回默认值
func (obj *Object) DefaultInt16(defaultValue int16) int16 {
	if obj.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 16)
	if err != nil {
		return defaultValue
	}
	return int16(value)
}

// Int32 转为int32类型
func (obj *Object) Int32() (int32, error) {
	if obj.err != nil || obj.rawValue == "" {
		return 0, obj.err
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

// DefaultInt32 转为int32类型，如果出错则只返回默认值
func (obj *Object) DefaultInt32(defaultValue int32) int32 {
	if obj.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int32(value)
}

// Int64 转为int64类型
func (obj *Object) Int64() (int64, error) {
	if obj.err != nil || obj.rawValue == "" {
		return 0, obj.err
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// DefaultInt64 转为int64类型，如果出错则只返回默认值
func (obj *Object) DefaultInt64(defaultValue int64) int64 {
	if obj.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

// Float32 转为float32类型
func (obj *Object) Float32() (float32, error) {
	if obj.err != nil || obj.rawValue == "" {
		return 0, obj.err
	}
	value, err := strconv.ParseFloat(obj.rawValue, 32)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}

// DefaultFloat32 转为float32类型，如果出错则只返回默认值
func (obj *Object) DefaultFloat32(defaultValue float32) float32 {
	if obj.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseFloat(obj.rawValue, 32)
	if err != nil {
		return defaultValue
	}
	return float32(value)
}

// Float64 转为float64类型
func (obj *Object) Float64() (float64, error) {
	if obj.err != nil || obj.rawValue == "" {
		return 0, obj.err
	}
	value, err := strconv.ParseFloat(obj.rawValue, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// DefaultFloat64 转为float64类型，如果出错则只返回默认值
func (obj *Object) DefaultFloat64(defaultValue float64) float64 {
	if obj.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseFloat(obj.rawValue, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

// Bool 转为bool类型
func (obj *Object) Bool() (bool, error) {
	if obj.err != nil || obj.rawValue == "" {
		return false, obj.err
	}
	value, err := strconv.ParseBool(obj.rawValue)
	if err != nil {
		return false, err
	}
	return value, nil
}

// DefaultBool 转为bool类型，如果出错则只返回默认值
func (obj *Object) DefaultBool(defaultValue bool) bool {
	if obj.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseBool(obj.rawValue)
	if err != nil {
		return defaultValue
	}
	return value
}
