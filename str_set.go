package filter

import (
	"errors"
	"reflect"
)

func (self *Str) setCheck(target interface{}) (*reflect.Value, error) {
	if self.err != nil {
		return nil, self.err
	}

	// 检查目标是否为nil
	if target == nil {
		return nil, errors.New("the assignment target cannot be a null pointer")
	}

	targetValueOf := reflect.ValueOf(target)
	// 检查对象是否是指针
	if targetValueOf.Kind() != reflect.Ptr {
		return nil, errors.New("the assignment target must be a pointer")
	}
	// 检查对象是否能赋值
	if !targetValueOf.Elem().CanSet() {
		return nil, errors.New("cannot assign to the target")
	}
	return &targetValueOf, nil
}

// 赋值到普通对象
func (self *Str) Set(target interface{}, customError ...string) error {
	targetValueOf, err := self.setCheck(target)
	if err != nil {
		return self.err
	}

	// 开始赋值
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	switch targetTypeOf { //nolint
	case reflect.String:
		targetValueOf.Elem().SetString(self.currentValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value, err := self.Int64()
		if err != nil {
			return self.err
		}
		if targetValueOf.Elem().OverflowInt(value) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self.err
		}
		targetValueOf.Elem().SetInt(value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value, err := self.Uint64()
		if err != nil {
			return self.err
		}
		if targetValueOf.Elem().OverflowUint(value) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self.err
		}
		targetValueOf.Elem().SetUint(value)
	case reflect.Float32, reflect.Float64:
		value, err := self.Float64()
		if err != nil {
			return self.err
		}
		if targetValueOf.Elem().OverflowFloat(value) {
			self.err = wrapError(self.name, InvalidErrorText, customError...)
			return self.err
		}
		targetValueOf.Elem().SetFloat(value)
	case reflect.Bool:
		value, err := self.Bool()
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().SetBool(value)
	case reflect.Slice:
		return errors.New("the '" + targetTypeOf.String() + "' type must be assigned using the 'SetSlice' method")
	default:
		self.err = wrapError(self.name, "Cannot assign to the '"+targetTypeOf.String()+"' type")
		return self.err
	}
	return nil
}

func (self *Str) SetSlice(target interface{}, sep string, trimSpace bool, customError ...string) error {
	if sep == "" {
		self.err = wrapError(self.name, "`sep`parameter cannot be empty", customError...)
		return self.err
	}
	targetValueOf, err := self.setCheck(target)
	if err != nil {
		self.err = wrapError(self.name, err.Error(), customError...)
		return self.err
	}

	// 检查对象是否是指针
	if targetValueOf.Kind() != reflect.Ptr {
		self.err = wrapError(self.name, "the assignment target must be a pointer", customError...)
		return self.err
	}
	// 检查对象是否能赋值
	if !targetValueOf.Elem().CanSet() {
		self.err = wrapError(self.name, "cannot assign to the target", customError...)
		return self.err
	}

	// 如果不是slice类型则调用set
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	if targetTypeOf != reflect.Slice {
		return self.Set(target, customError...)
	}

	// 开始赋值
	sliceType := targetValueOf.Elem().Type().String()
	switch sliceType {
	case "[]string":
		value, err := self.SliceString(sep, trimSpace)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]int":
		value, err := self.SliceInt(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]int8":
		value, err := self.SliceInt8(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]int16":
		value, err := self.SliceInt16(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]int32":
		value, err := self.SliceInt32(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]int64":
		value, err := self.SliceInt64(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]uint":
		value, err := self.SliceUint(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]uint8":
		value, err := self.SliceUint8(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]uint16":
		value, err := self.SliceUint16(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]uint32":
		value, err := self.SliceUint32(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]uint64":
		value, err := self.SliceUint64(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]float32":
		value, err := self.SliceFloat32(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]float64":
		value, err := self.SliceFloat64(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	case "[]bool":
		value, err := self.SliceBool(sep)
		if err != nil {
			return self.err
		}
		targetValueOf.Elem().Set(reflect.ValueOf(value))
	default:
		self.err = wrapError(self.name, "Cannot assign to the '"+sliceType+"' type")
		return self.err
	}
	return nil
}
