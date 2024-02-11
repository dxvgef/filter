package filter

// Str 输入字符串类型
type Str struct {
	name         string // 变量名称，用于拼接错误消息
	rawValue     string // 原始值
	currentValue string // 当前值
	err          error  // 错误
	require      bool   // 不能为零值
	requireErr   string
}

// FromStr 输入字符串类型的值
func FromStr(value string, name ...string) *Str {
	var str Str
	str.rawValue = value
	str.currentValue = value
	if len(name) > 0 {
		str.name = name[0]
	}
	return &str
}

// Require 必须有值(不能为零值)
func (self *Str) Require(customError ...string) *Str {
	self.require = true
	if len(customError) > 0 {
		self.requireErr = customError[0]
	}
	self.checkRequire()
	return self
}

// 检查require
func (self *Str) checkRequire() {
	if self.currentValue != "" {
		return
	}
	if self.requireErr != "" {
		self.err = wrapError(self.name, self.requireErr)
		return
	}
	self.err = wrapError(self.name, InvalidErrorText)
}

// IsRequire 判断是否必须
func (self *Str) IsRequire() bool {
	return self.require
}

// Error 获得错误信息
func (self *Str) Error() error {
	return self.err
}

type CustomFunc func(string) (string, error)

// Custom 自定义处理方法
func (self *Str) Custom(f CustomFunc) *Str {
	if self.err != nil {
		return self
	}
	value, err := f(self.currentValue)
	if err != nil {
		self.err = err
		return self
	}
	self.currentValue = value
	return self
}
