package filter

import "strings"

// 拼接string
func (self *Str) JoinStr(values ...string) *Str {
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

// JoinBytes 拼接bytes
func (self *Str) JoinBytes(values ...[]byte) *Str {
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

// JoinByte 拼接byte
func (self *Str) JoinByte(values ...byte) *Str {
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

// JoinRune 拼接rune
func (self *Str) JoinRune(values ...rune) *Str {
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
