package filter

// twoStrValidator 两个字符串入参的验证器
type twoStrValidator func(string, string) bool

// inValidator 值必须存在于slice的验证器
type inValidator func(string, []string) bool

// NewInValidator 创建值必须存在于slice的验证器
func NewInValidator(slice []string, validator inValidator, message string) Rule {
	return Rule{
		message: message,
		inSlice: struct {
			validator inValidator
			slice     []string
		}{validator: validator, slice: slice},
	}
}

// NewContainValidator 创建Contains的验证器
func NewContainValidator(sub string, validator twoStrValidator, message string) Rule {
	return Rule{
		message: message,
		contains: struct {
			validator twoStrValidator
			sub       string
		}{validator: validator, sub: sub},
	}
}

// NewHasPrefixValidator 创建HasPrefix的验证器
func NewHasPrefixValidator(sub string, validator twoStrValidator, message string) Rule {
	return Rule{
		message: message,
		contains: struct {
			validator twoStrValidator
			sub       string
		}{validator: validator, sub: sub},
	}
}
