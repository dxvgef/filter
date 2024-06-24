package filter

import (
	"strings"
)

// 字母转为大写
func (self *Str) ToUpper() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.ToUpper(self.currentValue)
	return self
}

// 字母转为小写
func (self *Str) ToLower() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.ToLower(self.currentValue)
	return self
}
