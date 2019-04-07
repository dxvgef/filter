package rule

import (
	"strconv"

	"github.com/dxvgef/filter"
)

// 内置验证器变量名
var (
	Required = filter.NewValidator(required, "值不能为空")
	Lower    = filter.NewValidator(lower, "只能是小写字母")
	Upper    = filter.NewValidator(upper, "只能是大写字母")
	Letter   = filter.NewValidator(letter, "只能是大小写字母")
	Int      = filter.NewValidator(intRule, "只能是0或整数")
	Uint     = filter.NewValidator(uintRule, "只能是0或正整数")
	Nint     = filter.NewValidator(nint, "只能是负整数")
	Float    = filter.NewValidator(floatRule, "只能是0或小数")
	Pfloat   = filter.NewValidator(pfloat, "只能是0或正小数")
	Nfloat   = filter.NewValidator(nfloat, "只能是负小数")
	JSON     = filter.NewValidator(jsonRule, "不是有效的JSON字符串")

	HasLower       = filter.NewValidator(hasLower, "必须含有至少一个小写字母")
	HasUpper       = filter.NewValidator(hasUpper, "必须含有至少一个大写字母")
	HasDigit       = filter.NewValidator(hasDigit, "必须含有至少一个数字")
	HasSpecialChar = filter.NewValidator(hasSpecialChar, "必须含有至少一个特殊字符")

	Email = filter.NewValidator(email, "不是有效的电邮地址")
	IP    = filter.NewValidator(ip, "不是有效的IP地址")
	UUID  = filter.NewValidator(uuidRule, "不是有效的UUID字符串")

	LowerAndDigit       = filter.NewValidator(lowerAndDigit, "只能是小写字母或数字")
	UpperAndDigit       = filter.NewValidator(upperAndDigit, "只能是大写字母或数字")
	LetterAndDigit      = filter.NewValidator(letterAndDigit, "只能是字母或数字")
	Chinese             = filter.NewValidator(chinese, "只能是汉字")
	ChineseTel          = filter.NewValidator(chineseTel, "只能是中国大陆地区的电话号码")
	ChineseMobile       = filter.NewValidator(chineseMobile, "只能是中国大陆地区的手机号码")
	ChineseIdentityCard = filter.NewValidator(chineseIdentityCard, "只能是中国大陆地区的身份证号码")
)

// 内置修剪器变量名
var (
	Trim                  = filter.NewTrimmer(trim)                  // 去除前后空格
	RemoveSpace           = filter.NewTrimmer(removeSpace)           // 去除所有的空格
	SnakeCaseToCamelCase  = filter.NewTrimmer(snakeCaseToCamelCase)  // 蛇形转驼峰
	SnakeCaseToPascalCase = filter.NewTrimmer(snakeCaseToPascalCase) // 蛇形转帕斯卡
	CamelCaseToSnakeCase  = filter.NewTrimmer(camelCaseToSnakeCase)  // 驼峰/帕斯卡转蛇形
)

// 内置特殊验证器
var (
	// 字符串最小长度
	MinLength = func(min int) filter.Rule {
		return filter.NewMinLengthValidator(min, minLength, "不能小于"+strconv.Itoa(min)+"个字符")
	}
	// 字符串最大长度
	MaxLength = func(max int) filter.Rule {
		return filter.NewMaxLengthValidator(max, maxLength, "不能大于"+strconv.Itoa(max)+"个字符")
	}
	// 最小整数值
	MinInteger = func(min int64) filter.Rule {
		return filter.NewMinIntegerValidator(min, minInteger, "数值不能小于"+strconv.FormatInt(min, 10))
	}
	// 最大整数值
	MaxInteger = func(max int64) filter.Rule {
		return filter.NewMaxIntegerValidator(max, maxInteger, "数值不能大于"+strconv.FormatInt(max, 10))
	}
	// 最小浮点数值
	MinFloat = func(min float64) filter.Rule {
		return filter.NewMinFloatValidator(min, minFloat, "数值不能小于"+strconv.FormatFloat(min, 'f', -1, 64))
	}
	// 最大整数值
	MaxFloat = func(max float64) filter.Rule {
		return filter.NewMaxFloatValidator(max, maxFloat, "数值不能大于"+strconv.FormatFloat(max, 'f', -1, 64))
	}
	// 值必须存在于[]string
	InStrings = func(slice []string) filter.Rule {
		return filter.NewInValidator(slice, inSlice, "不允许的值")
	}

	// 必须存在指定字符串
	Contains = func(sub string) filter.Rule {
		return filter.NewContainValidator(sub, contains, "必须包含字符串："+sub)
	}

	// 必须存在指定的前缀字符串
	HasPrefix = func(sub string) filter.Rule {
		return filter.NewHasPrefixValidator(sub, hasPrefix, "不允许的值")
	}
)
