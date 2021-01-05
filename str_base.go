package filter

// 输入字符串类型的值
func String(value string, name ...string) StrType {
	var str Str
	str.rawValue = value
	str.currentValue = value
	if len(name) > 0 {
		str.name = name[0]
	}
	return &str
}

// 必须有值(不能为零值)
func (self *Str) Require(customError ...string) StrType {
	self.require = true
	if len(customError) > 0 {
		self.requireErr = customError[0]
	}
	return self
}

// 检查require
func (self *Str) checkRequire() {
	if self.require && self.currentValue == "" {
		if self.requireErr != "" {
			self.err = wrapError(self.name, self.requireErr)
		} else {
			self.err = wrapError(self.name, "必要参数值不能为空")
		}
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
