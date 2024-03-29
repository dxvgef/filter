package filter

import "strings"

// Trim 删除左右的指定字符串
func (self *Str) Trim(sub string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.Trim(self.currentValue, sub)
	return self
}

// TrimSpace 删除左右的空格
func (self *Str) TrimSpace() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.TrimSpace(self.currentValue)
	return self
}

// TrimLeft 删除左边的指定字符串
func (self *Str) TrimLeft(sub string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.TrimLeft(self.currentValue, sub)
	return self
}

// TrimRight 删除右边的指定字符串
func (self *Str) TrimRight(sub string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.TrimRight(self.currentValue, sub)
	return self
}

// TrimPrefix 删除指定的前缀字符串
func (self *Str) TrimPrefix(sub string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.TrimPrefix(self.currentValue, sub)
	return self
}

// TrimSuffix 删除指定的后缀字符串
func (self *Str) TrimSuffix(sub string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.TrimSuffix(self.currentValue, sub)
	return self
}

// RemoveSpace 删除所有空格
func (self *Str) RemoveSpace() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.ReplaceAll(self.currentValue, " ", "")
	return self
}
