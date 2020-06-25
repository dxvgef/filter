package filter

import (
	"encoding/base64"
	"html"
	"net/url"
)

// base64 std编码
func (self *Str) Base64StdEncode() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = base64.StdEncoding.EncodeToString(strToBytes(self.value))
	return self
}

// base64 std解码
func (self *Str) Base64StdDecode(customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	bytes, err := base64.StdEncoding.DecodeString(self.value)
	if err != nil {
		self.err = self.wrapError(err, customError...)
		return self
	}
	self.value = bytesToStr(bytes)
	return self
}

// base64 RawStd编码
func (self *Str) Base64RawStdEncode() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = base64.RawStdEncoding.EncodeToString(strToBytes(self.value))
	return self
}

// base64 RawStd解码
func (self *Str) Base64RawStdDecode(customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	bytes, err := base64.RawStdEncoding.DecodeString(self.value)
	if err != nil {
		self.err = self.wrapError(err, customError...)
		return self
	}
	self.value = bytesToStr(bytes)
	return self
}

// base64 URL编码
func (self *Str) Base64URLEncode() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = base64.URLEncoding.EncodeToString(strToBytes(self.value))
	return self
}

// base64 URL解码
func (self *Str) Base64URLDecode(customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	bytes, err := base64.URLEncoding.DecodeString(self.value)
	if err != nil {
		self.err = self.wrapError(err, customError...)
		return self
	}
	self.value = bytesToStr(bytes)
	return self
}

// base64 RawURL编码
func (self *Str) Base64RawURLEncode() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = base64.RawURLEncoding.EncodeToString(strToBytes(self.value))
	return self
}

// base64 RawURL解码
func (self *Str) Base64RawURLDecode(customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	bytes, err := base64.RawURLEncoding.DecodeString(self.value)
	if err != nil {
		self.err = self.wrapError(err, customError...)
		return self
	}
	self.value = bytesToStr(bytes)
	return self
}

// html.UnescapeString
func (self *Str) HTMLUnescape() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = html.UnescapeString(self.value)
	return self
}

// html.EscapeString
func (self *Str) HTMLEscape() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = html.EscapeString(self.value)
	return self
}

// url.PathUnescape
func (self *Str) URLPathUnescape(customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	value, err := url.PathUnescape(self.value)
	if err != nil {
		self.err = self.wrapError(err, customError...)
		return self
	}
	self.value = value
	return self
}

// 与url.PathEscape相同
func (self *Str) URLPathEscape() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = url.PathEscape(self.value)
	return self
}

// 与url.QueryUnescape相同
func (self *Str) URLQueryUnescape(customError ...string) StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	value, err := url.QueryUnescape(self.value)
	if err != nil {
		self.err = self.wrapError(err, customError...)
		return self
	}
	self.value = value
	return self
}

// 与url.QueryEscape相同
func (self *Str) URLQueryEscape() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = url.QueryEscape(self.value)
	return self
}
