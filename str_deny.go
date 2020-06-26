package filter

import "strconv"

// DenyString 阻止存在于[]string中的值
func (self *Str) DenyString(slice []string, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for k := range slice {
		if slice[k] == self.currentValue {
			self.err = wrapError(self.name, "", customError...)
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
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	for k := range i {
		if value == i[k] {
			self.err = wrapError(self.name, "", customError...)
			return self
		}
	}
	return self
}

// DenyInt8 阻止存在于[]int8中的值
func (self *Str) DenyInt8(i []int8, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 8)
	if err != nil {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	value := int8(value64)
	for k := range i {
		if value == i[k] {
			self.err = wrapError(self.name, "", customError...)
			return self
		}
	}
	return self
}

// DenyInt16 阻止存在于[]int16中的值
func (self *Str) DenyInt16(i []int16, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 16)
	if err != nil {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	value := int16(value64)
	for k := range i {
		if value == i[k] {
			self.err = wrapError(self.name, "", customError...)
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
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	value := int32(value64)
	for k := range i {
		if value == i[k] {
			self.err = wrapError(self.name, "", customError...)
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
	value64, err := strconv.ParseInt(self.currentValue, 10, 64)
	if err != nil {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	for k := range i {
		if value64 == i[k] {
			self.err = wrapError(self.name, "", customError...)
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
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	value32 := float32(value)
	for k := range f {
		if value32 == f[k] {
			self.err = wrapError(self.name, "", customError...)
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
	value64, err := strconv.ParseFloat(self.currentValue, 64)
	if err != nil {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	for k := range f {
		if value64 == f[k] {
			self.err = wrapError(self.name, "", customError...)
			return self
		}
	}
	return self
}
