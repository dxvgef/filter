package rule

// Rule 规则
type Rule struct {
	Validate Validator // 验证器
	Trim     Trimmer   // 修剪器
	Message  string    // 自定义失败消息
	Length   struct {  // 长度
		Min int // 最小长度
		Max int // 最大长度
	}
	Range struct { // 值范围
		Min float64 // 最小值
		Max int     // 最大值
	}
}

// 内置验证器变量名
var (
	Must           = NewValidator(must, "值不能为空")
	Email          = NewValidator(email, "不是有效的电邮地址")
	IP             = NewValidator(ip, "不是有效的IP地址")
	Lower          = NewValidator(lower, "只能是小写字母")
	Upper          = NewValidator(upper, "只能是大写字母")
	Letter         = NewValidator(letter, "只能是大小写字母")
	Int            = NewValidator(intRule, "只能是0或整数")
	Uint           = NewValidator(uintRule, "只能是0或正整数")
	Nint           = NewValidator(nint, "只能是负整数")
	Float          = NewValidator(floatRule, "只能是0或小数")
	Pfloat         = NewValidator(pfloat, "只能是0或正小数")
	Nfloat         = NewValidator(nfloat, "只能是负小数")
	LowerAndDigit  = NewValidator(lowerAndDigit, "只能是小写字母或数字")
	UpperAndDigit  = NewValidator(upperAndDigit, "只能是大写字母或数字")
	LetterAndDigit = NewValidator(letterAndDigit, "只能是字母或数字")
	Chinese        = NewValidator(chinese, "只能是汉字")
	ChinaTel       = NewValidator(chinaTel, "只能是中国大陆地区的电话号码")
	ChinaMobile    = NewValidator(chinaMobile, "只能是中国大陆地区的手机号码")
	ChinaIDCard    = NewValidator(chinaIDCard, "只能是中国大陆地区的身份证号码")
	JSON           = NewValidator(jsonRule, "不能有效的序列化成JSON")
)

// 内置修剪器变量名
var (
	Trim = NewTrimmer(trim) //  去除前后空格
)
