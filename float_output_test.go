package filter

import (
	"testing"
)

// 测试 FromFloat().Error()
func TestFromFloatError(t *testing.T) {
	t.Log(FromFloat(3.14159265359))
}

// 测试 FromFloat().Result()
func TestFromFloatResult(t *testing.T) {
	t.Log(FromFloat(3.14159265359).Result())
}

// 测试 FromFloat().Set()
func TestFromFloatSet(t *testing.T) {
	var value32 float32
	var value64 float64
	t.Log(FromFloat(3.14159265359).Set(&value32))
	t.Log(FromFloat(3.14159265359).Set(&value64))
	t.Log(value32)
	t.Log(value64)
}

// 测试 FromFloat().Float32()
func TestFromFloatFloat32(t *testing.T) {
	t.Log(FromFloat(3.14159265359).Float32())
}

// 测试 FromFloat().DefaultFloat32()
func TestFromFloatDefaultFloat32(t *testing.T) {
	t.Log(FromFloat(0).DefaultFloat32(3.14159265359))
}

// 测试 FromFloat().Float64()
func TestFromFloatFloat64(t *testing.T) {
	t.Log(FromFloat(3.14159265359).Float64())
}

// 测试 FromFloat().DefaultFloat64()
func TestFromFloatDefaultFloat64(t *testing.T) {
	t.Log(FromFloat(0).DefaultFloat64(3.14159265359))
}
