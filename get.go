package filter

import (
	"errors"
)

// filter 过滤器
type filter struct {
	Value string
	Rules []Rule
}

// Filter 过滤器
var Filter = func(value string, rules ...Rule) filter {
	return filter{
		Value: value,
		Rules: rules,
	}
}

// GetAll 获得所有要过滤器的值
func GetAll(f ...filter) error {
	for k := range f {
		if _, err := Get(f[k].Value, f[k].Rules...); err != nil {
			return err
		}
	}
	return nil
}

// Get 获得要过滤的值
func Get(value string, rules ...Rule) (string, error) {
	// 运行修剪器
	for k := range rules {
		if rules[k].formatter != nil {
			value = rules[k].formatter(value)
		}
	}

	// 运行普通验证器
	for k := range rules {
		// 字符串长度验证
		if rules[k].length.validator != nil && rules[k].length.min > 0 {
			if result := rules[k].length.validator(value, rules[k].length.min); result != true {
				return "", errors.New(rules[k].message)
			}
		}
		if rules[k].length.validator != nil && rules[k].length.max > 0 {
			if result := rules[k].length.validator(value, rules[k].length.max); result != true {
				return "", errors.New(rules[k].message)
			}
		}
		// 验证值的范围
		if rules[k].intRange.validator != nil && rules[k].intRange.min != 0 {
			if result := rules[k].intRange.validator(value, rules[k].intRange.min); result != true {
				return "", errors.New(rules[k].message)
			}
		}
		if rules[k].intRange.validator != nil && rules[k].intRange.max != 0 {
			if result := rules[k].intRange.validator(value, rules[k].intRange.max); result != true {
				return "", errors.New(rules[k].message)
			}
		}
		if rules[k].floatRange.validator != nil && rules[k].floatRange.min != 0 {
			if result := rules[k].floatRange.validator(value, rules[k].floatRange.min); result != true {
				return "", errors.New(rules[k].message)
			}
		}
		if rules[k].floatRange.validator != nil && rules[k].floatRange.max != 0 {
			if result := rules[k].floatRange.validator(value, rules[k].floatRange.max); result != true {
				return "", errors.New(rules[k].message)
			}
		}
		// 普通验证
		if rules[k].validator != nil {
			if result := rules[k].validator(value); result != true {
				return "", errors.New(rules[k].message)
			}
		}

		// 验证值只能是slice中的值
		if rules[k].inSlice.validator != nil {
			if result := rules[k].inSlice.validator(value, rules[k].inSlice.slice); result != true {
				return "", errors.New(rules[k].message)
			}
		}

		// 验证值必须存在指定字符串
		if rules[k].contains.validator != nil && rules[k].contains.sub != "" {
			if result := rules[k].contains.validator(value, rules[k].contains.sub); result != true {
				return "", errors.New(rules[k].message)
			}
		}

		// 验证值必须存在指定的前缀字符串
		if rules[k].hasPrefix.validator != nil && rules[k].hasPrefix.sub != "" {
			if result := rules[k].hasPrefix.validator(value, rules[k].hasPrefix.sub); result != true {
				return "", errors.New(rules[k].message)
			}
		}

	}
	return value, nil
}
