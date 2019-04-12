package filter

import (
	"errors"
	"log"
	"reflect"
	"strconv"
)

// setFilter 赋值过滤器
type setFilter struct {
	Pointer interface{}
	Value   string
	Rules   []Rule
}

// 赋值过滤器
var SetFilter = func(pointer interface{}, value string, rules ...Rule) setFilter {
	return setFilter{
		Pointer: pointer,
		Value:   value,
		Rules:   rules,
	}
}

// Set 传入要接收过滤结果的变量指针，和要过滤的值
// 返回过滤结果
func Set(ptr interface{}, value string, rules ...Rule) error {
	if ptr == nil {
		return errors.New("参数不完整")
	}
	value, err := Get(value, rules...)
	if err != nil {
		return err
	}
	return setPtr(ptr, value)
}

// SetAll 验证所有传入的过滤器
// 返回首次过滤的结果
func SetAll(f ...setFilter) error {
	for k := range f {
		if err := Set(f[k].Pointer, f[k].Value, f[k].Rules...); err != nil {
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
