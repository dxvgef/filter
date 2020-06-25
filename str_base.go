package filter

import (
	"errors"
	"strings"
)

// 输入字符串类型的值
func InputString(value string, name ...string) StrType {
	var str Str
	str.rawValue = value
	str.currentValue = value
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

// 检查require
func (self *Str) checkRequire() {
	if self.require && self.currentValue == "" {
		self.err = self.wrapError("不能为空")
	}
}

// 获得错误信息
func (self *Str) Error() error {
	self.checkRequire()
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
		body.WriteString(DefaultErrorText)
	}
	return errors.New(body.String())
}

// // 设置默认值，如果出错则使用此值
// func (self *Str) SetDefault(value string) StrType {
// 	self.defaultValue = value
// 	return self
// }
