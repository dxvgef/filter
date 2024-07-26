package filter

import (
	"testing"
)

// 测试 FromFloatSlice().Require()
func TestFromFloatSliceRequire(t *testing.T) {
	t.Log(FromFloatSlice([]float64{}).Require().Error())
	t.Log(FromFloatSlice([]float64{0, 0}).Require().Error())
	t.Log(FromFloatSlice([]float64{0, 1}).Require().Error())
}

// 测试 FromFloatSlice().MinCount()
func TestFromFloatSliceMinCount(t *testing.T) {
	t.Log(FromFloatSlice([]float64{}).MinCount(1).Error())
	t.Log(FromFloatSlice([]float64{0}).MinCount(1).Error())
	t.Log(FromFloatSlice([]float64{0}).Require().MinCount(1).Error())
	t.Log(FromFloatSlice([]float64{3, 5}).MinCount(2).Error())
}

// 测试 FromFloatSlice().MaxCount()
func TestFromFloatSliceMaxCount(t *testing.T) {
	t.Log(FromFloatSlice([]float64{}).MaxCount(1).Error())
	t.Log(FromFloatSlice([]float64{0, 0}).MaxCount(1).Error())
	t.Log(FromFloatSlice([]float64{3, 5}).MaxCount(2).Error())
}

// 测试 FromFloatSlice().MinValue()
func TestFromFloatSliceMinValue(t *testing.T) {
	t.Log(FromFloatSlice([]float64{0, 3}).MinValue(2).Error())
	t.Log(FromFloatSlice([]float64{3, 5}).MinValue(2).Error())
}

// 测试 FromFloatSlice().MaxValue()
func TestFromFloatSliceMaxValue(t *testing.T) {
	t.Log(FromFloatSlice([]float64{1, 5}).MaxValue(5).Error())
}

// 测试 FromFloatSlice().AllowedValues()
func TestFromFloatSliceAllowedValues(t *testing.T) {
	t.Log(FromFloatSlice([]float64{1, 3}).AllowedValues([]float64{1, 2, 3}).Error())
	t.Log(FromFloatSlice([]float64{1, 2, 3}).AllowedValues([]float64{1, 3}).Error())
}

// 测试 FromFloatSlice().DeniedValues()
func TestFromFloatSliceDeniedValues(t *testing.T) {
	t.Log(FromFloatSlice([]float64{1, 2, 4}).DeniedValues([]float64{3, 5}).Error())
	t.Log(FromFloatSlice([]float64{1, 2, 3, 4}).DeniedValues([]float64{3, 5}).Error())
}
