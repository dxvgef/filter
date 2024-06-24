package filter

import (
	"unicode/utf8"
)

// MatchLength 匹配长度
func (self *Str) MatchLength(value int, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if len(self.currentValue) != value {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// MinLength 最小长度
func (self *Str) MinLength(min int, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if len(self.currentValue) < min {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// MinUTF8Length UTF8编码最小长度
func (self *Str) MinUTF8Length(min int, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if utf8.RuneCountInString(self.currentValue) < min {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// MaxLength 最大长度
func (self *Str) MaxLength(max int, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if len(self.currentValue) > max {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// MaxUTF8Length UTF8编码最大长度
func (self *Str) MaxUTF8Length(max int, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if utf8.RuneCountInString(self.currentValue) > max {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}
