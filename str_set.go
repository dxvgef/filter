package filter

import "reflect"

// 赋值到对象
// nolint:gocyclo
func (self *Str) Set(target interface{}, customError ...string) error {
	self.checkRequire()
	if self.err != nil {
		return self.err
	}

	// 检查目标是否为nil
	if target == nil {
		self.err = wrapError(self.name, "target cannot be nil", customError...)
		return self.err
	}

	targetValueOf := reflect.ValueOf(target)
	// 检查对象是否是指针
	if targetValueOf.Kind() != reflect.Ptr {
		self.err = wrapError(self.name, "target must be a pointer", customError...)
		return self.err
	}
	// 检查对象是否能赋值
	if !targetValueOf.Elem().CanSet() {
		self.err = wrapError(self.name, "cannot set the value of the target", customError...)
		return self.err
	}

	// 开始赋值
	targetTypeOf := targetValueOf.Elem().Type().Kind()
	switch targetTypeOf {
	case reflect.String:
		targetValueOf.Elem().SetString(self.currentValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value, err := self.Int64()
		if err != nil {
			return self.err
		}
		if targetValueOf.Elem().OverflowInt(value) {
			self.err = wrapError(self.name, "Assignment failure", customError...)
			return self.err
		}
		targetValueOf.Elem().SetInt(value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value, err := self.Uint64()
		if err != nil {
			return self.err
		}
		if targetValueOf.Elem().OverflowUint(value) {
			self.err = wrapError(self.name, "Assignment failure", customError...)
			return self.err
		}
		targetValueOf.Elem().SetUint(value)
	case reflect.Float32, reflect.Float64:
		value, err := self.Float64()
		if err != nil {
			return self.err
		}
		if targetValueOf.Elem().OverflowFloat(value) {
			self.err = wrapError(self.name, "Assignment failure", customError...)
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
		sliceType := targetValueOf.Elem().Type().String()
		switch sliceType {
		case "[]string":
			value, err := self.SliceString()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]int":
			value, err := self.SliceInt()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]int8":
			value, err := self.SliceInt8()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]int16":
			value, err := self.SliceInt16()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]int32":
			value, err := self.SliceInt32()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]int64":
			value, err := self.SliceInt64()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]uint":
			value, err := self.SliceUint()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]uint8":
			value, err := self.SliceUint8()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]uint16":
			value, err := self.SliceUint16()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]uint32":
			value, err := self.SliceUint32()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]uint64":
			value, err := self.SliceUint64()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]float32":
			value, err := self.SliceFloat32()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]float64":
			value, err := self.SliceFloat64()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		case "[]bool":
			value, err := self.SliceBool()
			if err != nil {
				return self.err
			}
			targetValueOf.Elem().Set(reflect.ValueOf(value))
		default:
			self.err = wrapError(self.name, "Cannot set "+sliceType+" type of target")
			return self.err
		}
	default:
		self.err = wrapError(self.name, "Cannot set "+targetTypeOf.String()+" type of target")
		return self.err
	}
	return nil
}
