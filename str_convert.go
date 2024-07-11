package filter

import (
	"strconv"
	"strings"
)

// ToStringSlice 转为[]string类型
func (strType *StringType) ToStringSlice(sep string) *StringSliceType {
	strSliceType := &StringSliceType{
		name: strType.name,
		err:  strType.err,
	}
	if strType.err != nil {
		return strSliceType
	}
	strSliceType.value = strings.Split(strType.value, sep)
	return strSliceType
}

// ToInteger 转为int64类型
func (strType *StringType) ToInteger(customError ...string) *IntegerType {
	intType := &IntegerType{
		name: strType.name,
		err:  strType.err,
	}
	if strType.err != nil {
		return intType
	}
	v, err := strconv.ParseInt(strType.value, 10, 64)
	if err != nil {
		intType.err = wrapError(strType.name, customError...)
		return intType
	}
	intType.value = v
	return intType
}

// ToIntegerSlice 转为[]int64类型
func (strType *StringType) ToIntegerSlice(sep string, customError ...string) *IntegerSliceType {
	intSliceType := &IntegerSliceType{
		name: strType.name,
		err:  strType.err,
	}
	if strType.err != nil {
		return intSliceType
	}
	strSlice := strings.Split(strType.value, sep)
	for k := range strSlice {
		v, err := strconv.ParseInt(strSlice[k], 10, 64)
		if err != nil {
			intSliceType.err = wrapError(strType.name, customError...)
			return intSliceType
		}
		intSliceType.value = append(intSliceType.value, v)
	}
	return intSliceType
}

// ToBooleanType 转[]bool类型
func (strType *StringType) ToBooleanType(customError ...string) *BooleanType {
	boolType := &BooleanType{
		name: strType.name,
		err:  strType.err,
	}
	if strType.err != nil {
		return boolType
	}
	v, err := strconv.ParseBool(strType.value)
	if err != nil {
		boolType.err = wrapError(strType.name, customError...)
		return boolType
	}
	boolType.value = v
	return boolType
}

// ToBooleanSliceType 转为[]bool]类型
func (strType *StringType) ToBooleanSliceType(sep string, customError ...string) *BooleanSliceType {
	boolSliceType := &BooleanSliceType{
		name: strType.name,
		err:  strType.err,
	}
	if strType.err != nil {
		return boolSliceType
	}
	strSlice := strings.Split(strType.value, sep)
	for k := range strSlice {
		v, err := strconv.ParseBool(strSlice[k])
		if err != nil {
			boolSliceType.err = wrapError(strType.name, customError...)
			return boolSliceType
		}
		boolSliceType.value = append(boolSliceType.value, v)
	}
	return boolSliceType
}
