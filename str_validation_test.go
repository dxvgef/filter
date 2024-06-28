package filter

import "testing"

// 测试 FromString().IsBool()
func TestStrIsBool(t *testing.T) {
	err := FromString("true").IsBool().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsLower()
func TestStrIsLower(t *testing.T) {
	err := FromString("abc").IsLower().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsUpper()
func TestStrIsUpper(t *testing.T) {
	err := FromString("ABC").IsUpper().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsLetter()
func TestStrIsLetter(t *testing.T) {
	err := FromString("abC").IsLetter().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsUnsigned()
func TestStrIsUnsigned(t *testing.T) {
	err := FromString("0").IsUnsigned().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsLowerOrNumber()
func TestStrIsLowerOrNumber(t *testing.T) {
	err := FromString("a0").IsLowerOrNumber().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsUpperOrNumber()
func TestStrIsUpperOrNumber(t *testing.T) {
	err := FromString("A0").IsUpperOrNumber().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsLetterOrNumber()
func TestStrIsLetterOrNumber(t *testing.T) {
	err := FromString("Ab0").IsLetterOrNumber().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsChinese()
func TestStrIsChinese(t *testing.T) {
	err := FromString("龙").IsChinese().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsMail()
func TestStrIsMail(t *testing.T) {
	err := FromString("a@b.c").IsMail().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsIP()
func TestStrIsIPv4(t *testing.T) {
	err := FromString("127.0.0.1").IsIP().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsIP()
func TestStrIsIPv6(t *testing.T) {
	err := FromString("2001:0db8:3c4d:0015:0000:0000:1a2f:1a2b").IsIP().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsTCPAddr()
func TestStrIsTCPAddr(t *testing.T) {
	err := FromString("localhost:80").IsTCPAddr().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsMAC()
func TestStrIsMAC(t *testing.T) {
	err := FromString("08:00:20:0A:8C:6D").IsMAC().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsJSON()
func TestStrIsJSON(t *testing.T) {
	err := FromString(`{"a": "0"}`).IsJSON().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsSQLObject()
func TestStrIsSQLObject(t *testing.T) {
	err := FromString("username").IsSQLObject().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsSQLObjects()
func TestStrIsSQLObjects(t *testing.T) {
	err := FromString("username,password").IsSQLObjects(",").Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsURL()
func TestStrIsURL(t *testing.T) {
	err := FromString("https://www.google.com").IsURL().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsUUID()
func TestStrIsUUID(t *testing.T) {
	err := FromString("123e4567-e89b-12d3-a456-426655440000").IsUUID().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().IsULID()
func TestStrIsULID(t *testing.T) {
	err := FromString("01J1C8KMEXV8JHMS0QDNCJTFSZ").IsULID().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().Length()
func TestStrLength(t *testing.T) {
	err := FromString("abc").Length(3).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().UTF8Length()
func TestStrUTF8Length(t *testing.T) {
	err := FromString("龙井").UTF8Length(2).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().MinLength()
func TestStrMinLength(t *testing.T) {
	err := FromString("abcdef").MinLength(3).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().UTF8MinLength()
func TestStrUTF8MinLength(t *testing.T) {
	err := FromString("龙井").UTF8MinLength(2).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().MaxLength()
func TestStrMaxLength(t *testing.T) {
	err := FromString("abdef").MaxLength(5).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().UTF8MaxLength()
func TestStrUTF8MaxLength(t *testing.T) {
	err := FromString("龙a").UTF8MaxLength(2).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().AllowedChars()
func TestStrAllowedChars(t *testing.T) {
	err := FromString("aaac").AllowedChars([]rune{'a', 'b', 'c'}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().AllowedSymbols()
func TestStrAllowedSymbols(t *testing.T) {
	err := FromString("aa_ac").AllowedSymbols([]rune{'-', '_', '.'}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().AllowedStrings()
func TestStrAllowedStrings(t *testing.T) {
	err := FromString("female").AllowedStrings([]string{"male", "female", "other"}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().DeniedChars()
func TestStrDeniedChars(t *testing.T) {
	err := FromString("adefg").DeniedChars([]rune{'b', 'c'}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().DeniedSymbols()
func TestStrDeniedSymbols(t *testing.T) {
	err := FromString("abc").DeniedSymbols([]rune{'@', '.'}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().DeniedStrings()
func TestStrDeniedStrings(t *testing.T) {
	err := FromString("abc").DeniedStrings([]string{"male", "female"}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().HasLetter()
func TestStrHasLetter(t *testing.T) {
	err := FromString("B0123").HasLetter().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().HasLower()
func TestStrHasLower(t *testing.T) {
	err := FromString("a0123").HasLower().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().HasUpper()
func TestStrHasUpper(t *testing.T) {
	err := FromString("aB0123").HasUpper().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().HasNumber()
func TestStrHasNumber(t *testing.T) {
	err := FromString("0abc").HasNumber().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().HasSymbol()
func TestStrHasSymbol(t *testing.T) {
	err := FromString("0abc@").HasSymbol().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().Contains()
func TestStrContains(t *testing.T) {
	err := FromString("abcdefg").Contains("def").Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().HasPrefix()
func TestStrHasPrefix(t *testing.T) {
	err := FromString("abcdefg").HasPrefix("ab").Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromString().HasSuffix()
func TestStrHasSuffix(t *testing.T) {
	err := FromString("abcdefg").HasSuffix("fg").Error()
	if err != nil {
		t.Error(err)
		return
	}
}
