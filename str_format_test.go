package filter

import "testing"

// 测试 FromString().ToUpper()
func TestStrToUpper(t *testing.T) {
	value, err := FromString("abc").ToUpper().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().ToLower()
func TestStrToLower(t *testing.T) {
	value, err := FromString("aBc").ToLower().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Trim()
func TestStrTrim(t *testing.T) {
	value, err := FromString("0x0x1110x0x").Trim("0x").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().TrimSpace()
func TestStrTrimSpace(t *testing.T) {
	value, err := FromString("   abc   ").TrimSpace().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().TrimLeft()
func TestStrTrimLeft(t *testing.T) {
	value, err := FromString("0x0xabc").TrimLeft("0x").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().TrimRight()
func TestStrTrimRight(t *testing.T) {
	value, err := FromString("abc0x0x").TrimRight("0x").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().TrimPrefix()
func TestStrTrimPrefix(t *testing.T) {
	value, err := FromString("0x0xabc").TrimPrefix("0x").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().TrimSuffix()
func TestStrTrimSuffix(t *testing.T) {
	value, err := FromString("abc0x0x").TrimSuffix("0x").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Replace()
func TestStrReplace(t *testing.T) {
	value, err := FromString("abc0x0x").Replace("0x", "", 1).Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().ReplaceAll()
func TestStrReplaceAll(t *testing.T) {
	value, err := FromString("abc0x0x").ReplaceAll("0x", "").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().RemoveSpace()
func TestStrRemoveSpace(t *testing.T) {
	value, err := FromString("   a  b  c    ").RemoveSpace().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Base64StdEncode()
func TestStrBase64StdEncode(t *testing.T) {
	value, err := FromString("abc").Base64StdEncode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Base64StdDecode()
func TestStrBase64StdDecode(t *testing.T) {
	value, err := FromString("YWJj").Base64StdDecode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Base64RawStdEncode()
func TestStrBase64RawStdEncode(t *testing.T) {
	value, err := FromString("abc").Base64RawStdEncode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Base64RawStdDecode()
func TestStrBase64RawStdDecode(t *testing.T) {
	value, err := FromString("YWJj").Base64RawStdDecode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Base64URLEncode()
func TestStrBase64URLEncode(t *testing.T) {
	value, err := FromString("abc").Base64URLEncode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Base64URLDecode()
func TestStrBase64URLDecode(t *testing.T) {
	value, err := FromString("YWJj").Base64URLDecode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Base64RawURLEncode()
func TestStrBase64RawURLEncode(t *testing.T) {
	value, err := FromString("abc").Base64RawURLEncode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().Base64RawURLDecode()
func TestStrBase64RawURLDecode(t *testing.T) {
	value, err := FromString("YWJj").Base64RawURLDecode().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().HTMLUnescape()
func TestStrHTMLEscape(t *testing.T) {
	value, err := FromString("abc").HTMLEscape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().HTMLUnescape()
func TestStrHTMLUnescape(t *testing.T) {
	value, err := FromString("abc").HTMLUnescape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().URLPathEscape()
func TestStrURLPathEscape(t *testing.T) {
	value, err := FromString("龙").URLPathEscape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().URLPathUnescape()
func TestStrURLPathUnescape(t *testing.T) {
	value, err := FromString("%E9%BE%99").URLPathUnescape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().URLQueryEscape()
func TestStrURLQueryEscape(t *testing.T) {
	value, err := FromString("龙").URLQueryEscape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}

// 测试 FromString().URLQueryUnescape()
func TestStrURLQueryUnescape(t *testing.T) {
	value, err := FromString("%E9%BE%99").URLQueryUnescape().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}
