package filter

import (
	"testing"
)

// 测试 FromBooleanSlice().Has()
func TestFromBooleanSliceHas(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, false}).Has(true).Result())
}

// 测试 FromBooleanSlice().Count()
func TestFromBooleanSliceCount(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, false}).Count(2).Result())
}

// 测试 FromBooleanSlice().MinCount()
func TestFromBooleanSliceMinCount(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, false}).MinCount(3).Result())
}

// 测试 FromBooleanSlice().MaxCount()
func TestFromBooleanSliceMaxCount(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, false}).MaxCount(2).Result())
}
