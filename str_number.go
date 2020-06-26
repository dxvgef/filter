package filter

import "strconv"

// MinInteger 最小整数值
func (self *Str) MinInteger(min int64, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 64)
	if err != nil {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	if value64 < min {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	return self
}

// MinInteger 最大整数值
func (self *Str) MaxInteger(max int64, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseInt(self.currentValue, 10, 64)
	if err != nil {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	if value64 > max {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	return self
}

// MinFloat 最小浮点值
func (self *Str) MinFloat(min float64, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseFloat(self.currentValue, 64)
	if err != nil {
		self.err = wrapError(self.name, "", customError...)
		return self
	}

	if value64 < min {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	return self
}

// MinFloat 最大浮点值
func (self *Str) MaxFloat(max float64, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value64, err := strconv.ParseFloat(self.currentValue, 64)
	if err != nil {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	if value64 > max {
		self.err = wrapError(self.name, "", customError...)
		return self
	}
	return self
}
