package filter

import "testing"

// 检查过滤结果
func TestCheck(t *testing.T) {
	err := InputString("输入数据", Config{Name: "数据名称"}).
		TrimSpace().
		MinLength(3, "不能小于3个字符").MaxLength(15).
		Error()
	if err != nil {
		t.Error(err)
	}
}

// 将过滤结果赋值到变量
func TestSet(t *testing.T) {
	var str string
	err := InputString("dxvgef", Config{Name: "数据名称"}).
		TrimSpace().
		MinLength(3, "不能小于3个字符").MaxLength(15).
		Set(&str)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(str)
}

// 类型转换
func TestConvert(t *testing.T) {
	i, err := InputString("123456", Config{Name: "字符串"}).
		TrimSpace().
		MinLength(3, "不能小于3个字符").MaxLength(15).
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
		InputString("dxvgef", Config{Name: "字符串1"}).MinLength(3, "不能小于3个字符").MaxLength(15).Error(),
		InputString("tsing", Config{Name: "字符串2"}).MinLength(3, "不能小于3个字符").MaxLength(15).Error(),
	)
	if err != nil {
		t.Error(err)
	}
}

// 批量赋值并检查处理结果
func TestBatchSet(t *testing.T) {
	var str1, str2 string
	err := Batch(
		InputString("dxvgef", Config{Name: "字符串1"}).TrimSpace().MinLength(3, "不能小于3个字符").MaxLength(15).Set(&str1),
		InputString("tsing", Config{Name: "字符串2"}).TrimSpace().MinLength(3, "不能小于3个字符").MaxLength(15).Set(&str2),
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(str1, str2)
}
