package filter

import (
	"strconv"
	"strings"
)

// String 转为string类型
func (obj *Object) String() (string, error) {
	if obj.err != nil {
		return "", obj.err
	}
	return obj.rawValue, nil
}

// MustString 转为string类型，如果出错则返回默认值
func (obj *Object) MustString(def string) string {
	if obj.err != nil {
		return def
	}
	return obj.rawValue
}

// Strings 转为[]string类型，sep为分隔符
func (obj *Object) Strings(sep string) []string {
	if obj.err != nil {
		return nil
	}
	value := strings.Split(obj.rawValue, sep)
	return value
}

// Int 转为int类型
func (obj *Object) Int() (int, error) {
	if obj.err != nil {
		return 0, obj.err
	}
	value, err := strconv.Atoi(obj.rawValue)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// MustInt 转为int类型，如果出错则返回默认值
func (obj *Object) MustInt(def int) int {
	if obj.err != nil {
		return def
	}
	value, err := strconv.Atoi(obj.rawValue)
	if err != nil {
		return def
	}
	return value
}

// Int8 转为int8类型
func (obj *Object) Int8() (int8, error) {
	if obj.err != nil {
		return 0, obj.err
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(value), nil
}

// MustInt8 转为int8类型，如果出错则返回默认值
func (obj *Object) MustInt8(def int8) int8 {
	if obj.err != nil {
		return def
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 8)
	if err != nil {
		return def
	}
	return int8(value)
}

// Int16 转为int16类型
func (obj *Object) Int16() (int16, error) {
	if obj.err != nil {
		return 0, obj.err
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(value), nil
}

// Int16 转为int16类型，如果出错则返回默认值
func (obj *Object) MustInt16(def int16) int16 {
	if obj.err != nil {
		return def
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 16)
	if err != nil {
		return def
	}
	return int16(value)
}

// Int32 转为int32类型
func (obj *Object) Int32() (int32, error) {
	if obj.err != nil {
		return 0, obj.err
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

// Int32 转为int32类型，如果出错则返回默认值
func (obj *Object) MustInt32(def int32) int32 {
	if obj.err != nil {
		return def
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}

// Int64 转为int64类型
func (obj *Object) Int64() (int64, error) {
	if obj.err != nil {
		return 0, obj.err
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// Int64 转为int64类型，如果出错则返回默认值
func (obj *Object) MustInt64(def int64) int64 {
	if obj.err != nil {
		return def
	}
	value, err := strconv.ParseInt(obj.rawValue, 10, 64)
	if err != nil {
		return def
	}
	return value
}

// Float32 转为float32类型
func (obj *Object) Float32() (float32, error) {
	if obj.err != nil {
		return 0, obj.err
	}
	value, err := strconv.ParseFloat(obj.rawValue, 32)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}

// Float32 转为float32类型，如果出错则返回默认值
func (obj *Object) MustFloat32(def float32) float32 {
	if obj.err != nil {
		return def
	}
	value, err := strconv.ParseFloat(obj.rawValue, 32)
	if err != nil {
		return def
	}
	return float32(value)
}

// Float64 转为float64类型
func (obj *Object) Float64() (float64, error) {
	if obj.err != nil {
		return 0, obj.err
	}
	value, err := strconv.ParseFloat(obj.rawValue, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// Float64 转为float64类型，如果出错则返回默认值
func (obj *Object) MustFloat64(def float64) float64 {
	if obj.err != nil {
		return def
	}
	value, err := strconv.ParseFloat(obj.rawValue, 64)
	if err != nil {
		return def
	}
	return value
}

// Bool 转为bool类型
func (obj *Object) Bool() (bool, error) {
	if obj.err != nil {
		return false, obj.err
	}
	value, err := strconv.ParseBool(obj.rawValue)
	if err != nil {
		return false, err
	}
	return value, nil
}

// Bool 转为bool类型，如果出错则返回默认值
func (obj *Object) MustBool(def bool) bool {
	if obj.err != nil {
		return def
	}
	value, err := strconv.ParseBool(obj.rawValue)
	if err != nil {
		return def
	}
	return value
}
