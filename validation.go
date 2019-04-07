package filter

import (
	"errors"

	"github.com/dxvgef/filter/rule"
	"fmt"
)

// Result
func Result(value string, rules ...rule.Rule) error {
	// 运行修剪器
	for k := range rules {
		if rules[k].Trim != nil {
			value = rules[k].Trim(value)
		}
		if rules[k].StringOpts != nil {
			opts := rules[k].StringOpts
			for _, fn := range opts {
				value = fn(value)
				fmt.Println(value)
			}
		}

	}
	// 运行普通验证器
	for k := range rules {
		// 字符串长度验证
		if rules[k].MinLength > 0 {
			if result := rules[k].LengthValidate(value, rules[k].MinLength); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		if rules[k].MaxLength > 0 {
			if result := rules[k].LengthValidate(value, rules[k].MaxLength); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		// 验证值的范围
		if rules[k].MinIntegerValue != 0 {
			if result := rules[k].IntegerRangeValidate(value, rules[k].MinIntegerValue); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		if rules[k].MaxIntegerValue != 0 {
			if result := rules[k].IntegerRangeValidate(value, rules[k].MaxIntegerValue); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		if rules[k].MinFloatValue != 0 {
			if result := rules[k].FloatRangeValidate(value, rules[k].MinFloatValue); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
		if rules[k].MaxFloatValue != 0 {
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

		// string 值存在验证
		if rules[k].StringValue != nil {
			if result := rules[k].StringValueValidate(value, rules[k].StringValue); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
	}
	return nil
}
