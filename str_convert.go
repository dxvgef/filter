package filter

import (
	"strconv"
	"strings"
)

// FromStr 转为string类型
func (self *Str) String() (string, error) {
	if self.err != nil || self.currentValue == "" {
		return "", self.err
	}
	return self.currentValue, nil
}

// DefaultString 转为string类型，如果出错则只返回默认值
func (self *Str) DefaultString(def string) string {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	return self.currentValue
}

// SliceString 使用SetSeparator()设置的分隔符拆分成[]string类型
func (self *Str) SliceString(sep string) ([]string, error) {
	if self.err != nil || self.currentValue == "" {
		return []string{}, self.err
	}
	value := strings.Split(self.currentValue, sep)
	if self.require {
		v := false
		for k := range value {
			if value[k] != "" {
				v = true
				break
			}
		}
		if !v {
			self.err = wrapError(self.name, self.requireErr)
			return []string{}, self.err
		}
	}
	return value, nil
}

// DefaultSliceString 使用SetSeparator()设置的分隔符拆分成[]string类型，如果出错则只返回默认值
func (self *Str) DefaultSliceString(sep string, def []string) []string {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceString(sep)
	if err != nil {
		return def
	}
	return value
}

// Int 转为int类型
func (self *Str) Int() (int, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	return strconv.Atoi(self.currentValue)
}

// DefaultInt 转为int类型，如果出错则只返回默认值
func (self *Str) DefaultInt(def int) int {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.Int()
	if err != nil {
		return def
	}
	return value
}

// SliceInt 转为[]int类型
func (self *Str) SliceInt(sep string) (result []int, err error) {
	if self.err != nil || self.currentValue == "" {
		return nil, self.err
	}
	var (
		values []string
		value  int
	)
	values, err = self.SliceString(sep)
	if err != nil {
		return
	}
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
func (self *Str) DefaultSliceInt(sep string, def []int) []int {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceInt(sep)
	if err != nil {
		return def
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
func (self *Str) DefaultUint(def uint) uint {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.Uint()
	if err != nil {
		return def
	}
	return value
}

// SliceUint 转为[]int类型
func (self *Str) SliceUint(sep string) (result []uint, err error) {
	if self.err != nil || self.currentValue == "" {
		return nil, self.err
	}
	var (
		values []string
		value  uint64
	)
	values, err = self.SliceString(sep)
	if err != nil {
		return
	}
	for k := range values {
		value, err = strconv.ParseUint(values[k], 10, 0)
		if err != nil {
			return
		}
		result = append(result, uint(value))
	}
	return
}

// DefaultSliceUint 转为[]int类型，出错则返回传入的默认值
func (self *Str) DefaultSliceUint(sep string, def []uint) []uint {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceUint(sep)
	if err != nil {
		return def
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
func (self *Str) DefaultInt8(def int8) int8 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 8)
	if err != nil {
		return def
	}
	return int8(value)
}

// SliceInt8 转为[]int8类型
func (self *Str) SliceInt8(sep string) (result []int8, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  int64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceInt8(sep string, def []int8) []int8 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceInt8(sep)
	if err != nil {
		return def
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
func (self *Str) DefaultUint8(def uint8) uint8 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 8)
	if err != nil {
		return def
	}
	return uint8(value)
}

// SliceUint8 转为[]int8类型
func (self *Str) SliceUint8(sep string) (result []uint8, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  uint64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceUint8(sep string, def []uint8) []uint8 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceUint8(sep)
	if err != nil {
		return def
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
func (self *Str) DefaultInt16(def int16) int16 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 16)
	if err != nil {
		return def
	}
	return int16(value)
}

// SliceInt16 转为[]int16类型
func (self *Str) SliceInt16(sep string) (result []int16, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  int64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceInt16(sep string, def []int16) []int16 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceInt16(sep)
	if err != nil {
		return def
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
func (self *Str) DefaultUint16(def uint16) uint16 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 16)
	if err != nil {
		return def
	}
	return uint16(value)
}

// SliceUint16 转为[]uint16类型
func (self *Str) SliceUint16(sep string) (result []uint16, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  uint64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceUint16(sep string, def []uint16) []uint16 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceUint16(sep)
	if err != nil {
		return def
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
func (self *Str) DefaultInt32(def int32) int32 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}

// SliceInt32 转为[]int32类型
func (self *Str) SliceInt32(sep string) (result []int32, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  int64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceInt32(sep string, def []int32) []int32 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceInt32(sep)
	if err != nil {
		return def
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
func (self *Str) DefaultUint32(def uint32) uint32 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 32)
	if err != nil {
		return def
	}
	return uint32(value)
}

// SliceUint32 转为[]uint32类型
func (self *Str) SliceUint32(sep string) (result []uint32, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  uint64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceUint32(sep string, def []uint32) []uint32 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceUint32(sep)
	if err != nil {
		return def
	}
	return value
}

// Int64 转为int64类型
func (self *Str) Int64() (int64, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	return strconv.ParseInt(self.currentValue, 10, 64)
}

// DefaultInt64 转为int64类型，如果出错则只返回默认值
func (self *Str) DefaultInt64(def int64) int64 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseInt(self.currentValue, 10, 64)
	if err != nil {
		return def
	}
	return value
}

// SliceInt64 转为[]int64类型
func (self *Str) SliceInt64(sep string) (result []int64, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  int64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceInt64(sep string, def []int64) []int64 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceInt64(sep)
	if err != nil {
		return def
	}
	return value
}

// Uint64 转为uint64类型
func (self *Str) Uint64() (uint64, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	return strconv.ParseUint(self.currentValue, 10, 64)
}

// DefaultUint64 转为uint64类型，如果出错则只返回默认值
func (self *Str) DefaultUint64(def uint64) uint64 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseUint(self.currentValue, 10, 64)
	if err != nil {
		return def
	}
	return value
}

// SliceUint64 转为[]uint64类型
func (self *Str) SliceUint64(sep string) (result []uint64, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  uint64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceUint64(sep string, def []uint64) []uint64 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceUint64(sep)
	if err != nil {
		return def
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
func (self *Str) DefaultFloat32(def float32) float32 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseFloat(self.currentValue, 32)
	if err != nil {
		return def
	}
	return float32(value)
}

// SliceFloat32 转为[]float32类型
func (self *Str) SliceFloat32(sep string) (result []float32, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  float64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceFloat32(sep string, def []float32) []float32 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceFloat32(sep)
	if err != nil {
		return def
	}
	return value
}

// Float64 转为float64类型
func (self *Str) Float64() (float64, error) {
	if self.err != nil || self.currentValue == "" {
		return 0, self.err
	}
	return strconv.ParseFloat(self.currentValue, 64)
}

// DefaultFloat64 转为float64类型，如果出错则只返回默认值
func (self *Str) DefaultFloat64(def float64) float64 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseFloat(self.currentValue, 64)
	if err != nil {
		return def
	}
	return value
}

// SliceFloat64 转为[]float64类型
func (self *Str) SliceFloat64(sep string) (result []float64, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  float64
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceFloat64(sep string, def []float64) []float64 {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceFloat64(sep)
	if err != nil {
		return def
	}
	return value
}

// Bool 转为bool类型
func (self *Str) Bool() (bool, error) {
	if self.err != nil || self.currentValue == "" {
		return false, self.err
	}
	return strconv.ParseBool(self.currentValue)
}

// DefaultBool 转为bool类型，如果出错则只返回默认值
func (self *Str) DefaultBool(def bool) bool {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := strconv.ParseBool(self.currentValue)
	if err != nil {
		return def
	}
	return value
}

// SliceBool 转为[]bool类型
func (self *Str) SliceBool(sep string) (result []bool, err error) {
	if self.err != nil || self.currentValue == "" {
		err = self.err
		return
	}
	var (
		values []string
		value  bool
	)
	values, err = self.SliceString(sep)
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
func (self *Str) DefaultSliceBool(sep string, def []bool) []bool {
	if self.err != nil || self.currentValue == "" {
		return def
	}
	value, err := self.SliceBool(sep)
	if err != nil {
		return def
	}
	return value
}
