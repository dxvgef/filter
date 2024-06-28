package filter

import (
	"math"
	"testing"
)

// 测试 FromInteger().Error()
func TestFromIntegerError(t *testing.T) {
	err := FromInteger(math.MaxInt, "id").Require().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromInteger().Result()
func TestFromIntegerResult(t *testing.T) {
	value, err := FromInteger(math.MaxInt, "id").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromInteger().Int8()
func TestFromIntegerInt8(t *testing.T) {
	value, err := FromInteger(math.MaxInt8, "id").Int8()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromInteger().Int16()
func TestFromIntegerInt16(t *testing.T) {
	value, err := FromInteger(math.MaxInt16, "id").Int16()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromInteger().Int32()
func TestFromIntegerInt32(t *testing.T) {
	value, err := FromInteger(math.MaxInt32, "id").Int32()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromInteger().Int64()
func TestFromIntegerInt64(t *testing.T) {
	value := FromInteger(math.MaxInt64, "id").Int64()
	t.Log(value)
}

// 测试 FromInteger().DefaultInt8()
func TestFromIntegerDefaultInt8(t *testing.T) {
	value := FromInteger(0, "id").Require().DefaultInt8(1)
	t.Log(value)
}

// 测试 FromInteger().DefaultInt16()
func TestFromIntegerDefaultInt16(t *testing.T) {
	value := FromInteger(0, "id").Require().DefaultInt16(1)
	t.Log(value)
}

// 测试 FromInteger().DefaultInt32()
func TestFromIntegerDefaultInt32(t *testing.T) {
	value := FromInteger(0, "id").Require().DefaultInt32(1)
	t.Log(value)
}

// 测试 FromInteger().DefaultInt64() 输出 int64
func TestFromIntegerDefaultInt64(t *testing.T) {
	value := FromInteger(0, "id").Require().DefaultInt64(1)
	t.Log(value)
}
