package filter

import (
	"testing"
)

// 测试 FromBoolean.Result()
func TestFromBooleanResult(t *testing.T) {
	t.Log(FromBoolean(false, "test").Result())
}

// 测试 FromBoolean().Value()
func TestFromBooleanValue(t *testing.T) {
	t.Log(FromBoolean(false).Value())
}

// 测试 FromBoolean().DefaultValue()
func TestFromBooleanDefaultValue(t *testing.T) {
	t.Log(FromBoolean(true).Value())
}

// 测试 FromBoolean().Error()
func TestFromBooleanError(t *testing.T) {
	t.Log(FromBoolean(false).Equal(true).Error())
	t.Log(FromBoolean(true).Equal(true).Error())
}

// 测试 FromBoolean().Set()
func TestFromBooleanSet(t *testing.T) {
	value := false
	t.Log(FromBoolean(true).Set(&value))
	t.Log(value)
}
