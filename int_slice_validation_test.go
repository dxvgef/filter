package filter

import (
	"testing"
)

// 测试 FromIntegerSlice().Require()
func TestFromIntegerSliceRequire(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{}).Require().Error())
	t.Log(FromIntegerSlice([]int64{0, 0}).Require().Error())
	t.Log(FromIntegerSlice([]int64{0, 1}).Require().Error())
}

// 测试 FromIntegerSlice().MinCount()
func TestFromIntegerSliceMinCount(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{}).MinCount(1).Error())
	t.Log(FromIntegerSlice([]int64{0}).MinCount(1).Error())
	t.Log(FromIntegerSlice([]int64{0}).Require().MinCount(1).Error())
	t.Log(FromIntegerSlice([]int64{3, 5}).MinCount(2).Error())
}

// 测试 FromIntegerSlice().MaxCount()
func TestFromIntegerSliceMaxCount(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{}).MaxCount(1).Error())
	t.Log(FromIntegerSlice([]int64{0, 0}).MaxCount(1).Error())
	t.Log(FromIntegerSlice([]int64{3, 5}).MaxCount(2).Error())
}

// 测试 FromIntegerSlice().MinValue()
func TestFromIntegerSliceMinValue(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 3}).MinValue(2).Error())
	t.Log(FromIntegerSlice([]int64{3, 5}).MinValue(2).Error())
}

// 测试 FromIntegerSlice().MaxValue()
func TestFromIntegerSliceMaxValue(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{1, 5}).MaxValue(5).Error())
}

// 测试 FromIntegerSlice().AllowedValues()
func TestFromIntegerSliceAllowedValues(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{1, 3}).AllowedValues([]int64{1, 2, 3}).Error())
	t.Log(FromIntegerSlice([]int64{1, 2, 3}).AllowedValues([]int64{1, 3}).Error())
}

// 测试 FromIntegerSlice().DeniedValues()
func TestFromIntegerSliceDeniedValues(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{1, 2, 4}).DeniedValues([]int64{3, 5}).Error())
	t.Log(FromIntegerSlice([]int64{1, 2, 3, 4}).DeniedValues([]int64{3, 5}).Error())
}
