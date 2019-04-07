package filter

import (
	"errors"
)

// 验证器
type validator func(string) bool

// NewValidator 创建新验证器
func NewValidator(validator validator, message string) Rule {
	return Rule{
		validator: validator,
		message:   message,
	}
}

// Result
func Result(value string, rules ...Rule) error {
	// 运行修剪器
	for k := range rules {
		if rules[k].trimmer != nil {
			value = rules[k].trimmer(value)
		}
	}

	// 运行普通验证器
	for k := range rules {
		// 字符串长度验证
		if rules[k].length.validator != nil && rules[k].length.min > 0 {
			if result := rules[k].length.validator(value, rules[k].length.min); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}
		if rules[k].length.validator != nil && rules[k].length.max > 0 {
			if result := rules[k].length.validator(value, rules[k].length.max); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}
		// 验证值的范围
		if rules[k].intRange.validator != nil && rules[k].intRange.min != 0 {
			if result := rules[k].intRange.validator(value, rules[k].intRange.min); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}
		if rules[k].intRange.validator != nil && rules[k].intRange.max != 0 {
			if result := rules[k].intRange.validator(value, rules[k].intRange.max); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}
		if rules[k].floatRange.validator != nil && rules[k].floatRange.min != 0 {
			if result := rules[k].floatRange.validator(value, rules[k].floatRange.min); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}
		if rules[k].floatRange.validator != nil && rules[k].floatRange.max != 0 {
			if result := rules[k].floatRange.validator(value, rules[k].floatRange.max); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}
		// 普通验证
		if rules[k].validator != nil {
			if result := rules[k].validator(value); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}

		// 验证值只能是slice中的值
		if rules[k].inSlice.validator != nil {
			if result := rules[k].inSlice.validator(value, rules[k].inSlice.slice); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}

		// 验证值必须存在指定字符串
		if rules[k].contains.validator != nil && rules[k].contains.sub != "" {
			if result := rules[k].contains.validator(value, rules[k].contains.sub); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}

		// 验证值必须存在指定的前缀字符串
		if rules[k].hasPrefix.validator != nil && rules[k].hasPrefix.sub != "" {
			if result := rules[k].hasPrefix.validator(value, rules[k].hasPrefix.sub); result != true {
				if rules[k].message != "" {
					return errors.New(rules[k].message)
				}
				return errors.New(rules[k].message)
			}
		}

	}
	return nil
}
