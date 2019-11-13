package filter

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// element 元素
type element struct {
	target interface{}
	source *Object
}

// El 设置元素
var El = func(target interface{}, source *Object) element {
	return element{
		target: target,
		source: source,
	}
}

// Set 将过滤结果赋值到指定对象
func Set(target interface{}, source *Object) error {
	if source.err != nil && source.defaultValue == nil {
		if source.silent {
			return nil
		}
		return source.err
	}

	// source.err = nil

	targetValueOf := reflect.ValueOf(target)
	if targetValueOf.Kind() != reflect.Ptr {
		if source.silent {
			return nil
		}
		return errors.New(source.name + "过滤规则的target参数必须传入指针类型")
	}
	if !targetValueOf.Elem().CanSet() {
		if source.silent {
			return nil
		}
		return errors.New(source.name + "过滤规则无法更改目标变量的值")
	}
	targetTypeOf := targetValueOf.Elem().Type().Kind()

	if source.err != nil && source.defaultValue != nil {
		if err := setDefaultValue(targetTypeOf, targetValueOf, source.defaultValue); err != nil {
			if source.silent {
				return nil
			}
			return err
		}
		return nil
	}

	if source.err == nil && source.rawValue != "" {
		if err := setRawValue(targetTypeOf, targetValueOf, source.rawValue, source.sep); err != nil {
			if source.silent {
				return nil
			}
			return errors.New(source.name + err.Error())
		}
	}

	return nil
}

// MSet 将多个过滤结果赋值到指定对象
func MSet(elements ...element) error {
	for k := range elements {
		err := Set(elements[k].target, elements[k].source)
		if err != nil {
			return err
		}
	}
	return nil
}

// 写target的值，值类型是string
func setRawValue(targetTypeOf reflect.Kind, targetValueOf reflect.Value, value string, sep string) error {
	switch targetTypeOf {
	case reflect.String:
		targetValueOf.Elem().SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return errors.New("必须是整数")
		}
		if targetValueOf.Elem().OverflowInt(v) {
			return errors.New("不能用" + targetTypeOf.String() + "类型赋值")
		}
		targetValueOf.Elem().SetInt(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return errors.New("必须是无符号整数")
		}
		if targetValueOf.Elem().OverflowUint(v) {
			return errors.New("不能用" + targetTypeOf.String() + "类型赋值")
		}
		targetValueOf.Elem().SetUint(v)
	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(value, 10)
		if err != nil {
			return errors.New("必须是小数")
		}
		if targetValueOf.Elem().OverflowFloat(v) {
			return errors.New("不能用" + targetTypeOf.String() + "类型赋值")
		}
	case reflect.Bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return errors.New("必须是布尔值")
		}
		targetValueOf.Elem().SetBool(v)
	case reflect.Slice:
		sliceType := targetValueOf.Elem().Type().String()
		if sliceType == "[]string" {
			if sep == "" {
				return errors.New("过滤规则的分隔符参数(sep)未定义")
			}
			targetValueOf.Elem().Set(reflect.ValueOf(strings.Split(value, sep)))
		}
	default:
		return errors.New("不能用" + targetTypeOf.String() + "类型赋值")
	}

	return nil
}

// 写target的值，值类型是interface
func setDefaultValue(targetTypeOf reflect.Kind, targetValueOf reflect.Value, value interface{}) error {
	valueTypeOf := reflect.TypeOf(value)
	if valueTypeOf.Kind() != targetTypeOf {
		return errors.New("值类型不相同")
	}
	targetValueOf.Elem().Set(reflect.ValueOf(value))
	return nil
}
