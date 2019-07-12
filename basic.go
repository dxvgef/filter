package filter

import (
	"encoding/json"
	"errors"
	"net"
	"strconv"
	"strings"
	"unicode"
)

type Object struct {
	err      error   // 结果错识信息
	name     string  // 变量名
	rawValue string  // 原始类型的值
	i64      int64   // int64类型的值
	f64      float64 // float64类型的值
}

func FromString(value string, name ...string) *Object {
	var obj Object
	obj.rawValue = value

	if len(name) > 0 {
		obj.name = name[0]
	}
	return &obj
}

func (obj *Object) setError(msg string, customMessage ...string) error {
	if len(customMessage) > 0 {
		return errors.New(customMessage[0])
	}
	return errors.New(obj.name + msg)
}

// IsLower 小写字母
func (obj *Object) IsLower(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	for _, v := range obj.rawValue {
		if unicode.IsLower(v) == false {
			obj.err = obj.setError("必须是小写字母", customError...)
			return obj
		}
	}
	return obj
}

// IsUpper 大写字母
func (obj *Object) IsUpper(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	for _, v := range obj.rawValue {
		if unicode.IsUpper(v) == false {
			obj.err = obj.setError("必须是大写字母", customError...)
			return obj
		}
	}
	return obj
}

// IsLetter 大小写字母
func (obj *Object) IsLetter(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	for _, v := range obj.rawValue {
		if unicode.IsLetter(v) == false {
			obj.err = obj.setError("必须是字母", customError...)
			return obj
		}
	}
	return obj
}

// IsDigit 数字
func (obj *Object) IsDigit(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	for _, v := range obj.rawValue {
		if unicode.IsDigit(v) == false {
			obj.err = obj.setError("必须是数字", customError...)
			return obj
		}
	}
	return obj
}

// IsLowerOrDigit 小写字母或数字
func (obj *Object) IsLowerOrDigit(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	for _, v := range obj.rawValue {
		if unicode.IsLower(v) == false && unicode.IsDigit(v) == false {
			obj.err = obj.setError("必须是小写字母或数字", customError...)
			return obj
		}
	}
	return obj
}

// IsUpperOrDigit 大写字母或数字
func (obj *Object) IsUpperOrDigit(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	for _, v := range obj.rawValue {
		if unicode.IsUpper(v) == false && unicode.IsDigit(v) == false {
			obj.err = obj.setError("必须是大写字母或数字", customError...)
			return obj
		}
	}
	return obj
}

// IsLetterOrDigit 字母或数字
func (obj *Object) IsLetterOrDigit(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	for _, v := range obj.rawValue {
		if unicode.IsLetter(v) == false && unicode.IsDigit(v) == false {
			obj.err = obj.setError("必须是字母或数字", customError...)
			return obj
		}
	}
	return obj
}

// IsChinese 汉字
func (obj *Object) IsChinese(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	for _, v := range obj.rawValue {
		if unicode.Is(unicode.Scripts["Han"], v) == false {
			obj.err = obj.setError("必须是汉字", customError...)
			return obj
		}
	}
	return obj
}

// IsChineseTel 中国大陆地区手机号码
func (obj *Object) IsChineseTel(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	telSlice := strings.Split(obj.rawValue, "-")
	if len(telSlice) != 2 {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	regionCode, err := strconv.Atoi(telSlice[0])
	if err != nil {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	if regionCode < 10 || regionCode > 999 {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}

	code, err := strconv.Atoi(telSlice[1])
	if err != nil {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	if code < 1000000 || code > 99999999 {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	return obj
}

// IsMail 电邮地址
func (obj *Object) IsMail(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	emailSlice := strings.Split(obj.rawValue, "@")
	if len(emailSlice) != 2 {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	if emailSlice[0] == "" || emailSlice[1] == "" {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}

	for k, v := range emailSlice[0] {
		if k == 0 && unicode.IsLetter(v) == false && unicode.IsDigit(v) == false {
			obj.err = obj.setError("格式不正确", customError...)
			return obj
		} else if unicode.IsLetter(v) == false && unicode.IsDigit(v) == false && v != '@' && v != '.' && v != '_' && v != '-' {
			obj.err = obj.setError("格式不正确", customError...)
			return obj
		}
	}

	domainSlice := strings.Split(emailSlice[1], ".")
	if len(domainSlice) < 2 {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	domainSliceLen := len(domainSlice)
	for i := 0; i < domainSliceLen; i++ {
		for k, v := range domainSlice[i] {
			if i != domainSliceLen-1 && k == 0 && unicode.IsLetter(v) == false && unicode.IsDigit(v) == false {
				obj.err = obj.setError("格式不正确", customError...)
				return obj
			} else if unicode.IsLetter(v) == false && unicode.IsDigit(v) == false && v != '.' && v != '_' && v != '-' {
				obj.err = obj.setError("格式不正确", customError...)
				return obj
			} else if i == domainSliceLen-1 && unicode.IsLetter(v) == false {
				obj.err = obj.setError("格式不正确", customError...)
				return obj
			}
		}
	}

	return obj
}

// IsIP IPv4/v6地址
func (obj *Object) IsIP(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	if net.ParseIP(obj.rawValue) == nil {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}

	return obj
}

// IsJSON 有效的JSON格式
func (obj *Object) IsJSON(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}

	var js json.RawMessage
	if json.Unmarshal([]byte(obj.rawValue), &js) != nil {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	return obj
}

// IsChineseIDNumber 中国大陆地区身份证号码
func (obj *Object) IsChineseIDNumber(customError ...string) *Object {
	if obj.err != nil {
		return obj
	}
	var idV int
	if obj.rawValue[17:] == "X" {
		idV = 88
	} else {
		var err error
		if idV, err = strconv.Atoi(obj.rawValue[17:]); err != nil {
			obj.err = obj.setError("格式不正确", customError...)
			return obj
		}
	}

	var verify int
	id := obj.rawValue[:17]
	arr := make([]int, 17)
	for i := 0; i < 17; i++ {
		arr[i], _ = strconv.Atoi(string(id[i]))
	}
	wi := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		res += arr[i] * wi[i]
	}
	verify = res % 11

	var temp int
	a18 := [11]int{1, 0, 88 /* 'X' */, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			break
		}
	}
	if temp != idV {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}

	return obj
}