package filter

import (
	"strconv"
)

// SetSeparator 设置过滤器缓存中的分隔符，之后都以此分隔符来拆解value成slice
func (self *Str) SetSeparator(sep string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.sep = sep
	return self
}

// EnumString 仅允许[]string中的值
func (self *Str) EnumString(slice []string, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for k := range slice {
		if slice[k] == self.currentValue {
			return self
		}
	}
	self.err = self.wrapError("", customError...)
	return self
}

// EnumInt 仅允许[]int中的值
func (self *Str) EnumInt(i []int, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.Atoi(self.currentValue)
	if err != nil {
		self.err = self.wrapError("", customError...)
		return self
	}
	for k := range i {
		if value == i[k] {
			return self
		}
	}
	self.err = self.wrapError("", customError...)
	return self
}

// EnumInt32 仅允许[]int32中的值
func (self *Str) EnumInt32(i []int32, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 32)
	if err != nil {
		self.err = self.wrapError("", customError...)
		return self
	}
	value32 := int32(value64)
	for k := range i {
		if value32 == i[k] {
			return self
		}
	}
	self.err = self.wrapError("", customError...)
	return self
}

// EnumInt64 仅允许[]int64中的值
func (self *Str) EnumInt64(i []int64, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if self.int64 == 0 {
		var err error
		self.int64, err = strconv.ParseInt(self.currentValue, 10, 64)
		if err != nil {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	for k := range i {
		if self.int64 == i[k] {
			return self
		}
	}
	self.err = self.wrapError("", customError...)
	return self
}

// EnumFloat32 仅允许[]float32中的值
func (self *Str) EnumFloat32(f []float32, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.ParseFloat(self.currentValue, 32)
	if err != nil {
		self.err = self.wrapError("", customError...)
		return self
	}
	value32 := float32(value)
	for k := range f {
		if value32 == f[k] {
			return self
		}
	}
	self.err = self.wrapError("", customError...)
	return self
}

// EnumFloat64 仅允许[]float64中的值
func (self *Str) EnumFloat64(f []float64, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if self.float64 == 0 {
		var err error
		self.float64, err = strconv.ParseFloat(self.currentValue, 64)
		if err != nil {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	for k := range f {
		if self.float64 == f[k] {
			return self
		}
	}
	self.err = self.wrapError("", customError...)
	return self
}

// ------------------------------------ Deny ---------------------------------------

// DenyString 阻止存在于[]string中的值
func (self *Str) DenyString(slice []string, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for k := range slice {
		if slice[k] == self.currentValue {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	return self
}

// DenyInt 阻止存在于[]int中的值
func (self *Str) DenyInt(i []int, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.Atoi(self.currentValue)
	if err != nil {
		self.err = self.wrapError("", customError...)
		return self
	}
	for k := range i {
		if value == i[k] {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	return self
}

// DenyInt32 阻止存在于[]int32中的值
func (self *Str) DenyInt32(i []int32, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 32)
	if err != nil {
		self.err = self.wrapError("", customError...)
		return self
	}
	value32 := int32(value64)
	for k := range i {
		if value32 == i[k] {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	return self
}

// DenyInt64 阻止存在于[]int64中的值
func (self *Str) DenyInt64(i []int64, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if self.int64 == 0 {
		var err error
		self.int64, err = strconv.ParseInt(self.currentValue, 10, 64)
		if err != nil {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	for k := range i {
		if self.int64 == i[k] {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	return self
}

// DenyFloat32 阻止存在于[]float32中的值
func (self *Str) DenyFloat32(f []float32, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.ParseFloat(self.currentValue, 32)
	if err != nil {
		self.err = self.wrapError("", customError...)
		return self
	}
	value32 := float32(value)
	for k := range f {
		if value32 == f[k] {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	return self
}

// DenyFloats64 阻止存在于[]float64中的值
func (self *Str) DenyInFloat64(f []float64, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if self.float64 == 0 {
		var err error
		self.float64, err = strconv.ParseFloat(self.currentValue, 64)
		if err != nil {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	for k := range f {
		if self.float64 == f[k] {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	return self
}
