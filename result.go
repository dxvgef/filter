package filter

import (
	"errors"
	"log"
	"reflect"
	"strconv"
)

// filter 赋值过滤器
type filter struct {
	Pointer interface{}
	Value   string
	Rules   []Rule
}

// 过滤器
var Filter = func(pointer interface{}, value string, rules ...Rule) filter {
	return filter{
		Pointer: pointer,
		Value:   value,
		Rules:   rules,
	}
}

// String 获得过滤后的字符串
func String(value string, rules ...Rule) (string, error) {
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

// Result 返回单个过滤和赋值的结果
func Result(ptr interface{}, value string, rules ...Rule) error {
	if ptr == nil {
		return errors.New("参数不完整")
	}
	value, err := String(value, rules...)
	if err != nil {
		return err
	}
	return setPtr(ptr, value)
}

// AllResult 返回多个过滤和赋值的结果
func AllResult(f ...filter) error {
	for k := range f {
		if err := Result(f[k].Pointer, f[k].Value, f[k].Rules...); err != nil {
			return err
		}
	}
	return nil
}

func setPtr(ptr interface{}, value string) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	ptrOf := reflect.ValueOf(ptr)
	if ptrOf.Kind() != reflect.Ptr {
		return errors.New("必须传入" + ptrOf.Kind().String() + "的指针")
	}

	if ptrOf.Elem().CanSet() == false {
		return errors.New("无法赋值到传入的变量")
	}

	ptrType := ptrOf.Elem().Type().Kind()
	switch ptrType {
	case reflect.String:
		ptrOf.Elem().SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return errors.New("无法将值转换为int类型")
		}
		if ptrOf.Elem().OverflowInt(v) == true {
			return errors.New("无法赋值给传入的参数类型")
		}
		ptrOf.Elem().SetInt(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return errors.New("无法将值转换为uint类型")
		}
		if ptrOf.Elem().OverflowUint(v) == true {
			return errors.New("无法赋值给传入的参数类型")
		}
		ptrOf.Elem().SetUint(v)
	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(value, 10)
		if err != nil {
			return errors.New("无法将值转换为float类型")
		}
		if ptrOf.Elem().OverflowFloat(v) == true {
			return errors.New("无法赋值给传入的参数类型")
		}
		ptrOf.Elem().SetFloat(v)
	case reflect.Bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return errors.New("无法将值转换为bool类型")
		}
		ptrOf.Elem().SetBool(v)
	}

	return nil
}
