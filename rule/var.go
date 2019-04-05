package rule

import "strconv"

// Rule 规则
type Rule struct {
	Validate       Validator       // 验证器
	Trim           Trimmer         // 修剪器
	Message        string          // 自定义失败消息
	LengthValidate LengthValidator // 长度验证器
	MinLength      int             // 最小长度
	MaxLength      int             // 最大长度
}

// 内置验证器变量名
var (
	Required            = NewValidator(required, "值不能为空")
	Email               = NewValidator(email, "不是有效的电邮地址")
	IP                  = NewValidator(ip, "不是有效的IP地址")
	Lower               = NewValidator(lower, "只能是小写字母")
	Upper               = NewValidator(upper, "只能是大写字母")
	Letter              = NewValidator(letter, "只能是大小写字母")
	Int                 = NewValidator(intRule, "只能是0或整数")
	Uint                = NewValidator(uintRule, "只能是0或正整数")
	Nint                = NewValidator(nint, "只能是负整数")
	Float               = NewValidator(floatRule, "只能是0或小数")
	Pfloat              = NewValidator(pfloat, "只能是0或正小数")
	Nfloat              = NewValidator(nfloat, "只能是负小数")
	LowerAndDigit       = NewValidator(lowerAndDigit, "只能是小写字母或数字")
	UpperAndDigit       = NewValidator(upperAndDigit, "只能是大写字母或数字")
	LetterAndDigit      = NewValidator(letterAndDigit, "只能是字母或数字")
	Chinese             = NewValidator(chinese, "只能是汉字")
	ChineseTel          = NewValidator(chineseTel, "只能是中国大陆地区的电话号码")
	ChineseMobile       = NewValidator(chineseMobile, "只能是中国大陆地区的手机号码")
	ChineseIdentityCard = NewValidator(chineseIdentityCard, "只能是中国大陆地区的身份证号码")
	JSON                = NewValidator(jsonRule, "不能有效的序列化成JSON")
)

// 内置修剪器变量名
var (
	Trim = NewTrimmer(trim) //  去除前后空格
)

// 内置特殊验证器
var (
	// 设置字符串最小长度
	MinLength = func(min int) Rule {
		return SetMinLength(min, minLength, "不能小于"+strconv.Itoa(min)+"个字符")
	}
	// 设置字符串最大长度
	MaxLength = func(max int) Rule {
		return SetMaxLength(max, maxLength, "不能大于"+strconv.Itoa(max)+"个字符")
	}
)
