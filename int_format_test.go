package filter

import (
	"testing"
)

// 测试 FromInteger().Replace()
func TestFromIntegerReplace(t *testing.T) {
	t.Log(FromInteger(5).Replace(5, 1).Result())
	t.Log(FromInteger(1).Replace(3, 10).Result())
}
