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
	// 运行验证器
	for k := range rules {
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
