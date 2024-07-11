package filter

import (
	"testing"
)

// 测试 FromBooleanSlice().Result()
func TestFromBooleanSliceResult(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, true}, "test").Result())
}

// 测试 FromBooleanSlice().Value()
func TestFromBooleanSliceValue(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, true}).Value())
}

// 测试 FromBooleanSlice().Error()
func TestFromBooleanSliceError(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, true}).Error())
}

// 测试 FromBooleanSlice().DefaultValue()
func TestFromBooleanSliceDefaultValue(t *testing.T) {
	t.Log(FromBooleanSlice([]bool{false, false}).Has(true).DefaultValue([]bool{false, true}))
}
