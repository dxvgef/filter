package filter

import (
	"encoding/base64"
	"html"
	"net/url"
)

// base64 std编码
func (self *Str) Base64StdEncode() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = base64.StdEncoding.EncodeToString(strToBytes(self.currentValue))
	return self
}

// base64 std解码
func (self *Str) Base64StdDecode(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	bytes, err := base64.StdEncoding.DecodeString(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	self.currentValue = bytesToStr(bytes)
	return self
}

// base64 RawStd编码
func (self *Str) Base64RawStdEncode() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = base64.RawStdEncoding.EncodeToString(strToBytes(self.currentValue))
	return self
}

// base64 RawStd解码
func (self *Str) Base64RawStdDecode(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	bytes, err := base64.RawStdEncoding.DecodeString(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	self.currentValue = bytesToStr(bytes)
	return self
}

// base64 URL编码
func (self *Str) Base64URLEncode() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = base64.URLEncoding.EncodeToString(strToBytes(self.currentValue))
	return self
}

// base64 URL解码
func (self *Str) Base64URLDecode(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	bytes, err := base64.URLEncoding.DecodeString(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	self.currentValue = bytesToStr(bytes)
	return self
}

// base64 RawURL编码
func (self *Str) Base64RawURLEncode() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = base64.RawURLEncoding.EncodeToString(strToBytes(self.currentValue))
	return self
}

// base64 RawURL解码
func (self *Str) Base64RawURLDecode(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	bytes, err := base64.RawURLEncoding.DecodeString(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	self.currentValue = bytesToStr(bytes)
	return self
}

// html.UnescapeString
func (self *Str) HTMLUnescape() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = html.UnescapeString(self.currentValue)
	return self
}

// html.EscapeString
func (self *Str) HTMLEscape() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = html.EscapeString(self.currentValue)
	return self
}

// url.PathUnescape
func (self *Str) URLPathUnescape(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := url.PathUnescape(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	self.currentValue = value
	return self
}

// 与url.PathEscape相同
func (self *Str) URLPathEscape() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = url.PathEscape(self.currentValue)
	return self
}

// 与url.QueryUnescape相同
func (self *Str) URLQueryUnescape(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	value, err := url.QueryUnescape(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	self.currentValue = value
	return self
}

// URLQueryEscape 与url.QueryEscape相同
func (self *Str) URLQueryEscape() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = url.QueryEscape(self.currentValue)
	return self
}
