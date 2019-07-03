package filter

import (
	"encoding/json"
	"errors"
	"net"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stringObject struct {
	err      error  // 结果错识信息
	name     string // 变量名
	rawValue string // 原始类型的值
	int      int    // int类型的值

}

func Value(value string, name ...string) *stringObject {
	var obj stringObject
	obj.rawValue = value

	if len(name) > 0 {
		obj.name = name[0]
	}
	return &obj
}

func (obj *stringObject) setError(msg string, customMessage ...string) error {
	if len(customMessage) > 0 {
		return errors.New(customMessage[0])
	}
	return errors.New(obj.name + msg)
}

func (obj *stringObject) Trim() *stringObject {
	obj.rawValue = strings.TrimSpace(obj.rawValue)
	return obj
}
func (obj *stringObject) ReplaceAll(old, new string) *stringObject {
	count := strings.Count(obj.rawValue, old)
	for i := 0; i < count; i++ {
		obj.rawValue = strings.ReplaceAll(obj.rawValue, old, new)
	}
	return obj
}

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

func (obj *stringObject) MinLength(min int, customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}
	if len(obj.rawValue) < min {
		obj.err = obj.setError("长度不能小于"+strconv.Itoa(min)+"个字符", customError...)
	}
	return obj
}

func (obj *stringObject) MinUTF8Length(min int, customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}
	if utf8.RuneCountInString(obj.rawValue) < min {
		obj.err = obj.setError("长度不能小于"+strconv.Itoa(min)+"个字符", customError...)
	}
	return obj
}

func (obj *stringObject) MaxLength(max int, customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}
	if len(obj.rawValue) > max {
		obj.err = obj.setError("长度不能大于"+strconv.Itoa(max)+"个字符", customError...)
	}
	return obj
}
func (obj *stringObject) MaxUTF8Length(max int, customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}
	if utf8.RuneCountInString(obj.rawValue) > max {
		obj.err = obj.setError("长度不能大于"+strconv.Itoa(max)+"个字符", customError...)
	}
	return obj
}

func (obj *stringObject) IsLower(customError ...string) *stringObject {
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

func (obj *stringObject) IsUpper(customError ...string) *stringObject {
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

func (obj *stringObject) IsLetter(customError ...string) *stringObject {
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

func (obj *stringObject) IsLowerAndDigit(customError ...string) *stringObject {
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

func (obj *stringObject) IsUpperAndDigit(customError ...string) *stringObject {
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

func (obj *stringObject) IsLetterAndDigit(customError ...string) *stringObject {
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

func (obj *stringObject) IsChinese(customError ...string) *stringObject {
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

func (obj *stringObject) IsChineseTel(customError ...string) *stringObject {
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

func (obj *stringObject) IsChineseMobile(customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}
	v, err := strconv.Atoi(obj.rawValue)
	if err != nil {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	obj.int = v
	if len(obj.rawValue) != 11 {
		obj.err = obj.setError("格式不正确", customError...)
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
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	suffix := obj.rawValue[3:]
	if suffix == "00000000" {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	return obj
}

func (obj *stringObject) IsChineseIdNumber(customError ...string) *stringObject {
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

func (obj *stringObject) IsMail(customError ...string) *stringObject {
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

func (obj *stringObject) IsIP(customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}
	if net.ParseIP(obj.rawValue) == nil {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}

	return obj
}

func xtob(x1, x2 byte) (byte, bool) {
	b1 := xvalues[x1]
	b2 := xvalues[x2]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}

func (obj *stringObject) IsUUID(customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}

	str := obj.rawValue

	var uuid [16]byte
	switch len(str) {
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36:

	// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36 + 9:
		if strings.ToLower(str[:9]) != "urn:uuid:" {
			obj.err = obj.setError("格式不正确", customError...)
			return obj
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
				obj.err = obj.setError("格式不正确", customError...)
				return obj
			}
		}
		return obj
	default:
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	// s is now at least 36 bytes long
	// it must be of the form  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		obj.err = obj.setError("格式不正确", customError...)
		return obj
	}
	for i, x := range [16]int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
		v, ok := xtob(str[x], str[x+1])
		if !ok {
			obj.err = obj.setError("格式不正确", customError...)
			return obj
		}
		uuid[i] = v
	}

	return obj
}

func (obj *stringObject) IsJSON(customError ...string) *stringObject {
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

func (obj *stringObject) HasLower(customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsLower(v) == true {
			return obj
		}
	}

	obj.err = obj.setError("格式不正确，必须存在小写字母", customError...)
	return obj
}

func (obj *stringObject) HasUpper(customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsUpper(v) == true {
			return obj
		}
	}

	obj.err = obj.setError("格式不正确，必须存在大写字母", customError...)
	return obj
}

func (obj *stringObject) HasDigit(customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsDigit(v) == true {
			return obj
		}
	}

	obj.err = obj.setError("格式不正确，必须存在大写字母", customError...)
	return obj
}

func (obj *stringObject) HasSymbol(customError ...string) *stringObject {
	if obj.err != nil {
		return obj
	}

	for _, v := range obj.rawValue {
		if unicode.IsDigit(v) == false || unicode.IsLetter(v) == false || unicode.Is(unicode.Han, v) == false {
			return obj
		}
	}

	obj.err = obj.setError("格式不正确，必须存在符号", customError...)
	return obj
}

func (obj *stringObject) String() (string, error) {
	if obj.err != nil {
		return "", obj.err
	}
	return obj.rawValue, nil
}
func (obj *stringObject) MustString(def string) string {
	if obj.err != nil {
		return def
	}
	return obj.rawValue
}

func (obj *stringObject) Int() (int, error) {
	if obj.err != nil {
		return 0, obj.err
	}
	if obj.int != 0 {
		return obj.int, nil
	}
	var err error
	obj.int, err = strconv.Atoi(obj.rawValue)
	if err != nil {
		return 0, err
	}
	return obj.int, nil
}

func (obj *stringObject) MustInt(def int) int {
	if obj.err != nil {
		return def
	}
	if obj.int != 0 {
		return obj.int
	}
	var err error
	obj.int, err = strconv.Atoi(obj.rawValue)
	if err != nil {
		return def
	}
	return obj.int
}

func (obj *stringObject) Int32() (int, error) {
	if obj.err != nil {
		return 0, obj.err
	}
	if obj.int != 0 {
		return obj.int, nil
	}
	var err error
	obj.int, err = strconv.Atoi(obj.rawValue)
	if err != nil {
		return 0, err
	}
	return obj.int, nil
}

func (obj *stringObject) MustInt(def int) int {
	if obj.err != nil {
		return def
	}
	if obj.int != 0 {
		return obj.int
	}
	var err error
	obj.int, err = strconv.Atoi(obj.rawValue)
	if err != nil {
		return def
	}
	return obj.int
}
