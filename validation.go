package filter

import (
	"errors"

	"github.com/dxvgef/filter/rule"
)

// Result
func Result(value string, rules ...rule.Rule) error {
	// 运行修剪器
	for k := range rules {
		if rules[k].Trimmer != nil {
			value = rules[k].Trimmer(value)
		}
	}
	// 运行普通验证器
	for k := range rules {
		// 字符串长度验证
		if rules[k].LengthValidate != nil && rules[k].MinLength > 0 {
			if result := rules[k].LengthValidate(value, rules[k].MinLength); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		if rules[k].LengthValidate != nil && rules[k].MaxLength > 0 {
			if result := rules[k].LengthValidate(value, rules[k].MaxLength); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		// 验证值的范围
		if rules[k].IntegerRangeValidate != nil && rules[k].MinIntegerValue != 0 {
			if result := rules[k].IntegerRangeValidate(value, rules[k].MinIntegerValue); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		if rules[k].IntegerRangeValidate != nil && rules[k].MaxIntegerValue != 0 {
			if result := rules[k].IntegerRangeValidate(value, rules[k].MaxIntegerValue); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		if rules[k].FloatRangeValidate != nil && rules[k].MinFloatValue != 0 {
			if result := rules[k].FloatRangeValidate(value, rules[k].MinFloatValue); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		if rules[k].FloatRangeValidate != nil && rules[k].MaxFloatValue != 0 {
			if result := rules[k].FloatRangeValidate(value, rules[k].MaxFloatValue); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		// 普通验证
		if rules[k].Validate != nil {
			if result := rules[k].Validate(value); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}

		// 验证值只能是slice中的值
		if rules[k].InValidate != nil {
			if result := rules[k].InValidate(value, rules[k].InValues); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}

		// 验证值必须存在指定字符
		if rules[k].ContainValidate != nil {

		}

	}
	return nil
}
