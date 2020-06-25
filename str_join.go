package filter

import "strings"

// 拼接string
func (self *Str) JoinStr(values ...string) StrType {
	if self.err != nil {
		return self
	}
	var value strings.Builder
	value.WriteString(self.currentValue)
	for k := range values {
		value.WriteString(values[k])
	}
	self.currentValue = value.String()
	return self
}

// 拼接bytes
func (self *Str) JoinBytes(values ...[]byte) StrType {
	if self.err != nil {
		return self
	}
	var value strings.Builder
	value.WriteString(self.currentValue)
	for k := range values {
		value.Write(values[k])
	}
	self.currentValue = value.String()
	return self
}

// 拼接byte
func (self *Str) JoinByte(values ...byte) StrType {
	if self.err != nil {
		return self
	}
	var value strings.Builder
	value.WriteString(self.currentValue)
	for k := range values {
		value.WriteByte(values[k])
	}
	self.currentValue = value.String()
	return self
}

// 拼接rune
func (self *Str) JoinRune(values ...rune) StrType {
	if self.err != nil {
		return self
	}
	var value strings.Builder
	value.WriteString(self.currentValue)
	for k := range values {
		value.WriteRune(values[k])
	}
	self.currentValue = value.String()
	return self
}
