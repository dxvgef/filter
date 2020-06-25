package filter

import (
	"strconv"
	"strings"
)

// String 转为string类型
func (self *Str) String() (string, error) {
	if self.err != nil || self.currentValue == "" {
		return "", self.err
	}
	return self.currentValue, nil
}

// DefaultString 转为string类型，如果出错则只返回默认值
func (self *Str) DefaultString(defaultValue string) string {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	return self.currentValue
}

// SliceString 使用SetSeparator()设置的分隔符拆分成[]string类型
func (self *Str) SliceString() ([]string, error) {
	if self.err != nil || self.currentValue == "" {
		return []string{}, self.err
	}
	value := strings.Split(self.currentValue, self.sep)
	if value[0] == "" {
		return []string{}, self.err
	}
	return value, nil
}

// DefaultSliceString 使用SetSeparator()设置的分隔符拆分成[]string类型，如果出错则只返回默认值
func (self *Str) DefaultSliceString(defaultValue []string) []string {
	if self.err != nil {
		return defaultValue
	}
	value, err := self.SliceString()
	if err != nil {
		return defaultValue
	}
	return value
}

// Int 转为int类型
func (self *Str) Int() (int, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.Atoi(self.currentValue)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// DefaultInt 转为int类型，如果出错则只返回默认值
func (self *Str) DefaultInt(defaultValue int) int {
	if self.err != nil {
		return defaultValue
	}
	value, err := self.Int()
	if err != nil {
		return defaultValue
	}
	return value
}

// IntSlice 转为[]int类型
func (self *Str) IntSlice() (result []int, err error) {
	if self.err != nil || self.currentValue == "" {
		return
	}
	values := strings.Split(self.currentValue, self.sep)
	var value int
	for k := range values {
		value, err = strconv.Atoi(values[k])
		if err != nil {
			return
		}
		result = append(result, value)
	}
	return
}

// DefaultIntSlice 转为[]int类型，出错则返回传入的默认值
func (self *Str) DefaultIntSlice(defaultValue []int) []int {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.IntSlice()
	if err != nil {
		return defaultValue
	}
	return value
}

// Int8 转为int8类型
func (self *Str) Int8() (int8, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(value), nil
}

// DefaultInt8 转为int8类型，如果出错则只返回默认值
func (self *Str) DefaultInt8(defaultValue int8) int8 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 8)
	if err != nil {
		return defaultValue
	}
	return int8(value)
}

// Int8 转为[]int8类型
func (self *Str) Int8Slice() (result []int8, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  int64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseInt(values[k], 10, 8)
		if err != nil {
			return
		}
		result = append(result, int8(value))
	}
	return
}

// DefaultInt8Slice 转为[]int8类型，如果出错则只返回默认值
func (self *Str) DefaultInt8Slice(defaultValue []int8) []int8 {
	if self.err != nil {
		return defaultValue
	}
	value, err := self.Int8Slice()
	if err != nil {
		return defaultValue
	}
	return value
}

// Int16 转为int16类型
func (self *Str) Int16() (int16, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(value), nil
}

// DefaultInt16 转为int16类型，如果出错则只返回默认值
func (self *Str) DefaultInt16(defaultValue int16) int16 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 16)
	if err != nil {
		return defaultValue
	}
	return int16(value)
}

// Int16 转为[]int16类型
func (self *Str) Int16Slice() (result []int16, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  int64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseInt(values[k], 10, 16)
		if err != nil {
			return
		}
		result = append(result, int16(value))
	}
	return
}

// DefaultInt16Slice 转为[]int16类型
func (self *Str) DefaultInt16Slice(defaultValue []int16) []int16 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.Int16Slice()
	if err != nil {
		return defaultValue
	}
	return value
}

// Int32 转为int32类型
func (self *Str) Int32() (int32, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

// DefaultInt32 转为int32类型，如果出错则只返回默认值
func (self *Str) DefaultInt32(defaultValue int32) int32 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int32(value)
}

// Int32Slice 转为[]int32类型
func (self *Str) Int32Slice() (result []int32, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  int64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseInt(values[k], 10, 32)
		if err != nil {
			return
		}
		result = append(result, int32(value))
	}
	return
}

// DefaultInt32 转为[]int8类型
func (self *Str) DefaultInt32Slice(defaultValue []int32) []int32 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.Int32Slice()
	if err != nil {
		return defaultValue
	}
	return value
}

// Int64 转为int64类型
func (self *Str) Int64() (int64, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// DefaultInt64 转为int64类型，如果出错则只返回默认值
func (self *Str) DefaultInt64(defaultValue int64) int64 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

// Int64 转为[]int64x类型
func (self *Str) Int64Slice() (result []int64, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  int64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseInt(values[k], 10, 64)
		if err != nil {
			return
		}
		result = append(result, value)
	}
	return
}

// DefaultInt64Slice 转为[]int64类型，如果出错则只返回默认值
func (self *Str) DefaultInt64Slice(defaultValue []int64) []int64 {
	if self.err != nil {
		return defaultValue
	}
	value, err := self.Int64Slice()
	if err != nil {
		return defaultValue
	}
	return value
}

// Float32 转为float32类型
func (self *Str) Float32() (float32, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseFloat(self.currentValue, 32)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}

// DefaultFloat32 转为float32类型，如果出错则只返回默认值
func (self *Str) DefaultFloat32(defaultValue float32) float32 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseFloat(self.currentValue, 32)
	if err != nil {
		return defaultValue
	}
	return float32(value)
}

// Float32Slice 转为[]float32类型
func (self *Str) Float32Slice() (result []float32, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  float64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseFloat(values[k], 32)
		if err != nil {
			return
		}
		result = append(result, float32(value))
	}
	return
}

// DefaultFloat32Slice 转为[]float32类型
func (self *Str) DefaultFloat32Slice(defaultValue []float32) []float32 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.Float32Slice()
	if err != nil {
		return defaultValue
	}
	return value
}

// Float64 转为float64类型
func (self *Str) Float64() (float64, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseFloat(self.currentValue, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// DefaultFloat64 转为float64类型，如果出错则只返回默认值
func (self *Str) DefaultFloat64(defaultValue float64) float64 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseFloat(self.currentValue, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

// Float64Slice 转为[]float64类型
func (self *Str) Float64Slice() (result []float64, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  float64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseFloat(values[k], 64)
		if err != nil {
			return
		}
		result = append(result, value)
	}
	return
}

// DefaultFloat64Slice 转为[]float64类型
func (self *Str) DefaultFloat64Slice(defaultValue []float64) []float64 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.Float64Slice()
	if err != nil {
		return defaultValue
	}
	return value
}

// Bool 转为bool类型
func (self *Str) Bool() (bool, error) {
	if self.err != nil || self.currentValue == "" {
		return false, self.err
	}
	value, err := strconv.ParseBool(self.currentValue)
	if err != nil {
		return false, err
	}
	return value, nil
}

// DefaultBool 转为bool类型，如果出错则只返回默认值
func (self *Str) DefaultBool(defaultValue bool) bool {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseBool(self.currentValue)
	if err != nil {
		return defaultValue
	}
	return value
}
