package filter

import (
	"errors"
	"reflect"
	"strings"
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
