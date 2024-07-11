package filter

import (
	"testing"
)

// 测试 FromInteger().Require()
func TestFromIntegerRequire(t *testing.T) {
	t.Log(FromInteger(0).Require().Error())
	t.Log(FromInteger(1).Require().Error())
}

// 测试 FromInteger().MinCount()
func TestFromIntegerMin(t *testing.T) {
	t.Log(FromInteger(3).MinValue(5).Error())
	t.Log(FromInteger(5).MinValue(5).Error())
}

// 测试 FromInteger().MaxCount()
func TestFromIntegerMax(t *testing.T) {
	t.Log(FromInteger(5).MaxValue(5).Error())
	t.Log(FromInteger(7).MaxValue(5).Error())
}

// 测试 FromInteger().AllowedValues()
func TestFromIntegerAllowedValues(t *testing.T) {
	t.Log(FromInteger(2).AllowedValues([]int64{1, 3}).Error())
	t.Log(FromInteger(2).AllowedValues([]int64{1, 2, 3}).Error())
}

// 测试 FromInteger().DeniedValues()
func TestFromIntegerDeniedValues(t *testing.T) {
	t.Log(FromInteger(2).DeniedValues([]int64{1, 2, 3}).Error())
	t.Log(FromInteger(4).DeniedValues([]int64{1, 2, 3}).Error())
}
