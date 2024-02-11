package filter

import (
	"errors"
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
func wrapError(name, err string, custom ...string) error {
	var body strings.Builder
	body.WriteString(name)
	body.WriteString(": ")
	switch {
	case len(custom) > 0 && custom[0] != "":
		body.WriteString(custom[0])
	case err != "":
		body.WriteString(err)
	default:
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
