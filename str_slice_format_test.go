package filter

import "testing"

// 测试 FromStringSlice().Trim()
func TestStrSliceTrim(t *testing.T) {
	value, err := FromStringSlice([]string{"  aa  ", "  bb  "}).Trim(" ").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().TrimSpace()
func TestStrSliceTrimSpace(t *testing.T) {
	value, err := FromStringSlice([]string{"  aa  ", "  bb  "}).TrimSpace().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().TrimLeft()
func TestStrSliceTrimLeft(t *testing.T) {
	value, err := FromStringSlice([]string{"  aa  ", "  bb  "}).TrimLeft(" ").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().TrimRight()
func TestStrSliceTrimRight(t *testing.T) {
	value, err := FromStringSlice([]string{"  aa  ", "  bb  "}).TrimRight(" ").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().TrimPrefix()
func TestStrSliceTrimPrefix(t *testing.T) {
	value, err := FromStringSlice([]string{"  aa  ", "  bb  "}).TrimPrefix(" ").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().TrimSuffix()
func TestStrSliceTrimSuffix(t *testing.T) {
	value, err := FromStringSlice([]string{"  aa  ", "  bb  "}).TrimSuffix(" ").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().DeleteEmpty()
func TestStrSliceDeleteEmpty(t *testing.T) {
	value, err := FromStringSlice([]string{"  aa  ", "    ", "cc"}).TrimSpace().DeleteEmpty().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().RemoveSpace()
func TestStrSliceRemoveSpace(t *testing.T) {
	value, err := FromStringSlice([]string{"  aa  ", "     b      b          "}).RemoveSpace().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().Base64StdEncode(()
func TestStrSliceBase64StdEncode(t *testing.T) {
	value, err := FromStringSlice([]string{"  aa  ", "     b      b          "}).RemoveSpace().Base64StdEncode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().Base64StdDecode()
func TestStrSliceBase64StdDecode(t *testing.T) {
	value, err := FromStringSlice([]string{"YWE=", "YmI="}).Base64StdDecode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().Base64RawStdEncode()
func TestStrSliceBase64RawStdEncode(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).Base64RawStdEncode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().Base64RawStdDecode()
func TestStrSliceBase64RawStdDecode(t *testing.T) {
	value, err := FromStringSlice([]string{"YWE", "YmI"}).Base64RawStdDecode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().Base64URLEncode()
func TestStrSliceBase64URLEncode(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).Base64URLEncode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().Base64URLDecode()
func TestStrSliceBase64URLDecode(t *testing.T) {
	value, err := FromStringSlice([]string{"YWE=", "YmI="}).Base64URLDecode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().Base64URLEncode()
func TestStrSliceBase64RawURLEncode(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).Base64RawURLEncode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().Base64URLDecode()
func TestStrSliceBase64RawURLDecode(t *testing.T) {
	value, err := FromStringSlice([]string{"YWE", "YmI"}).Base64RawURLDecode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().HTMLEscape()
func TestStrSliceHTMLEscape(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).HTMLEscape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().HTMLUnescape()
func TestStrSliceHTMLUnescape(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).HTMLUnescape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().URLPathEscape()
func TestStrSliceURLPathEscape(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).URLPathEscape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().URLPathUnescape()
func TestStrSliceURLPathUnescape(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).URLPathUnescape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().URLQueryEscape()
func TestStrSliceURLQueryEscape(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).URLQueryEscape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromStringSlice().URLQueryUnescape()
func TestStrSliceURLQueryUnescape(t *testing.T) {
	value, err := FromStringSlice([]string{"aa", "bb"}).URLQueryUnescape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}
