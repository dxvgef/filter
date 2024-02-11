package filter

import (
	"testing"
)

func customTestFunc(s *Str) (string, error) {
	return s.Value() + "def", nil
	// return s.Value() + "def", errors.New("错误信息")
}

func TestCustomFunc(t *testing.T) {
	str, err := FromStr("abc", "账号").Custom(customTestFunc).String()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(str)
}

// 检查过滤结果
func TestCheck(t *testing.T) {
	err := FromStr("   ", "账号").Require().
		TrimSpace().Require().MinLength(3, "不能少于3个字符").MaxLength(15).
		Error()
	if err != nil {
		t.Error(err)
	}
}

// 检查过滤结果
func TestIsMail(t *testing.T) {
	err := FromStr("dxvgef@a.com", "邮箱").Require().
		IsMail().
		Error()
	if err != nil {
		t.Error(err)
	}
}

// 检查错误提示
func TestError(t *testing.T) {
	err := FromStr("ab", "账号").MinLength(5).Error()
	if err != nil {
		t.Error(err)
	}
}

// 将过滤结果赋值到普通变量
func TestSet(t *testing.T) {
	var str string
	err := FromStr("dxvgef", "账号").Require().
		TrimSpace().MinLength(3, "不能少于3个字符").MaxLength(15).
		Set(&str)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(str)
}

// 将过滤结果赋值到切片变量
func TestSetSlice(t *testing.T) {
	var str []string
	err := FromStr("dxv,gef", "账号").Require().
		TrimSpace().MinLength(3, "不能少于3个字符").MaxLength(15).
		SetSlice(&str, ",")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(str)
}

// 类型转换
func TestConvert(t *testing.T) {
	i, err := FromStr("18", "年龄").Require().
		TrimSpace().MinLength(18, "不能小于18岁").
		Int()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(i)
}

// 封装过滤规则
func TestWrapRule(t *testing.T) {
	// 定义检查账号的规则
	var checkUsername = func(value string, name ...string) *Str {
		return FromStr(value, name...).Require().MinLength(3, "不能少于3个字符").MaxLength(15)
	}

	err := checkUsername("dxvgef", "账号1")
	if err != nil {
		t.Error(err.Error())
	}

	err = checkUsername("dx", "账号2")
	if err != nil {
		t.Error(err.Error())
	}
}

// 批量处理检查结果
func TestBatchCheck(t *testing.T) {
	err := Batch(
		FromStr("dxvgef", "账号").Require().MinLength(3, "不能少于3个字符").MaxLength(15).Error(),
		FromStr("123456", "密码").Require().MinLength(6, "不能少于6个字符").MaxLength(15).Error(),
	)
	if err != nil {
		t.Error(err)
	}
}

// 批量赋值并检查处理结果
func TestBatchSet(t *testing.T) {
	var str1, str2 string
	err := Batch(
		FromStr("dxvgef", "账号").Require().MinLength(3, "不能少于3个字符").MaxLength(15).Set(&str1),
		FromStr("123456", "密码").Require().MinLength(6, "不能少于6个字符").MaxLength(15).Set(&str2),
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(str1, str2)
}

// 测试require和slice组合检查
func TestSliceRequire(t *testing.T) {
	str, err := FromStr(",", "账号").Require("不能为空").SliceInt(",")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(str, len(str))
}

// 测试EnumSliceString
func TestEnumSliceString(t *testing.T) {
	allow := []string{"a", "b", "c"}
	if err := FromStr("a,c", "账号").Require("不能为空").EnumStrSlice(",", allow).Error(); err != nil {
		t.Error(err)
		return
	}
}
