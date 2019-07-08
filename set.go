package filter

import (
	"errors"
	"log"
	"reflect"
	"strconv"
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
		return source.err
	}
	return setTarget(target, source.rawValue)
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

func setTarget(ptr interface{}, value string) error {
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
