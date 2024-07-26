package filter

import (
	"math"
	"testing"
)

// 测试 FromIntegerSlice().Result()
func TestFromIntIntegerSliceResult(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{1, 2}).Result())
}

// 测试 FromIntegerSlice().Value()
func TestFromIntIntegerSliceValue(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{1, 2}).Value())
}

// 测试 FromIntegerSlice().Error()
func TestFromIntIntegerSliceError(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Require().Error())
	t.Log(FromIntegerSlice([]int64{0, 0}).Require().Error())
}

// 测试 FromIntegerSlice().Set()
func TestFromIntIntegerSliceSet(t *testing.T) {
	var i []int
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&i))
	t.Log(i)
	var i8 []int8
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&i8))
	t.Log(i8)
	var i16 []int16
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&i16))
	t.Log(i16)
	var i32 []int32
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&i32))
	t.Log(i32)
	var i64 []int64
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&i64))
	t.Log(i64)

	var ui []uint
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&ui))
	t.Log(ui)
	var ui8 []uint8
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&ui8))
	t.Log(ui8)
	var ui16 []int16
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&ui16))
	t.Log(ui16)
	var ui32 []uint32
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&ui32))
	t.Log(ui32)
	var ui64 []uint64
	t.Log(FromIntegerSlice([]int64{0, 1}).Set(&ui64))
	t.Log(ui64)
}

// 测试 FromIntegerSlice().Int8Slice()
func TestFromIntIntegerSliceInt8Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Int8Slice())
	t.Log(FromIntegerSlice([]int64{0, math.MaxInt8 + 1}).Int8Slice())
}

// 测试 FromIntegerSlice().DefaultInt8Slice()
func TestFromIntIntegerSliceDefaultInt8(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Require().DefaultInt8Slice([]int8{1, 2}))
	t.Log(FromIntegerSlice([]int64{}).Require().DefaultInt8Slice([]int8{0, 1}))
}

// 测试 FromIntegerSlice().Int16Slice()
func TestFromIntIntegerSliceInt16Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Int16Slice())
	t.Log(FromIntegerSlice([]int64{0, math.MaxInt16 + 1}).Int16Slice())
}

// 测试 FromIntegerSlice().DefaultInt16Slice()
func TestFromIntIntegerSliceDefaultInt16Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Require().DefaultInt16Slice([]int16{1, 2}))
	t.Log(FromIntegerSlice([]int64{}).Require().DefaultInt16Slice([]int16{0, 1}))
}

// 测试 FromIntegerSlice().Int32Slice()
func TestFromIntIntegerSliceInt32Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Int32Slice())
	t.Log(FromIntegerSlice([]int64{0, math.MaxInt32 + 1}).Int32Slice())
}

// 测试 FromIntegerSlice().DefaultInt32Slice()
func TestFromIntIntegerSliceDefaultInt32Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Require().DefaultInt32Slice([]int32{1, 2}))
	t.Log(FromIntegerSlice([]int64{}).Require().DefaultInt32Slice([]int32{0, 1}))
}

// 测试 FromIntegerSlice().Int64Slice()
func TestFromIntIntegerSliceInt64Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Int64Slice())
	t.Log(FromIntegerSlice([]int64{}).Require().Int64Slice())
}

// 测试 FromIntegerSlice().DefaultInt64Slice()
func TestFromIntIntegerSliceDefaultInt64Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Require().DefaultInt64Slice([]int64{1, 2}))
	t.Log(FromIntegerSlice([]int64{}).Require().DefaultInt64Slice([]int64{0, 1}))
}

// 测试 FromIntegerSlice().IntSlice()
func TestFromIntIntegerSliceIntSlice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).IntSlice())
	t.Log(FromIntegerSlice([]int64{}).Require().IntSlice())
}

// 测试 FromIntegerSlice().DefaultIntSlice()
func TestFromIntIntegerSliceDefaultIntSlice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Require().DefaultIntSlice([]int{1, 2}))
	t.Log(FromIntegerSlice([]int64{}).Require().DefaultIntSlice([]int{0, 1}))
}

// 测试 FromIntegerSlice().Uint8Slice()
func TestFromIntIntegerSliceUint8Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Uint8Slice())
	t.Log(FromIntegerSlice([]int64{0, -1}).Uint8Slice())
}

// 测试 FromIntegerSlice().DefaultUint8Slice()
func TestFromIntIntegerSliceDefaultUint8Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).DefaultUint8Slice([]uint8{1, 2}))
	t.Log(FromIntegerSlice([]int64{0, -1}).DefaultUint8Slice([]uint8{0, 1}))
}

// 测试 FromIntegerSlice().Uint16Slice()
func TestFromIntIntegerSliceUint16Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Uint16Slice())
	t.Log(FromIntegerSlice([]int64{0, -1}).Uint16Slice())
}

// 测试 FromIntegerSlice().DefaultUint16Slice()
func TestFromIntIntegerSliceDefaultUint16Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).DefaultUint16Slice([]uint16{1, 2}))
	t.Log(FromIntegerSlice([]int64{0, -1}).DefaultUint16Slice([]uint16{0, 1}))
}

// 测试 FromIntegerSlice().Uint32Slice()
func TestFromIntIntegerSliceUint32Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Uint32Slice())
	t.Log(FromIntegerSlice([]int64{0, -1}).Uint32Slice())
}

// 测试 FromIntegerSlice().DefaultUint32Slice()
func TestFromIntIntegerSliceDefaultUint32Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).DefaultUint32Slice([]uint32{1, 2}))
	t.Log(FromIntegerSlice([]int64{0, -1}).DefaultUint32Slice([]uint32{0, 1}))
}

// 测试 FromIntegerSlice().Uint64Slice()
func TestFromIntIntegerSliceUint64Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).Uint64Slice())
	t.Log(FromIntegerSlice([]int64{}).Require().Uint64Slice())
}

// 测试 FromIntegerSlice().DefaultUint64Slice()
func TestFromIntIntegerSliceDefaultUint64Slice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{0, 1}).DefaultUint64Slice([]uint64{1, 2}))
	t.Log(FromIntegerSlice([]int64{}).Require().DefaultUint64Slice([]uint64{0, 1}))
}

// 测试 FromIntegerSlice().UintSlice()
func TestFromIntIntegerSliceUintSlice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{-1, 1}).UintSlice())
	t.Log(FromIntegerSlice([]int64{1, 2}).UintSlice())
}

// 测试 FromIntegerSlice().DefaultUintSlice()
func TestFromIntIntegerSliceDefaultUintSlice(t *testing.T) {
	t.Log(FromIntegerSlice([]int64{}).DefaultUintSlice([]uint{1, 2}))
	t.Log(FromIntegerSlice([]int64{0, 1}).DefaultUintSlice([]uint{}))
}
