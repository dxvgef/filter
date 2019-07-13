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
	if source.err != nil {
		if source.required == true && source.rawValue == "" {
			return source.requiredError
		}
		return source.err
	}

	ptrOf := reflect.ValueOf(target)
	if ptrOf.Kind() != reflect.Ptr {
		return errors.New("必须传入" + ptrOf.Kind().String() + "的指针")
	}

	if ptrOf.Elem().CanSet() == false {
		return errors.New("无法赋值到传入的变量")
	}

	ptrType := ptrOf.Elem().Type().Kind()
	switch ptrType {
	case reflect.String:
		ptrOf.Elem().SetString(source.rawValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(source.rawValue, 10, 64)
		if err != nil {
			return errors.New("无法将值转换为int类型")
		}
		if ptrOf.Elem().OverflowInt(v) == true {
			return errors.New("无法赋值给传入的参数类型")
		}
		ptrOf.Elem().SetInt(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := strconv.ParseUint(source.rawValue, 10, 64)
		if err != nil {
			return errors.New("无法将值转换为uint类型")
		}
		if ptrOf.Elem().OverflowUint(v) == true {
			return errors.New("无法赋值给传入的参数类型")
		}
		ptrOf.Elem().SetUint(v)
	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(source.rawValue, 10)
		if err != nil {
			return errors.New("无法将值转换为float类型")
		}
		if ptrOf.Elem().OverflowFloat(v) == true {
			return errors.New("无法赋值给传入的参数类型")
		}
	case reflect.Bool:
		v, err := strconv.ParseBool(source.rawValue)
		if err != nil {
			return errors.New("无法将值转换为bool类型")
		}
		ptrOf.Elem().SetBool(v)
	case reflect.Slice:
		sliceType := ptrOf.Elem().Type().String()
		switch sliceType {
		case "[]string":
			if source.sep == "" {
				return errors.New("无法赋值到传入的变量，分隔符sep参数未定义")
			}
			ptrOf.Elem().Set(reflect.ValueOf(strings.Split(source.rawValue, source.sep)))
		}
	default:
		return errors.New("无法赋值到传入的变量，不是预期的值类型")
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
