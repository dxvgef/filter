package filter

import (
	"encoding/json"
	"net"
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

// Equal 与指定的字符串相等
func (self *Str) Equal(str string, customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if self.currentValue != str {
		self.err = wrapError(self.name, "值不匹配", customError...)
	}
	return self
}

// IsBool 布尔值
func (self *Str) IsBool(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if _, err := strconv.ParseBool(self.currentValue); err != nil {
		self.err = wrapError(self.name, "必须是布尔类型", customError...)
	}
	return self
}

// IsLower 小写字母
func (self *Str) IsLower(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsLower(v) {
			self.err = wrapError(self.name, "必须是小写字母", customError...)
			return self
		}
	}
	return self
}

// IsUpper 大写字母
func (self *Str) IsUpper(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsUpper(v) {
			self.err = wrapError(self.name, "必须是大写字母", customError...)
			return self
		}
	}
	return self
}

// IsLetter 大小写字母
func (self *Str) IsLetter(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsLetter(v) {
			self.err = wrapError(self.name, "必须是字母", customError...)
			return self
		}
	}
	return self
}

// IsDigit 无符号数字
func (self *Str) IsDigit(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsDigit(v) {
			self.err = wrapError(self.name, "必须是无符号数字", customError...)
			return self
		}
	}
	return self
}

// IsLowerOrDigit 小写字母或数字
func (self *Str) IsLowerOrDigit(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsLower(v) && !unicode.IsDigit(v) {
			self.err = wrapError(self.name, "必须是小写字母或无符号数字", customError...)
			return self
		}
	}
	return self
}

// IsUpperOrDigit 大写字母或数字
func (self *Str) IsUpperOrDigit(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsUpper(v) && !unicode.IsDigit(v) {
			self.err = wrapError(self.name, "必须是大写字母或无符号数字", customError...)
			return self
		}
	}
	return self
}

// IsLetterOrDigit 字母或数字
func (self *Str) IsLetterOrDigit(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			self.err = wrapError(self.name, "必须是字母或无符号数字", customError...)
			return self
		}
	}
	return self
}

// IsChinese 汉字
func (self *Str) IsChinese(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	for _, v := range self.currentValue {
		if !unicode.Is(unicode.Scripts["Han"], v) {
			self.err = wrapError(self.name, "必须是汉字", customError...)
			return self
		}
	}
	return self
}

// IsChinaTel 中国固定电话号码
func (self *Str) IsChinaTel(customError ...string) StrType {
	errMsg := "不是有效的中国大陆固定电话号码"
	if self.err != nil || self.currentValue == "" {
		return self
	}
	telSlice := strings.Split(self.currentValue, "-")
	if len(telSlice) != 2 {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	regionCode, err := strconv.Atoi(telSlice[0])
	if err != nil {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	if regionCode < 10 || regionCode > 999 {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}

	code, err := strconv.Atoi(telSlice[1])
	if err != nil {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	if code < 1000000 || code > 99999999 {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	return self
}

// IsMail 电邮地址
func (self *Str) IsMail(customError ...string) StrType {
	errMsg := "不是有效的电邮地址"
	if self.err != nil || self.currentValue == "" {
		return self
	}
	emailSlice := strings.Split(self.currentValue, "@")
	if len(emailSlice) != 2 {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	if emailSlice[0] == "" || emailSlice[1] == "" {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}

	for k, v := range emailSlice[0] {
		if k == 0 && !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			self.err = wrapError(self.name, errMsg, customError...)
			return self
		} else if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '@' && v != '.' && v != '_' && v != '-' {
			self.err = wrapError(self.name, errMsg, customError...)
			return self
		}
	}

	domainSlice := strings.Split(emailSlice[1], ".")
	if len(domainSlice) < 2 {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	domainSliceLen := len(domainSlice)
	for i := 0; i < domainSliceLen; i++ {
		for k, v := range domainSlice[i] {
			// nolint
			if i != domainSliceLen-1 && k == 0 && !unicode.IsLetter(v) && !unicode.IsDigit(v) {
				self.err = wrapError(self.name, errMsg, customError...)
				return self
			} else if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '.' && v != '_' && v != '-' {
				self.err = wrapError(self.name, errMsg, customError...)
				return self
			} else if i == domainSliceLen-1 && !unicode.IsLetter(v) {
				self.err = wrapError(self.name, errMsg, customError...)
				return self
			}
		}
	}

	return self
}

// IsIP IPv4/v6地址
func (self *Str) IsIP(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if net.ParseIP(self.currentValue) == nil {
		self.err = wrapError(self.name, "不是有效的IP地址", customError...)
		return self
	}

	return self
}

// IsTCPAddr IP:Port
func (self *Str) IsTCPAddr(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if _, err := net.ResolveTCPAddr("tcp", self.currentValue); err != nil {
		self.err = wrapError(self.name, "不是有效的TCP地址(IP:Port)", customError...)
		return self
	}

	return self
}

// IsMAC MAC地址
func (self *Str) IsMAC(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if _, err := net.ParseMAC(self.currentValue); err != nil {
		self.err = wrapError(self.name, "不是有效的MAC地址", customError...)
		return self
	}

	return self
}

// IsJSON JSON格式
func (self *Str) IsJSON(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	var js json.RawMessage
	if json.Unmarshal([]byte(self.currentValue), &js) != nil {
		self.err = wrapError(self.name, "不是有效的JSON字符串", customError...)
		return self
	}
	return self
}

// IsChinaIDNumber 中国身份证号码
func (self *Str) IsChinaIDNumber(customError ...string) StrType {
	errMsg := "不是有效的中国大陆身份证号码"
	if self.err != nil || self.currentValue == "" {
		return self
	}
	var idV int
	if self.currentValue[17:] == "X" {
		idV = 88
	} else {
		var err error
		if idV, err = strconv.Atoi(self.currentValue[17:]); err != nil {
			self.err = wrapError(self.name, errMsg, customError...)
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
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}

	return self
}

// nolint:gocyclo
// IsChinaMobile 中国手机号码
func (self *Str) IsChinaMobile(customError ...string) StrType {
	errMsg := "不是有效的中国大陆手机号码"
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if len(self.currentValue) != 11 {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	var (
		prefix      uint8
		prefixValid bool
	)
	if prefix64, err := strconv.ParseUint(self.currentValue[0:3], 10, 8); err != nil {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	} else {
		prefix = uint8(prefix64)
	}
	for k := range chinaMobilePrefix {
		if chinaMobilePrefix[k] == prefix {
			prefixValid = true
			break
		}
	}
	if !prefixValid {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	if _, err := strconv.ParseUint(self.currentValue[3:], 10, 32); err != nil {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
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
func (self *Str) IsSQLObject(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if !isSQLObject(self.currentValue) {
		self.err = wrapError(self.name, "不是有效的SQL对象名", customError...)
		return self
	}
	return self
}

// IsSQLObjects SQL对象集合（库名、表名、字段名）
func (self *Str) IsSQLObjects(sep string, customError ...string) StrType {
	errMsg := "不是有效的SQL对象合集"
	if self.err != nil || self.currentValue == "" {
		return self
	}
	values := strings.Split(self.currentValue, sep)
	if values[0] == "" || values[0] == " " {
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	for _, v := range values {
		if !isSQLObject(v) {
			self.err = wrapError(self.name, errMsg, customError...)
			return self
		}
	}
	return self
}

// IsURL 是有效的URL
func (self *Str) IsURL(customError ...string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	if _, err := url.ParseRequestURI(self.currentValue); err != nil {
		self.err = wrapError(self.name, "不是有效的URL", customError...)
	}
	return self
}

// IsUUID UUID格式
func (self *Str) IsUUID(customError ...string) StrType {
	errMsg := "不是有效的UUID"
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
			self.err = wrapError(self.name, errMsg, customError...)
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
				self.err = wrapError(self.name, errMsg, customError...)
				return self
			}
		}
		return self
	default:
		self.err = wrapError(self.name, errMsg, customError...)
		return self
	}
	// s is now at least 36 bytes long
	// it must be of the form  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		self.err = wrapError(self.name, errMsg, customError...)
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
			self.err = wrapError(self.name, errMsg, customError...)
			return self
		}
		uuid[i] = v
	}

	return self
}
