package filter

import (
	"testing"
)

// 测试 FromBoolean().Equal()
func TestBooleanEqual(t *testing.T) {
	t.Log(FromBoolean(false, "test").Equal(true).Error())
	t.Log(FromBoolean(false, "test").Equal(false).Error())
}

// 测试 FromBoolean().NotEqual()
func TestBooleanNotEqual(t *testing.T) {
	t.Log(FromBoolean(false, "test").NotEqual(true).Error())
	t.Log(FromBoolean(false, "test").NotEqual(false).Error())
}
