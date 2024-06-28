package filter

import "strconv"

// ToInteger 转为int64对象
func (strType *StringType) ToInteger() *IntegerType {
	intType := &IntegerType{
		name: strType.name,
		err:  strType.err,
	}
	if strType.err != nil {
		return intType
	}
	intType.value, intType.err = strconv.ParseInt(strType.value, 10, 64)
	return intType
}
