package rule

import "strconv"

// 整数值范围验证器
type IntRangeValidator func(string, int64) bool

//  浮点数值范围验证器
type FloatRangeValidator func(string, float64) bool

// NewMinIntegerValidator 创建整数最小值验证器
func NewMinIntegerValidator(min int64, validator IntRangeValidator, message string) Rule {
	return Rule{
		Message:              message,
		IntegerRangeValidate: validator,
		MinIntegerValue:      min,
	}
}

// NewMaxIntegerValidator 创建整数最大值验证器
func NewMaxIntegerValidator(max int64, validator IntRangeValidator, message string) Rule {
	return Rule{
		Message:              message,
		IntegerRangeValidate: validator,
		MaxIntegerValue:      max,
	}
}

// NewMinFloatValidator 创建浮点数最小值验证器
func NewMinFloatValidator(min float64, validator FloatRangeValidator, message string) Rule {
	return Rule{
		Message:            message,
		FloatRangeValidate: validator,
		MinFloatValue:      min,
	}
}

// NewMaxFloatValidator 创建浮点数最大值验证器
func NewMaxFloatValidator(max float64, validator FloatRangeValidator, message string) Rule {
	return Rule{
		Message:            message,
		FloatRangeValidate: validator,
		MaxFloatValue:      max,
	}
}

// 验证整数最小值
func minInteger(value string, min int64) bool {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false
	}
	if v < min {
		return false
	}
	return true
}

// 验证整数最大值
func maxInteger(value string, max int64) bool {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false
	}
	if v > max {
		return false
	}
	return true
}

// 验证浮点数最小值
func minFloat(value string, min float64) bool {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	if v < min {
		return false
	}
	return true
}

// 验证浮点数最大值
func maxFloat(value string, max float64) bool {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	if v > max {
		return false
	}
	return true
}
