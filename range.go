package filter

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
