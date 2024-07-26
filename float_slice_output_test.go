package filter

import (
	"testing"
)

// 测试 FromFloatSlice().Result()
func TestFromFloatSliceResult(t *testing.T) {
	t.Log(FromFloatSlice([]float64{1, 2}).Result())
}

// 测试 FromFloatSlice().Value()
func TestFromFloatSliceValue(t *testing.T) {
	t.Log(FromFloatSlice([]float64{1, 2}).Value())
}

// 测试 FromFloatSlice().Error()
func TestFromFloatSliceError(t *testing.T) {
	t.Log(FromFloatSlice([]float64{0, 1}).Require().Error())
	t.Log(FromFloatSlice([]float64{0, 0}).Require().Error())
}

// 测试 FromFloatSlice().Set()
func TestFromFloatSliceSet(t *testing.T) {
	var value64 []float64
	var value32 []float32
	t.Log(FromFloatSlice([]float64{0, 1}).Set(&value64))
	t.Log(value64)
	t.Log(FromFloatSlice([]float64{0, 1}).Set(&value32))
	t.Log(value32)
}

// 测试 FromFloatSlice().Float32Slice()
func TestFromFloatSliceFloat32Slice(t *testing.T) {
	t.Log(FromFloatSlice([]float64{0, 1}).Float32Slice())
	t.Log(FromFloatSlice([]float64{0, 3.5e38}).Float32Slice())
	t.Log(FromFloatSlice([]float64{-3.5e38, 0}).Float32Slice())
}

// 测试 FromFloatSlice().DefaultFloat32Slice()
func TestFromFloatSliceDefaultFloat32Slice(t *testing.T) {
	t.Log(FromFloatSlice([]float64{0, 1}).Require().DefaultFloat32Slice([]float32{1, 2}))
	t.Log(FromFloatSlice([]float64{}).Require().DefaultFloat32Slice([]float32{0, 1}))
}

// 测试 FromFloatSlice().Float64Slice()
func TestFromFloatSliceInt16Slice(t *testing.T) {
	t.Log(FromFloatSlice([]float64{0, 1}).Float64Slice())
	t.Log(FromFloatSlice([]float64{0, 3.5e38}).Float64Slice())
	t.Log(FromFloatSlice([]float64{-3.5e38, 0}).Float64Slice())
}

// 测试 FromFloatSlice().DefaultFloat64Slice()
func TestFromFloatSliceDefaultFloat64Slice(t *testing.T) {
	t.Log(FromFloatSlice([]float64{0, 1}).Require().DefaultFloat64Slice([]float64{1, 2}))
	t.Log(FromFloatSlice([]float64{}).Require().DefaultFloat64Slice([]float64{0, 1}))
}
