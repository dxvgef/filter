package filter

// 验证器
type validator func(string) bool

// NewValidator 创建新验证器
func NewValidator(validator validator, message string) Rule {
	return Rule{
		validator: validator,
		message:   message,
	}
}

// 格式化器
type Formatter func(string) string

// NewFormatter 创建新的格式化器
func NewFormatter(formatter Formatter) Rule {
	return Rule{
		formatter: formatter,
	}
}

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

// 整数值范围验证器
type IntRangeValidator func(string, int64) bool

//  浮点数值范围验证器
type FloatRangeValidator func(string, float64) bool

// NewMinIntegerValidator 创建最大整数值验证器
func NewMinIntegerValidator(min int64, validator IntRangeValidator, message string) Rule {
	return Rule{
		message: message,
		intRange: struct {
			validator IntRangeValidator
			min       int64
			max       int64
		}{validator: validator, min: min},
	}
}

// NewMaxIntegerValidator 创建最小整数值验证器
func NewMaxIntegerValidator(max int64, validator IntRangeValidator, message string) Rule {
	return Rule{
		message: message,
		intRange: struct {
			validator IntRangeValidator
			min       int64
			max       int64
		}{validator: validator, max: max},
	}
}

// NewMinFloatValidator 创建最大浮点值验证器
func NewMinFloatValidator(min float64, validator FloatRangeValidator, message string) Rule {
	return Rule{
		message: message,
		floatRange: struct {
			validator FloatRangeValidator
			min       float64
			max       float64
		}{validator: validator, min: min},
	}
}

// NewMaxFloatValidator 创建最小浮点值验证器
func NewMaxFloatValidator(max float64, validator FloatRangeValidator, message string) Rule {
	return Rule{
		message: message,
		floatRange: struct {
			validator FloatRangeValidator
			min       float64
			max       float64
		}{validator: validator, max: max},
	}
}
