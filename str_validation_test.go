package filter

import "testing"

// 测试 FromString().Require()
func TestFromStringRequire(t *testing.T) {
	t.Log(FromString("      ").TrimSpace().Require().Error())
	t.Log(FromString("aa").TrimSpace().Require().Error())
}

// 测试 FromString().Length()
func TestStrLength(t *testing.T) {
	t.Log(FromString("ab").Length(3).Error())
	t.Log(FromString("a龙").Length(3).Error())
	t.Log(FromString("abc").Length(3).Error())
}

// 测试 FromString().UTF8Length()
func TestStrUTF8Length(t *testing.T) {
	t.Log(FromString("a龙井").UTF8Length(2).Error())
	t.Log(FromString("ab").UTF8Length(2).Error())
	t.Log(FromString("a井").UTF8Length(2).Error())
	t.Log(FromString("龙井").UTF8Length(2).Error())
}

// 测试 FromString().MinLength()
func TestStrMinLength(t *testing.T) {
	t.Log(FromString("ab").MinLength(3).Error())
	t.Log(FromString("abcdef").MinLength(3).Error())
}

// 测试 FromString().UTF8MinLength()
func TestStrUTF8MinLength(t *testing.T) {
	t.Log(FromString("龙").UTF8MinLength(2).Error())
	t.Log(FromString("龙井").UTF8MinLength(2).Error())
}

// 测试 FromString().MaxLength()
func TestStrMaxLength(t *testing.T) {
	t.Log(FromString("abdef龙").MaxLength(5).Error())
	t.Log(FromString("龙井").MaxLength(5).Error())
	t.Log(FromString("abdef").MaxLength(5).Error())
}

// 测试 FromString().UTF8MaxLength()
func TestStrUTF8MaxLength(t *testing.T) {
	t.Log(FromString("a龙井").UTF8MaxLength(2).Error())
	t.Log(FromString("龙井").UTF8MaxLength(2).Error())
	t.Log(FromString("a龙").UTF8MaxLength(2).Error())
}

// 测试 FromString().IsLower()
func TestStrIsLower(t *testing.T) {
	t.Log(FromString("aBc").IsLower().Error())
	t.Log(FromString("abc").IsLower().Error())
}

// 测试 FromString().IsUpper()
func TestStrIsUpper(t *testing.T) {
	t.Log(FromString("AbC").IsUpper().Error())
	t.Log(FromString("ABC").IsUpper().Error())
}

// 测试 FromString().IsLetter()
func TestStrIsLetter(t *testing.T) {
	t.Log(FromString("ab0").IsLetter().Error())
	t.Log(FromString("abC").IsLetter().Error())
}

// 测试 FromString().IsLowerOrNumber()
func TestStrIsLowerOrNumber(t *testing.T) {
	t.Log(FromString("aB0").IsLowerOrNumber().Error())
	t.Log(FromString("a0").IsLowerOrNumber().Error())
}

// 测试 FromString().IsUpperOrNumber()
func TestStrIsUpperOrNumber(t *testing.T) {
	t.Log(FromString("Ab0").IsUpperOrNumber().Error())
	t.Log(FromString("AB0").IsUpperOrNumber().Error())
}

// 测试 FromString().IsLetterOrNumber()
func TestStrIsLetterOrNumber(t *testing.T) {
	t.Log(FromString("A-0").IsLetterOrNumber().Error())
	t.Log(FromString("Ab0").IsLetterOrNumber().Error())
}

// 测试 FromString().IsChinese()
func TestStrIsChinese(t *testing.T) {
	t.Log(FromString("a龙井").IsChinese().Error())
	t.Log(FromString("龙井").IsChinese().Error())
}

// 测试 FromString().IsMail()
func TestStrIsMail(t *testing.T) {
	t.Log(FromString("@localhost").IsMail().Error())
	t.Log(FromString("a@b.c").IsMail().Error())
}

// 测试 FromString().IsIP()
func TestStrIsIPv4(t *testing.T) {
	t.Log(FromString("127.0.0.").IsIP().Error())
	t.Log(FromString("127.0.0.1").IsIP().Error())
	t.Log(FromString("2001:0db8:3c4d:0015:0000:0000:1a2f:1a2b").IsIP().Error())
}

// 测试 FromString().IsTCPAddr()
func TestStrIsTCPAddr(t *testing.T) {
	t.Log(FromString("localhost").IsTCPAddr().Error())
	t.Log(FromString("2001:0db8:3c4d:0015:0000:0000:1a2f:1a2b:80").IsTCPAddr().Error())
	t.Log(FromString("127.0.0.1:80").IsTCPAddr().Error())
	t.Log(FromString("localhost:80").IsTCPAddr().Error())
}

// 测试 FromString().IsMAC()
func TestStrIsMAC(t *testing.T) {
	t.Log(FromString("0800200A8C6D").IsMAC().Error())
	t.Log(FromString("08-00-20-0A-8C-6D").IsMAC().Error())
	t.Log(FromString("08:00:20:0A:8C:6D").IsMAC().Error())
}

// 测试 FromString().IsJSON()
func TestStrIsJSON(t *testing.T) {
	t.Log(FromString(`{"a": "0"`).IsJSON().Error())
	t.Log(FromString(`{"a": "0"}`).IsJSON().Error())
}

// 测试 FromString().IsChineseIDCard()
func TestStrIsChineseIDCard(t *testing.T) {
	t.Log(FromString("000000000000000000").IsChineseIDCard().Error())
	t.Log(FromString("000000000000000001").IsChineseIDCard().Error())
	t.Log(FromString("00000000000000001X").IsChineseIDCard().Error())
}

// 测试 FromString().IsSQLObject()
func TestStrIsSQLObject(t *testing.T) {
	t.Log(FromString("user;name").IsSQLObject().Error())
	t.Log(FromString("username").IsSQLObject().Error())
}

// 测试 FromString().IsURL()
func TestStrIsURL(t *testing.T) {
	t.Log(FromString("www.google.com").IsURL().Error())
	t.Log(FromString("//www.google.com").IsURL().Error())
	t.Log(FromString("custom://www.google.com").IsURL().Error())
	t.Log(FromString("ftp://www.google.com").IsURL().Error())
	t.Log(FromString("https://www.google.com").IsURL().Error())
}

// 测试 FromString().IsUUID()
func TestStrIsUUID(t *testing.T) {
	t.Log(FromString("123e4567:e89b:12d3:a456:426655440000").IsUUID().Error())
	t.Log(FromString("123e4567e89b12d3a456426655440000").IsUUID().Error())
	t.Log(FromString("123e4567-e89b-12d3-a456-426655440000").IsUUID().Error())
}

// 测试 FromString().IsULID()
func TestStrIsULID(t *testing.T) {
	t.Log(FromString("01J1C8KMEXV8JHMS0QDNCJTFS").IsULID().Error())
	t.Log(FromString("01J1C8KMEXV8JHMS0QDNCJTFSZ").IsULID().Error())
}

// 测试 FromString().AllowedChars()
func TestStrAllowedChars(t *testing.T) {
	t.Log(FromString("aaacd").AllowedChars([]rune{'a', 'b', 'c'}).Error())
	t.Log(FromString("aaac").AllowedChars([]rune{'a', 'b', 'c'}).Error())
}

// 测试 FromString().AllowedSymbols()
func TestStrAllowedSymbols(t *testing.T) {
	t.Log(FromString("aa;ac").AllowedSymbols([]rune{'-', '_', '.'}).Error())
	t.Log(FromString("aa_ac").AllowedSymbols([]rune{'-', '_', '.'}).Error())
}

// 测试 FromString().AllowedStrings()
func TestStrAllowedStrings(t *testing.T) {
	t.Log(FromString("mmmm").AllowedStrings([]string{"male", "female", "other"}).Error())
	t.Log(FromString("female").AllowedStrings([]string{"male", "female", "other"}).Error())
}

// 测试 FromString().DeniedChars()
func TestStrDeniedChars(t *testing.T) {
	t.Log(FromString("abdefg").DeniedChars([]rune{'b', 'c'}).Error())
	t.Log(FromString("adefg").DeniedChars([]rune{'b', 'c'}).Error())
}

// 测试 FromString().DeniedSymbols()
func TestStrDeniedSymbols(t *testing.T) {
	t.Log(FromString("a@bc").DeniedSymbols([]rune{'@', '.'}).Error())
	t.Log(FromString("abc").DeniedSymbols([]rune{'@', '.'}).Error())
}

// 测试 FromString().DeniedStrings()
func TestStrDeniedStrings(t *testing.T) {
	t.Log(FromString("abc").DeniedStrings([]string{"male", "abc"}).Error())
	t.Log(FromString("abc").DeniedStrings([]string{"male", "female"}).Error())
}

// 测试 FromString().HasLetter()
func TestStrHasLetter(t *testing.T) {
	t.Log(FromString("0123").HasLetter().Error())
	t.Log(FromString("B0123").HasLetter().Error())
}

// 测试 FromString().HasLower()
func TestStrHasLower(t *testing.T) {
	t.Log(FromString("A0123").HasLower().Error())
	t.Log(FromString("a0123").HasLower().Error())
}

// 测试 FromString().HasUpper()
func TestStrHasUpper(t *testing.T) {
	t.Log(FromString("ab0123").HasUpper().Error())
	t.Log(FromString("aB0123").HasUpper().Error())
}

// 测试 FromString().HasNumber()
func TestStrHasNumber(t *testing.T) {
	t.Log(FromString("abc").HasNumber().Error())
	t.Log(FromString("0abc").HasNumber().Error())
}

// 测试 FromString().HasSymbol()
func TestStrHasSymbol(t *testing.T) {
	t.Log(FromString("0abc").HasSymbol().Error())
	t.Log(FromString("0abc@").HasSymbol().Error())
}

// 测试 FromString().Contains()
func TestStrContains(t *testing.T) {
	t.Log(FromString("abcdefg").Contains("abd").Error())
	t.Log(FromString("abcdefg").Contains("def").Error())
}

// 测试 FromString().HasPrefix()
func TestStrHasPrefix(t *testing.T) {
	t.Log(FromString("abcdefg").HasPrefix("ac").Error())
	t.Log(FromString("abcdefg").HasPrefix("ab").Error())
}

// 测试 FromString().HasSuffix()
func TestStrHasSuffix(t *testing.T) {
	t.Log(FromString("abcdefg").HasSuffix("gf").Error())
	t.Log(FromString("abcdefg").HasSuffix("fg").Error())
}
