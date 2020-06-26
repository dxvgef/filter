package filter

// 输入字符串类型的值
func String(value string, config ...Config) StrType {
	var str Str
	str.rawValue = value
	str.currentValue = value
	if len(config) > 0 {
		str.name = config[0].Name
		str.require = config[0].Require
		str.sep = config[0].Separator
	}
	return &str
}

// 必须有值(不能为零值)
func (self *Str) SetRequire() StrType {
	self.require = true
	return self
}

// SetSeparator 设置过滤器缓存中的分隔符，之后都以此分隔符来拆解value成slice
func (self *Str) SetSeparator(sep string) StrType {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.sep = sep
	return self
}

// 检查require
func (self *Str) checkRequire() {
	if self.require && self.currentValue == "" {
		self.err = wrapError(self.name, "不能为空")
	}
}

// 获得错误信息
func (self *Str) Error() error {
	self.checkRequire()
	return self.err
}

// // 设置默认值，如果出错则使用此值
// func (self *Str) SetDefault(value string) StrType {
// 	self.defaultValue = value
// 	return self
// }
