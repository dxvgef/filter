package filter

// Rules 规则
type Rule struct {
	message string // 自定义失败消息

	formatter Formatter // 格式化器
	validator validator // 验证器
	// 长度
	length struct {
		validator LengthValidator
		min       int
		max       int
	}
	// 整数值范围
	intRange struct {
		validator IntRangeValidator
		min       int64
		max       int64
	}
	// 浮点值范围
	floatRange struct {
		validator FloatRangeValidator
		min       float64
		max       float64
	}
	// 必须存在于slice
	inSlice struct {
		validator inValidator
		slice     []string
	}
	// 必须存在指定字符串
	contains struct {
		validator twoStrValidator
		sub       string
	}
	// 必须存在指定的前缀字符串
	hasPrefix struct {
		validator twoStrValidator
		sub       string
	}
}

// Error 为验证器自定义错误消息
func (v Rule) Error(message string) Rule {
	v.message = message
	return v
}
