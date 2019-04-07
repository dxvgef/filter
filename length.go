package filter

// 长度验证器
type LengthValidator func(string, int) bool

// NewMinLengthValidator 创建字符串最小长度验证器
func NewMinLengthValidator(min int, validator LengthValidator, message string) Rule {
	return Rule{
		message: message,
		length: struct {
			validator LengthValidator
			min       int
			max       int
		}{validator: validator, min: min},
	}
}

// NewMaxLengthValidator 创建字符串最大长度验证器
func NewMaxLengthValidator(max int, validator LengthValidator, message string) Rule {
	return Rule{
		message: message,
		length: struct {
			validator LengthValidator
			min       int
			max       int
		}{validator: validator, max: max},
	}
}
