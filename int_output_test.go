package filter

import (
	"math"
	"testing"
)

// 测试 FromInteger().Error()
func TestFromIntegerError(t *testing.T) {
	t.Log(FromInteger(1))
}

// 测试 FromInteger().Result()
func TestFromIntegerResult(t *testing.T) {
	t.Log(FromInteger(1).Result())
}

// 测试 FromInteger().Int()
func TestFromIntegerInt(t *testing.T) {
	t.Log(FromInteger(1).Int())
}

// 测试 FromInteger().Int8()
func TestFromIntegerInt8(t *testing.T) {
	t.Log(FromInteger(1).Int8())
	t.Log(FromInteger(math.MaxInt8 + 1).Int8())
}

// 测试 FromInteger().Int16()
func TestFromIntegerInt16(t *testing.T) {
	t.Log(FromInteger(1).Int16())
	t.Log(FromInteger(math.MaxInt16 + 1).Int16())
}

// 测试 FromInteger().Int32()
func TestFromIntegerInt32(t *testing.T) {
	t.Log(FromInteger(1).Int32())
	t.Log(FromInteger(math.MaxInt32 + 1).Int32())
}

// 测试 FromInteger().Int64()
func TestFromIntegerInt64(t *testing.T) {
	t.Log(FromInteger(1).Int64())
}

// 测试 FromInteger().DefaultInt()
func TestFromIntegerDefaultInt(t *testing.T) {
	t.Log(FromInteger(1).DefaultInt(2))
	t.Log(FromInteger(0).Require().DefaultInt(1))
}

// 测试 FromInteger().DefaultInt8Slice()
func TestFromIntegerDefaultInt8(t *testing.T) {
	t.Log(FromInteger(1).DefaultInt8(2))
	t.Log(FromInteger(math.MaxInt8 + 1).DefaultInt8(1))
}

// 测试 FromInteger().DefaultInt16()
func TestFromIntegerDefaultInt16(t *testing.T) {
	t.Log(FromInteger(1).DefaultInt16(2))
	t.Log(FromInteger(math.MaxInt16 + 1).DefaultInt16(1))
}

// 测试 FromInteger().DefaultInt32()
func TestFromIntegerDefaultInt32(t *testing.T) {
	t.Log(FromInteger(1).DefaultInt32(2))
	t.Log(FromInteger(math.MaxInt32 + 1).DefaultInt32(1))
}

// 测试 FromInteger().DefaultInt64() 输出 int64
func TestFromIntegerDefaultInt64(t *testing.T) {
	t.Log(FromInteger(1).Require().DefaultInt64(2))
	t.Log(FromInteger(0).Require().DefaultInt64(1))
}

// 测试 FromInteger().Uint()
func TestFromIntegerUint(t *testing.T) {
	t.Log(FromInteger(1).Uint())
	t.Log(FromInteger(-1).Uint())
}

// 测试 FromInteger().Uint8()
func TestFromIntegerUint8(t *testing.T) {
	t.Log(FromInteger(1).Uint8())
	t.Log(FromInteger(-1).Uint8())
}

// 测试 FromInteger().Uint16()
func TestFromIntegerUint16(t *testing.T) {
	t.Log(FromInteger(1).Uint16())
	t.Log(FromInteger(-1).Uint16())
}

// 测试 FromInteger().Uint32()
func TestFromIntegerUint32(t *testing.T) {
	t.Log(FromInteger(1).Uint32())
	t.Log(FromInteger(-1).Uint32())
}

// 测试 FromInteger().Uint64()
func TestFromIntegerUint64(t *testing.T) {
	t.Log(FromInteger(1).Uint64())
	t.Log(FromInteger(-1).Uint64())
}

// 测试 FromInteger().DefaultUint()
func TestFromIntegerDefaultUint(t *testing.T) {
	t.Log(FromInteger(1).DefaultUint(2))
	t.Log(FromInteger(-1).DefaultUint(1))
}

// 测试 FromInteger().DefaultUint8()
func TestFromIntegerDefaultUint8(t *testing.T) {
	t.Log(FromInteger(1).DefaultUint8(2))
	t.Log(FromInteger(-1).DefaultUint8(1))
}

// 测试 FromInteger().DefaultUint16()
func TestFromIntegerDefaultUint16(t *testing.T) {
	t.Log(FromInteger(1).DefaultUint16(2))
	t.Log(FromInteger(-1).DefaultUint16(1))
}

// 测试 FromInteger().DefaultUint32()
func TestFromIntegerDefaultUint32(t *testing.T) {
	t.Log(FromInteger(1).DefaultUint32(2))
	t.Log(FromInteger(-1).DefaultUint32(1))
}

// 测试 FromInteger().DefaultUint64()
func TestFromIntegerDefaultUint64(t *testing.T) {
	t.Log(FromInteger(1).Require().DefaultUint64(2))
	t.Log(FromInteger(-1).DefaultUint64(1))
}

// 测试 FromInteger().Set()
func TestFromIntegerSet(t *testing.T) {
	var value64 int64
	var value32 int32
	var ui8 uint8
	t.Log(FromInteger(math.MaxInt64).Set(&value64), value64)
	t.Log(FromInteger(math.MaxInt64).Set(&value32), value32)
	t.Log(FromInteger(math.MaxInt32).Set(&value32), value32)
	t.Log(FromInteger(math.MaxInt32).Set(&ui8), ui8)
	t.Log(FromInteger(math.MaxUint8).Set(&ui8), ui8)
}
