package filter

import (
	"errors"
	"reflect"
	"strings"
	"unicode"
	"unsafe"
)

// InvalidErrorText 无效数据错误提示
const InvalidErrorText = "invalid data"

// Batch 批量处理
func Batch(errs ...error) error {
	for k := range errs {
		if errs[k] != nil {
			return errs[k]
		}
	}
	return nil
}

// 封装错误信息
func wrapError(name string, customError ...string) error {
	var body strings.Builder
	body.WriteString(name)
	if name != "" {
		body.WriteString(": ")
	}
	if len(customError) > 0 {
		body.WriteString(customError[0])
	} else {
		body.WriteString(InvalidErrorText)
	}
	return errors.New(body.String())
}

// []byte转string
func bytesToStr(value []byte) string {
	return *(*string)(unsafe.Pointer(&value))
}

// string转[]byte
func strToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 检查对象能否赋值
func setCheck(target interface{}) (*reflect.Value, error) {
	// 检查目标是否为nil
	if target == nil {
		return nil, errors.New(InvalidErrorText)
	}

	targetValueOf := reflect.ValueOf(target)
	// 检查对象是否是指针
	if targetValueOf.Kind() != reflect.Ptr {
		return nil, errors.New(InvalidErrorText)
	}
	// 检查对象是否能赋值
	if !targetValueOf.Elem().CanSet() {
		return nil, errors.New(InvalidErrorText)
	}
	return &targetValueOf, nil
}

func isSQLObject(value string) bool {
	for _, v := range value {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '_' {
			return false
		}
	}
	return true
}

// 用于校验uuid
var xvalues = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// 用于校验uuid
func xtob(x1, x2 byte) (byte, bool) {
	b1 := xvalues[x1]
	b2 := xvalues[x2]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}

func isUUID(str string) bool {
	var uuid [16]byte
	switch len(str) {
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36:

	// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36 + 9:
		if strings.EqualFold(strings.ToLower(str[:9]), "urn:uuid:") {
			return false
		}
		str = str[9:]
	// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
	case 36 + 2:
		str = str[1:]
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	case 32:
		var ok bool
		for i := range uuid {
			uuid[i], ok = xtob(str[i*2], str[i*2+1])
			if !ok {
				return false
			}
		}
		return true
	default:
		return false
	}
	// s is now at least 36 bytes long
	// it must be of the form  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		return false
	}
	for i, x := range [16]int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
		v, ok := xtob(str[x], str[x+1])
		if !ok {
			return false
		}
		uuid[i] = v
	}
	return true
}
