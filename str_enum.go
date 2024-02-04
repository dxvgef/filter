package filter

import (
	"strconv"
)

// EnumStr 仅允许[]string中的值
func (self *Str) EnumStr(slice []string, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for k := range slice {
		if slice[k] == self.currentValue {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt 仅允许[]int中的值
func (self *Str) EnumInt(i []int, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.Atoi(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, ConvErrorText, customError...)
		return self
	}
	for k := range i {
		if value == i[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt8 仅允许[]int8中的值
func (self *Str) EnumInt8(i []int8, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 8)
	if err != nil {
		self.err = wrapError(self.name, ConvErrorText, customError...)
		return self
	}
	value := int8(value64)
	for k := range i {
		if value == i[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt16 仅允许[]int16中的值
func (self *Str) EnumInt16(i []int16, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 16)
	if err != nil {
		self.err = wrapError(self.name, ConvErrorText, customError...)
		return self
	}
	value := int16(value64)
	for k := range i {
		if value == i[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt32 仅允许[]int32中的值
func (self *Str) EnumInt32(i []int32, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 32)
	if err != nil {
		self.err = wrapError(self.name, ConvErrorText, customError...)
		return self
	}
	value32 := int32(value64)
	for k := range i {
		if value32 == i[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt64 仅允许[]int64中的值
func (self *Str) EnumInt64(i []int64, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 64)
	if err != nil {
		self.err = wrapError(self.name, ConvErrorText, customError...)
		return self
	}
	for k := range i {
		if value64 == i[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumFloat32 仅允许[]float32中的值
func (self *Str) EnumFloat32(f []float32, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.ParseFloat(self.currentValue, 32)
	if err != nil {
		self.err = wrapError(self.name, ConvErrorText, customError...)
		return self
	}
	value32 := float32(value)
	for k := range f {
		if value32 == f[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumFloat64 仅允许[]float64中的值
func (self *Str) EnumFloat64(f []float64, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseFloat(self.currentValue, 64)
	if err != nil {
		self.err = wrapError(self.name, ConvErrorText, customError...)
		return self
	}
	for k := range f {
		if value64 == f[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumStrSlice 检查[]string中的元素，仅允许指定[]string中的值
func (self *Str) EnumStrSlice(sep string, slice []string, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	values, err := self.SliceString(sep)
	if err != nil {
		self.err = wrapError(self.name, ConvErrorText, customError...)
		return self
	}
	for k := range values {
		if !inStrings(values[k], slice) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// EnumIntSlice 检查[]string中的元素，仅允许指潘秉衡[]int中的值
func (self *Str) EnumIntSlice(sep string, slice []int, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	values, err := self.SliceString(sep)
	if err != nil {
		self.err = wrapError(self.name, ConvErrorText, customError...)
		return self
	}
	for k := range values {
		v, err := strconv.Atoi(values[k])
		if err != nil {
			self.err = wrapError(self.name, ConvErrorText, customError...)
			return self
		}
		if !inInt(v, slice) {
			self.err = wrapError(self.name, ConvErrorText, customError...)
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

func inStrings(v string, s []string) bool {
	for k := range s {
		if s[k] == v {
			return true
		}
	}
	return false
}

func inInt(v int, s []int) bool {
	for k := range s {
		if s[k] == v {
			return true
		}
	}
	return false
}
