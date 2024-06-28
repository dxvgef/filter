package filter

import (
	"testing"
)

// 测试 FromInteger().Min()
func TestFromIntegerMin(t *testing.T) {
	err := FromInteger(5, "id").Min(5).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromInteger().Max()
func TestFromIntegerMax(t *testing.T) {
	err := FromInteger(5, "id").Max(5).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromInteger().Equal()
func TestFromIntegerEqual(t *testing.T) {
	err := FromInteger(5, "id").Equal(5).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromInteger().NotEqual()
func TestFromIntegerNotEqual(t *testing.T) {
	err := FromInteger(6, "id").NotEqual(5).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromInteger().AllowedValues()
func TestFromIntegerAllowedValues(t *testing.T) {
	err := FromInteger(2, "id").AllowedValues([]int64{1, 2, 3}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromInteger().DeniedValues()
func TestFromIntegerDeniedValues(t *testing.T) {
	err := FromInteger(4, "id").DeniedValues([]int64{1, 2, 3}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}
