package filter

import "strconv"

// nolint:gocyclo
// IsChineseMobile 中国大陆地区手机号码
func (obj *Object) IsChineseMobile(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	_, err := strconv.Atoi(obj.rawValue)
	if err != nil {
		obj.err = obj.setError(customError...)
		return obj
	}
	if len(obj.rawValue) != 11 {
		obj.err = obj.setError(customError...)
		return obj
	}
	prefix := obj.rawValue[0:3]
	// 移动
	if prefix != "139" &&
		prefix != "138" &&
		prefix != "137" &&
		prefix != "136" &&
		prefix != "135" &&
		prefix != "134" &&
		prefix != "147" &&
		prefix != "150" &&
		prefix != "151" &&
		prefix != "152" &&
		prefix != "157" &&
		prefix != "158" &&
		prefix != "159" &&
		prefix != "165" &&
		prefix != "178" &&
		prefix != "182" &&
		prefix != "183" &&
		prefix != "184" &&
		prefix != "187" &&
		prefix != "188" &&
		prefix != "198" &&
		// 联通
		prefix != "130" &&
		prefix != "131" &&
		prefix != "132" &&
		prefix != "155" &&
		prefix != "156" &&
		prefix != "166" &&
		prefix != "167" &&
		prefix != "185" &&
		prefix != "186" &&
		prefix != "145" &&
		prefix != "175" &&
		prefix != "176" &&
		// 电信
		prefix != "133" &&
		prefix != "153" &&
		prefix != "162" &&
		prefix != "177" &&
		prefix != "173" &&
		prefix != "180" &&
		prefix != "181" &&
		prefix != "189" &&
		prefix != "191" &&
		prefix != "199" &&
		// 虚拟运营商
		prefix != "170" && prefix != "171" {
		obj.err = obj.setError(customError...)
		return obj
	}
	suffix := obj.rawValue[3:]
	if suffix == "00000000" {
		obj.err = obj.setError(customError...)
		return obj
	}
	return obj
}
