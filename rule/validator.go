package rule

import (
	"encoding/json"
	"net"
	"strconv"
	"strings"
	"unicode"
)

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

// 不能为空
func required(value string) bool {
	if value == "" {
		return false
	}
	return true
}

// 电邮地址
func email(str string) bool {
	emailSlice := strings.Split(str, "@")
	if len(emailSlice) != 2 {
		return false
	}
	if emailSlice[0] == "" || emailSlice[1] == "" {
		return false
	}

	for k, v := range emailSlice[0] {
		if k == 0 && unicode.IsLetter(v) == false && unicode.IsDigit(v) == false {
			return false
		} else if unicode.IsLetter(v) == false && unicode.IsDigit(v) == false && v != '@' && v != '.' && v != '_' && v != '-' {
			return false
		}
	}

	domainSlice := strings.Split(emailSlice[1], ".")
	if len(domainSlice) < 2 {
		return false
	}
	domainSliceLen := len(domainSlice)
	for i := 0; i < domainSliceLen; i++ {
		for k, v := range domainSlice[i] {
			if i != domainSliceLen-1 && k == 0 && unicode.IsLetter(v) == false && unicode.IsDigit(v) == false {
				return false
			} else if unicode.IsLetter(v) == false && unicode.IsDigit(v) == false && v != '.' && v != '_' && v != '-' {
				return false
			} else if i == domainSliceLen-1 && unicode.IsLetter(v) == false {
				return false
			}
		}
	}

	return true
}

// ip地址
func ip(str string) bool {
	if net.ParseIP(str) == nil {
		return false
	}
	return true
}

// 小写字母
func lower(str string) bool {
	for _, v := range str {
		if unicode.IsLower(v) == false {
			return false
		}
	}
	return true
}

// 大写字母
func upper(str string) bool {
	for _, v := range str {
		if unicode.IsUpper(v) == false {
			return false
		}
	}
	return true
}

// 大小写字母
func letter(str string) bool {
	for _, v := range str {
		if unicode.IsLetter(v) == false {
			return false
		}
	}
	return true
}

// 整数(含0)
func intRule(str string) bool {
	_, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false
	}
	return true
}

// >=0的整数
func uintRule(str string) bool {
	_, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return false
	}
	return true
}

// 负整数
func nint(str string) bool {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false
	}
	if i >= 0 {
		return false
	}
	return true
}

// 小数(含0)
func floatRule(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	return true
}

// 正小数(含0)
func pfloat(str string) bool {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	if i < 0 {
		return false
	}
	return true
}

// 负小数
func nfloat(str string) bool {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	if i >= 0 {
		return false
	}
	return true
}

// 小写字母和数字
func lowerAndDigit(str string) bool {
	for _, v := range str {
		if unicode.IsLower(v) == false && unicode.IsDigit(v) == false {
			return false
		}
	}
	return true
}

// 大写字母和数字
func upperAndDigit(str string) bool {
	for _, v := range str {
		if unicode.IsUpper(v) == false && unicode.IsDigit(v) == false {
			return false
		}
	}
	return true
}

// 字母和数字
func letterAndDigit(str string) bool {
	for _, v := range str {
		if unicode.IsLetter(v) == false && unicode.IsDigit(v) == false {
			return false
		}
	}
	return true
}

// 中国大陆地区座机号码
func chineseTel(str string) bool {
	telSlice := strings.Split(str, "-")
	if len(telSlice) != 2 {
		return false
	}
	regionCode, err := strconv.Atoi(telSlice[0])
	if err != nil {
		return false
	}
	if regionCode < 10 || regionCode > 999 {
		return false
	}

	code, err := strconv.Atoi(telSlice[1])
	if err != nil {
		return false
	}
	if code < 1000000 || code > 99999999 {
		return false
	}
	return true
}

// 中国大陆地区手机号码
func chineseMobile(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if len(str) != 11 {
		return false
	}
	prefix := str[0:3]
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
		return false
	}
	suffix := str[3:]
	if suffix == "00000000" {
		return false
	}
	return true
}

// 中文
func chinese(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Scripts["Han"], v) == false {
			return false
		}
	}
	return true
}

// JSON类型
func jsonRule(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func xtob(x1, x2 byte) (byte, bool) {
	b1 := xvalues[x1]
	b2 := xvalues[x2]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}
func uuidRule(str string) bool {
	var uuid [16]byte
	switch len(str) {
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36:

	// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36 + 9:
		if strings.ToLower(str[:9]) != "urn:uuid:" {
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

// 中国大陆地区身份证号码
func chineseIdentityCard(str string) bool {
	var idV int
	if str[17:] == "X" {
		idV = 88
	} else {
		var err error
		if idV, err = strconv.Atoi(str[17:]); err != nil {
			return false
		}
	}

	var verify int
	id := str[:17]
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
	if temp == idV {
		return true
	}

	return false
}
