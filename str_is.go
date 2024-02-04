package filter

import (
	"encoding/json"
	"net"
	"net/mail"
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

// Equal 与指定的字符串相等
func (self *Str) Equal(str string, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if self.currentValue != str {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// IsBool 布尔值
func (self *Str) IsBool(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if _, err := strconv.ParseBool(self.currentValue); err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// IsLower 小写字母
func (self *Str) IsLower(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsLower(v) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// IsUpper 大写字母
func (self *Str) IsUpper(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsUpper(v) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// IsLetter 大小写字母
func (self *Str) IsLetter(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsLetter(v) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// IsUnsigned 无符号数值
func (self *Str) IsUnsigned(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsDigit(v) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// IsLowerOrNumber 小写字母或数字
func (self *Str) IsLowerOrNumber(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsLower(v) && !unicode.IsDigit(v) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// IsUpperOrNumber 大写字母或数字
func (self *Str) IsUpperOrNumber(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsUpper(v) && !unicode.IsDigit(v) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// IsLetterOrNumber 字母或数字
func (self *Str) IsLetterOrNumber(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// IsChinese 汉字
func (self *Str) IsChinese(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.Is(unicode.Scripts["Han"], v) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// IsChinaTel 中国固定电话号码
func (self *Str) IsChinaTel(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	telSlice := strings.Split(self.currentValue, "-")
	if len(telSlice) != 2 {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	regionCode, err := strconv.Atoi(telSlice[0])
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	if regionCode < 10 || regionCode > 999 {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}

	code, err := strconv.Atoi(telSlice[1])
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	if code < 1000000 || code > 99999999 {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	return self
}

// IsMail 电邮地址
func (self *Str) IsMail(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	_, err := mail.ParseAddress(self.currentValue)
	if err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// IsIP IPv4/v6地址
func (self *Str) IsIP(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if net.ParseIP(self.currentValue) == nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// IsTCPAddr IP:Port
func (self *Str) IsTCPAddr(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if _, err := net.ResolveTCPAddr("tcp", self.currentValue); err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// IsMAC MAC地址
func (self *Str) IsMAC(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if _, err := net.ParseMAC(self.currentValue); err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}

	return self
}

// IsJSON JSON格式
func (self *Str) IsJSON(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	var js json.RawMessage
	if json.Unmarshal([]byte(self.currentValue), &js) != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	return self
}

// IsChinaIDNumber 中国身份证号码
func (self *Str) IsChinaIDNumber(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	var idV int
	if self.currentValue[17:] == "X" {
		idV = 88
	} else {
		var err error
		if idV, err = strconv.Atoi(self.currentValue[17:]); err != nil {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}

	var verify int
	id := self.currentValue[:17]
	arr := make([]int, 17)
	for i := 0; i < 17; i++ {
		arr[i], self.err = strconv.Atoi(string(id[i]))
		if self.err != nil {
			return self
		}
	}
	wi := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		res += arr[i] * wi[i]
	}
	verify = res % 11

	var temp int
	a18 := [11]int{1, 0, 88 /* 'X' */, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			break
		}
	}
	if temp != idV {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}

	return self
}

// IsChinaMobile 中国手机号码
func (self *Str) IsChinaMobile(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if len(self.currentValue) != 11 {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	var (
		prefix      uint8
		prefixValid bool
	)
	if prefix64, err := strconv.ParseUint(self.currentValue[0:3], 10, 8); err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	} else { //nolint
		prefix = uint8(prefix64)
	}
	for k := range chinaMobilePrefix {
		if chinaMobilePrefix[k] == prefix {
			prefixValid = true
			break
		}
	}
	if !prefixValid {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	if _, err := strconv.ParseUint(self.currentValue[3:], 10, 32); err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

func isSQLObject(value string) bool {
	for _, v := range value {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '-' && v != '_' {
			return false
		}
	}
	return true
}

// IsSQLObject SQL对象（库名、表名、字段名）
func (self *Str) IsSQLObject(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if !isSQLObject(self.currentValue) {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// IsSQLObjects SQL对象集合（库名、表名、字段名）
func (self *Str) IsSQLObjects(sep string, customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	values := strings.Split(self.currentValue, sep)
	if values[0] == "" || values[0] == " " {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for _, v := range values {
		if !isSQLObject(v) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
	}
	return self
}

// IsURL 是有效的URL
func (self *Str) IsURL(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if _, err := url.ParseRequestURI(self.currentValue); err != nil {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
	}
	return self
}

// IsUUID UUID格式
func (self *Str) IsUUID(customError ...string) *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}

	str := self.currentValue

	var uuid [16]byte
	switch len(str) {
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36:

	// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36 + 9:
		if strings.EqualFold(strings.ToLower(str[:9]), "urn:uuid:") {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
		str = str[9:]
	// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
	case 36 + 2:
		str = str[1:]
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	case 32:
		var ok bool
		for i := range uuid {
			uuid[i], ok = xtob(str[i*2], str[i*2+1])
			if !ok {
				self.err = wrapError(self.name, InvalidErrorText, customError...)
				return self
			}
		}
		return self
	default:
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	// s is now at least 36 bytes long
	// it must be of the form  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		self.err = wrapError(self.name, InvalidErrorText, customError...)
		return self
	}
	for i, x := range [16]int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
		v, ok := xtob(str[x], str[x+1])
		if !ok {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self
		}
		uuid[i] = v
	}

	return self
}
