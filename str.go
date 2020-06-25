package filter

import (
	"errors"
	"strings"
)

type StrType interface {
	Raw() string
	Error() error
	JoinStr(...string) StrType
	JoinBytes(...[]byte) StrType
	JoinByte(...byte) StrType
	JoinRune(...rune) StrType
	Trim(string) StrType
	TrimSpace() StrType
	TrimLeft(string) StrType
	TrimRight(string) StrType
	TrimPrefix(string) StrType
	TrimSuffix(string) StrType
}

type Str struct {
	name         string  // 变量名称，用于拼接错误消息
	rawValue     string  // 原始值
	currentValue string  // 当前值
	defaultValue string  // 默认值
	err          error   // 错误
	int64        int64   // 转成int64类型的值，用于做数值运算，减少类型转换的次数
	float64      float64 // 转成float64类型的值，用于做数值运算，减少类型转换的次数
	require      bool    // 不能为零值
	sep          string  // 分隔符
}

// 输入字符串类型的值
func InputString(value string, name ...string) StrType {
	var str Str
	str.rawValue = value
	if len(name) > 0 {
		str.name = name[0]
	}
	return &str
}

// 必须有值(不能为零值)
func (self *Str) Require() StrType {
	self.require = true
	return self
}

// 获得原始值
func (self *Str) Raw() string {
	return self.rawValue
}

// 获得错误信息
func (self *Str) Error() error {
	return self.err
}

// 封装错误信息
// nolint:unparam
func (self *Str) wrapError(err string, custom ...string) error {
	var body strings.Builder
	body.WriteString(self.name)
	body.WriteString(": ")

	// nolint:gocritic
	if len(custom) > 0 && custom[0] != "" {
		body.WriteString(custom[0])
	} else if err != "" {
		body.WriteString(err)
	} else {
		body.WriteString(defaultError)
	}
	return errors.New(body.String())
}

// 设置默认值，如果出错则使用此值
func (self *Str) SetDefault(value string) StrType {
	self.defaultValue = value
	return self
}
