package filter

import (
	"slices"
	"strconv"
	"unicode"
)

// DenyStr 阻止denyValues中的值
func (self *Str) DenyStr(denyValues []string, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for k := range denyValues {
		if denyValues[k] == self.currentValue {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// DenyOtherSymbol 阻止allowValues之外的符号
func (self *Str) DenyOtherSymbol(allowValues []rune, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, r := range self.currentValue {
		// 如果是标点符号
		if unicode.IsPunct(r) {
			if !slices.Contains(allowValues, r) {
				self.err = wrapError(self.name, InvalidErrorText, customError...)
				return self
			}
		}
	}
	return self
}

// DenyInt 阻止denyValues中的值
func (self *Str) DenyInt(denyValues []int, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.Atoi(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for k := range denyValues {
		if value == denyValues[k] {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// DenyInt8 阻止denyValues中的值
func (self *Str) DenyInt8(denyValues []int8, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 8)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	value := int8(value64)
	for k := range denyValues {
		if value == denyValues[k] {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// DenyInt16 阻止denyValues中的值
func (self *Str) DenyInt16(denyValues []int16, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 16)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	value := int16(value64)
	for k := range denyValues {
		if value == denyValues[k] {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// DenyInt32 阻止denyValues中的值
func (self *Str) DenyInt32(denyValues []int32, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 32)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	value := int32(value64)
	for k := range denyValues {
		if value == denyValues[k] {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// DenyInt64 阻止denyValues中的值
func (self *Str) DenyInt64(denyValues []int64, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 64)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for k := range denyValues {
		if value64 == denyValues[k] {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// DenyFloat32 阻止denyValues中的值
func (self *Str) DenyFloat32(denyValues []float32, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := strconv.ParseFloat(self.currentValue, 32)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	value32 := float32(value)
	for k := range denyValues {
		if value32 == denyValues[k] {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// DenyFloats64 阻止denyValues中的值
func (self *Str) DenyFloat64(denyValues []float64, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseFloat(self.currentValue, 64)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for k := range denyValues {
		if value64 == denyValues[k] {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}
