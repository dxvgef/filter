package filter

import "strings"

// 删除左右的指定字符串
func (self *Str) Trim(sub string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = strings.Trim(self.value, sub)
	return self
}

// 删除左右的空格
func (self *Str) TrimSpace() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = strings.TrimSpace(self.value)
	return self
}

// 删除左边的指定字符串
func (self *Str) TrimLeft(sub string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = strings.TrimLeft(self.value, sub)
	return self
}

// 删除右边的指定字符串
func (self *Str) TrimRight(sub string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = strings.TrimRight(self.value, sub)
	return self
}

// 删除指定的前缀字符串
func (self *Str) TrimPrefix(sub string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = strings.TrimPrefix(self.value, sub)
	return self
}

// 删除指定的后缀字符串
func (self *Str) TrimSuffix(sub string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = strings.TrimSuffix(self.value, sub)
	return self
}

// 删除所有空格
func (self *Str) RemoveSpace(sub string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = strings.ReplaceAll(self.value, " ", "")
	return self
}
