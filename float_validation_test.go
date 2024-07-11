package filter

import (
	"testing"
)

// 测试 FromFloat().Require()
func TestFromFloatRequire(t *testing.T) {
	t.Log(FromFloat(0).Require().Error())
	t.Log(FromFloat(1).Require().Error())
}

// 测试 FromFloat().MinCount()
func TestFromFloatMin(t *testing.T) {
	t.Log(FromFloat(3).MinValue(5).Error())
	t.Log(FromFloat(5).MinValue(5).Error())
}

// 测试 FromFloat().MaxCount()
func TestFromFloatMax(t *testing.T) {
	t.Log(FromFloat(5).MaxValue(5).Error())
	t.Log(FromFloat(7).MaxValue(5).Error())
}

// 测试 FromFloat().AllowedValues()
func TestFromFloatAllowedValues(t *testing.T) {
	t.Log(FromFloat(2).AllowedValues([]float64{1, 3}).Error())
	t.Log(FromFloat(2).AllowedValues([]float64{1, 2, 3}).Error())
}

// 测试 FromFloat().DeniedValues()
func TestFromFloatDeniedValues(t *testing.T) {
	t.Log(FromFloat(2).DeniedValues([]float64{1, 2, 3}).Error())
	t.Log(FromFloat(4).DeniedValues([]float64{1, 2, 3}).Error())
}
