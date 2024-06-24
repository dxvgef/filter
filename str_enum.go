package filter

import (
	"strconv"
)

// EnumStr 仅允许allows中的值
func (self *Str) EnumStr(allows []string, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for k := range allows {
		if allows[k] == self.currentValue {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt 仅允许allows中的值
func (self *Str) EnumInt(allows []int, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.Atoi(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for k := range allows {
		if value == allows[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt8 仅允许allows中的值
func (self *Str) EnumInt8(allows []int8, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 8)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	value := int8(value64)
	for k := range allows {
		if value == allows[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt16 仅允许allows中的值
func (self *Str) EnumInt16(allows []int16, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 16)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	value := int16(value64)
	for k := range allows {
		if value == allows[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt32 仅允许allows中的值
func (self *Str) EnumInt32(allows []int32, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 32)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	value32 := int32(value64)
	for k := range allows {
		if value32 == allows[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumInt64 仅允许allows中的值
func (self *Str) EnumInt64(allows []int64, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 64)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for k := range allows {
		if value64 == allows[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumFloat32 仅允许allows中的值
func (self *Str) EnumFloat32(allows []float32, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.ParseFloat(self.currentValue, 32)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	value32 := float32(value)
	for k := range allows {
		if value32 == allows[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumFloat64 仅允许allows中的值
func (self *Str) EnumFloat64(allows []float64, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseFloat(self.currentValue, 64)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for k := range allows {
		if value64 == allows[k] {
			return self
		}
	}
	self.err = wrapError(self.name, InvalidErrorText, customError...)
	return self
}

// EnumStrSlice 仅允许使用allows中的值
func (self *Str) EnumStrSlice(sep string, allows []string, trimSpace bool, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	values, err := self.SliceString(sep, trimSpace)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for k := range values {
		if !inStrings(values[k], allows) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// EnumIntSlice 仅允许使用allows中的值
func (self *Str) EnumIntSlice(sep string, allows []int, trimSpace bool, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	values, err := self.SliceString(sep, trimSpace)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for k := range values {
		v, err := strconv.Atoi(values[k])
		if err != nil {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
		if !inInt(v, allows) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
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
