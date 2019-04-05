package filter

import (
	"errors"

	"github.com/dxvgef/filter/rule"
)

// Result
func Result(value string, rules ...rule.Rule) error {
	// 运行修剪器
	for k := range rules {
		if rules[k].Trim != nil {
			value = rules[k].Trim(value)
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
		// 普通验证
		if rules[k].Validate != nil {
			if result := rules[k].Validate(value); result != true {
				if rules[k].Message != "" {
					return errors.New(rules[k].Message)
				}
				return errors.New(rules[k].Message)
			}
		}
	}
	return nil
}
