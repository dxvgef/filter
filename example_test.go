package filter

import "testing"

// 检查过滤结果
func TestCheck(t *testing.T) {
	err := String("dxvgef", Config{Name: "账号", Require: true}).
		TrimSpace().
		MinLength(3, "不能少于3个字符").MaxLength(15).
		Error()
	if err != nil {
		t.Error(err)
	}
}

// 将过滤结果赋值到变量
func TestSet(t *testing.T) {
	var str string
	err := String("dxvgef", Config{Name: "账号", Require: true}).
		TrimSpace().
		MinLength(3, "不能少于3个字符").MaxLength(15).
		Set(&str)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(str)
}

// 类型转换
func TestConvert(t *testing.T) {
	i, err := String("18", Config{Name: "年龄", Require: true}).
		TrimSpace().
		MinLength(18, "不能小于18岁").
		Int()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(i)
}

// 批量处理检查结果
func TestBatchCheck(t *testing.T) {
	err := Batch(
		String("dxvgef", Config{Name: "账号", Require: true}).
			MinLength(3, "不能少于3个字符").MaxLength(15).Error(),
		String("123456", Config{Name: "密码", Require: true}).
			MinLength(6, "不能少于6个字符").MaxLength(15).Error(),
	)
	if err != nil {
		t.Error(err)
	}
}

// 批量赋值并检查处理结果
func TestBatchSet(t *testing.T) {
	var str1, str2 string
	err := Batch(
		String("dxvgef", Config{Name: "账号", Require: true}).
			MinLength(3, "不能少于3个字符").MaxLength(15).
			Set(&str1),
		String("123456", Config{Name: "密码", Require: true}).
			MinLength(6, "不能少于6个字符").MaxLength(15).
			Set(&str2),
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(str1, str2)
}
