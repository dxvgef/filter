package filter

import "strconv"

// UpdateInt64 使用最新的value转换成int64类型的值并在内部缓存，用于之后的数值计算
func (self *Str) UpdateInt64(customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	i64, err := strconv.ParseInt(self.value, 10, 64)
	if err != nil {
		self.err = self.wrapError("", customError...)
		return self
	}
	self.int64 = i64
	return self
}

// UpdateFloat64 使用最新的value转换成float64类型的值并在内部缓存，用于之后的数值计算
func (self *Str) UpdateFloat64(customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	f64, err := strconv.ParseFloat(self.value, 64)
	if err != nil {
		self.err = self.wrapError("", customError...)
		return self
	}
	self.float64 = f64
	return self
}

// MinInteger 最小整数值
func (self *Str) MinInteger(min int64, customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	if self.int64 == 0 {
		var err error
		self.int64, err = strconv.ParseInt(self.value, 10, 64)
		if err != nil {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	if self.int64 < min {
		self.err = self.wrapError("", customError...)
		return self
	}
	return self
}

// MinInteger 最大整数值
func (self *Str) MaxInteger(max int64, customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	if self.int64 == 0 {
		var err error
		self.int64, err = strconv.ParseInt(self.value, 10, 64)
		if err != nil {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	if self.int64 > max {
		self.err = self.wrapError("", customError...)
		return self
	}
	return self
}

// MinFloat 最小浮点值
func (self *Str) MinFloat(min float64, customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	if self.float64 == 0 {
		var err error
		self.float64, err = strconv.ParseFloat(self.value, 64)
		if err != nil {
			self.err = self.wrapError("", customError...)
			return self
		}
	}

	if self.float64 < min {
		self.err = self.wrapError("", customError...)
		return self
	}
	return self
}

// MinFloat 最大浮点值
func (self *Str) MaxFloat(max float64, customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	if self.float64 == 0 {
		var err error
		self.float64, err = strconv.ParseFloat(self.value, 64)
		if err != nil {
			self.err = self.wrapError("", customError...)
			return self
		}
	}
	if self.float64 > max {
		self.err = self.wrapError("", customError...)
		return self
	}
	return self
}
