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

// SliceInt 转为[]int类型
func (self *Str) SliceInt() (result []int, err error) {
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

// DefaultSliceInt 转为[]int类型，出错则返回传入的默认值
func (self *Str) DefaultSliceInt(defaultValue []int) []int {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.SliceInt()
	if err != nil {
		return defaultValue
	}
	return value
}

// Uint 转为uint类型
func (self *Str) Uint() (uint, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

// DefaultUint 转为uint类型，如果出错则只返回默认值
func (self *Str) DefaultUint(defaultValue uint) uint {
	if self.err != nil {
		return defaultValue
	}
	value, err := self.Uint()
	if err != nil {
		return defaultValue
	}
	return value
}

// SliceInt 转为[]int类型
func (self *Str) SliceUint() (result []uint, err error) {
	if self.err != nil || self.currentValue == "" {
		return
	}
	values := strings.Split(self.currentValue, self.sep)
	var value uint64
	for k := range values {
		value, err = strconv.ParseUint(values[k], 10, 0)
		if err != nil {
			return
		}
		result = append(result, uint(value))
	}
	return
}

// DefaultSliceInt 转为[]int类型，出错则返回传入的默认值
func (self *Str) DefaultSliceUint(defaultValue []uint) []uint {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.SliceUint()
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

// SliceInt8 转为[]int8类型
func (self *Str) SliceInt8() (result []int8, err error) {
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

// DefaultSliceInt8 转为[]int8类型，如果出错则只返回默认值
func (self *Str) DefaultSliceInt8(defaultValue []int8) []int8 {
	if self.err != nil {
		return defaultValue
	}
	value, err := self.SliceInt8()
	if err != nil {
		return defaultValue
	}
	return value
}

// Uint8 转为uint8类型
func (self *Str) Uint8() (uint8, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(value), nil
}

// DefaultUint8 转为uint8类型，如果出错则只返回默认值
func (self *Str) DefaultUint8(defaultValue uint8) uint8 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 8)
	if err != nil {
		return defaultValue
	}
	return uint8(value)
}

// SliceUint8 转为[]int8类型
func (self *Str) SliceUint8() (result []uint8, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  uint64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseUint(values[k], 10, 8)
		if err != nil {
			return
		}
		result = append(result, uint8(value))
	}
	return
}

// DefaultSliceUint8 转为[]uint8类型，如果出错则只返回默认值
func (self *Str) DefaultSliceUint8(defaultValue []uint8) []uint8 {
	if self.err != nil {
		return defaultValue
	}
	value, err := self.SliceUint8()
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

// SliceInt16 转为[]int16类型
func (self *Str) SliceInt16() (result []int16, err error) {
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

// DefaultSliceInt16 转为[]int16类型
func (self *Str) DefaultSliceInt16(defaultValue []int16) []int16 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.SliceInt16()
	if err != nil {
		return defaultValue
	}
	return value
}

// Uint16 转为uint16类型
func (self *Str) Uint16() (uint16, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(value), nil
}

// DefaultUint16 转为uint16类型，如果出错则只返回默认值
func (self *Str) DefaultUint16(defaultValue uint16) uint16 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 16)
	if err != nil {
		return defaultValue
	}
	return uint16(value)
}

// SliceUint16 转为[]uint16类型
func (self *Str) SliceUint16() (result []uint16, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  uint64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseUint(values[k], 10, 16)
		if err != nil {
			return
		}
		result = append(result, uint16(value))
	}
	return
}

// DefaultSliceUint16 转为[]uint16类型
func (self *Str) DefaultSliceUint16(defaultValue []uint16) []uint16 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.SliceUint16()
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

// SliceInt32 转为[]int32类型
func (self *Str) SliceInt32() (result []int32, err error) {
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

// DefaultSliceInt32 转为[]int8类型
func (self *Str) DefaultSliceInt32(defaultValue []int32) []int32 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.SliceInt32()
	if err != nil {
		return defaultValue
	}
	return value
}

// Uint32 转为uint32类型
func (self *Str) Uint32() (uint32, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(value), nil
}

// DefaultUint32 转为uint32类型，如果出错则只返回默认值
func (self *Str) DefaultUint32(defaultValue uint32) uint32 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 32)
	if err != nil {
		return defaultValue
	}
	return uint32(value)
}

// SliceUint32 转为[]uint32类型
func (self *Str) SliceUint32() (result []uint32, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  uint64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseUint(values[k], 10, 32)
		if err != nil {
			return
		}
		result = append(result, uint32(value))
	}
	return
}

// DefaultSliceUint32 转为[]uint8类型
func (self *Str) DefaultSliceUint32(defaultValue []uint32) []uint32 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.SliceUint32()
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

// SliceInt64 转为[]int64类型
func (self *Str) SliceInt64() (result []int64, err error) {
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

// DefaultSliceInt64 转为[]int64类型，如果出错则只返回默认值
func (self *Str) DefaultSliceInt64(defaultValue []int64) []int64 {
	if self.err != nil {
		return defaultValue
	}
	value, err := self.SliceInt64()
	if err != nil {
		return defaultValue
	}
	return value
}

// Uint64 转为uint64类型
func (self *Str) Uint64() (uint64, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// DefaultUint64 转为uint64类型，如果出错则只返回默认值
func (self *Str) DefaultUint64(defaultValue uint64) uint64 {
	if self.err != nil {
		return defaultValue
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

// SliceUint64 转为[]uint64类型
func (self *Str) SliceUint64() (result []uint64, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  uint64
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseUint(values[k], 10, 64)
		if err != nil {
			return
		}
		result = append(result, value)
	}
	return
}

// DefaultSliceUint64 转为[]uint64类型，如果出错则只返回默认值
func (self *Str) DefaultSliceUint64(defaultValue []uint64) []uint64 {
	if self.err != nil {
		return defaultValue
	}
	value, err := self.SliceUint64()
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

// SliceFloat32 转为[]float32类型
func (self *Str) SliceFloat32() (result []float32, err error) {
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

// DefaultSliceFloat32 转为[]float32类型
func (self *Str) DefaultSliceFloat32(defaultValue []float32) []float32 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.SliceFloat32()
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

// SliceFloat64 转为[]float64类型
func (self *Str) SliceFloat64() (result []float64, err error) {
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

// DefaultSliceFloat64 转为[]float64类型
func (self *Str) DefaultSliceFloat64(defaultValue []float64) []float64 {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.SliceFloat64()
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

// SliceBool 转为[]bool类型
func (self *Str) SliceBool() (result []bool, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  bool
	)
	values, err = self.SliceString()
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseBool(values[k])
		if err != nil {
			return
		}
		result = append(result, value)
	}
	return
}

// DefaultSliceBool 转为[]bool类型
func (self *Str) DefaultSliceBool(defaultValue []bool) []bool {
	if self.err != nil || self.currentValue == "" {
		return defaultValue
	}
	value, err := self.SliceBool()
	if err != nil {
		return defaultValue
	}
	return value
}
