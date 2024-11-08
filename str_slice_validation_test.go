package filter

import "testing"

// 测试 FromStringSlice().Require()
func TestStrSliceRequire(t *testing.T) {
	err := FromStringSlice([]string{"a", ""}).Require().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().MinCount()
func TestStrSliceMinCount(t *testing.T) {
	err := FromStringSlice([]string{"", "", ""}).MinCount(3).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().MaxCount()
func TestStrSliceMaxCount(t *testing.T) {
	err := FromStringSlice([]string{"", ""}).MaxCount(2).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().EqualCount()
func TestStrSliceEqualCount(t *testing.T) {
	err := FromStringSlice([]string{"", ""}).EqualCount(2).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().NotEqualCount()
func TestStrSliceNotEqualCount(t *testing.T) {
	err := FromStringSlice([]string{"", ""}).NotEqualCount(3).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().StringLength()
func TestStrSliceLength(t *testing.T) {
	err := FromStringSlice([]string{"aaa", "bbb"}).StringLength(3).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().UTF8StringLength()
func TestStrSliceUTF8Length(t *testing.T) {
	err := FromStringSlice([]string{"龙井", "龙井"}).UTF8StringLength(2).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().StringMinLength()
func TestStrSliceMinLength(t *testing.T) {
	err := FromStringSlice([]string{"aa", "bbb"}).StringMinLength(2).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().StringUTF8MinLength()
func TestStrSliceUTF8MinLength(t *testing.T) {
	err := FromStringSlice([]string{"龙井", "龙井云"}).StringUTF8MinLength(2).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().StringMaxLength()
func TestStrSliceMaxLength(t *testing.T) {
	err := FromStringSlice([]string{"aa", "bbb"}).StringMaxLength(3).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().StringUTF8MaxLength()
func TestStrSliceUTF8MaxLength(t *testing.T) {
	err := FromStringSlice([]string{"龙", "龙井"}).StringUTF8MaxLength(2).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsLower()
func TestStrSliceIsLower(t *testing.T) {
	err := FromStringSlice([]string{"aa", "bb"}).IsLower().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsUpper()
func TestStrSliceIsUpper(t *testing.T) {
	err := FromStringSlice([]string{"AA", "BB"}).IsUpper().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsLetter()
func TestStrSliceIsLetter(t *testing.T) {
	err := FromStringSlice([]string{"Aa", "Bb"}).IsLetter().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsLowerOrNumber()
func TestStrSliceIsLowerOrNumber(t *testing.T) {
	err := FromStringSlice([]string{"a1", "bb"}).IsLowerOrNumber().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsUpperOrNumber()
func TestStrSliceIsUpperOrNumber(t *testing.T) {
	err := FromStringSlice([]string{"A1", "BB"}).IsUpperOrNumber().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsLetterOrNumber()
func TestStrSliceIsLetterOrNumber(t *testing.T) {
	err := FromStringSlice([]string{"A1", "bb"}).IsLetterOrNumber().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsChinese()
func TestStrSliceIsChinese(t *testing.T) {
	err := FromStringSlice([]string{"龙", "井"}).IsChinese().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsMail()
func TestStrSliceIsMail(t *testing.T) {
	err := FromStringSlice([]string{"a@b.com", "b@b.com"}).IsMail().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsIP()
func TestStrSliceIsIP(t *testing.T) {
	err := FromStringSlice([]string{"127.0.0.1", "2001:0db8:3c4d:0015:0000:0000:1a2f:1a2b"}).IsIP().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsTCPAddr()
func TestStrSliceIsTCPAddr(t *testing.T) {
	err := FromStringSlice([]string{"127.0.0.1:80", "localhost:80"}).IsTCPAddr().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsMAC()
func TestStrSliceIsMAC(t *testing.T) {
	err := FromStringSlice([]string{"01:84:F5:04:E5:67", "7F-DB-6C-B3-B5-D1"}).IsMAC().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsSQLObject()
func TestStrSliceIsSQLObject(t *testing.T) {
	err := FromStringSlice([]string{"username", "password"}).IsSQLObject().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsUUID()
func TestStrSliceIsUUID(t *testing.T) {
	err := FromStringSlice([]string{"e1776581-f158-e7b0-eeea-4db795b47ff0", "D73E30E5-6BB3-7BBA-6A4C-500E8CEC3050"}).IsUUID().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().IsULID()
func TestStrSliceIsULID(t *testing.T) {
	err := FromStringSlice([]string{"01J1FEEMCXF8YDQ1NFFS2R8B36", "01J1FEEMCXWB6NQA4Y4PF8A9DK"}).IsULID().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().Contains()
func TestStrSliceContains(t *testing.T) {
	err := FromStringSlice([]string{"Aai1", "B@ib2", "C:ci3"}).Contains("i").Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().HasPrefix()
func TestStrSliceHasPrefix(t *testing.T) {
	err := FromStringSlice([]string{"TestAa1", "TestBb2", "TestCc1-"}).HasPrefix("Test").Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().HasSuffix()
func TestStrSliceHasSuffix(t *testing.T) {
	err := FromStringSlice([]string{"Aa1Test", "Bb2Test", "Cc1-Test"}).HasSuffix("Test").Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().HasLetter()
func TestStrSliceHasLetter(t *testing.T) {
	err := FromStringSlice([]string{"1Aa1", "2Bb2", "3Cc3"}).HasLetter().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().HasLower()
func TestStrSliceHasLower(t *testing.T) {
	err := FromStringSlice([]string{"1Aa1", "2Bb2", "3c3"}).HasLower().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().HasUpper()
func TestStrSliceHasUpper(t *testing.T) {
	err := FromStringSlice([]string{"1Aa1", "2Bb2", "3Cc3"}).HasUpper().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().HasNumber()
func TestStrSliceHasNumber(t *testing.T) {
	err := FromStringSlice([]string{"Aa1", "Bb2", "Cc3"}).HasNumber().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().HasSymbol()
func TestStrSliceHasSymbol(t *testing.T) {
	err := FromStringSlice([]string{"A-a1", "B@b2", "C:c3"}).HasSymbol().Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().AllowedChars()
func TestStrSliceAllowedChars(t *testing.T) {
	err := FromStringSlice([]string{"aaa", "bbb"}).AllowedChars([]rune{'a', 'b'}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().AllowedSymbols()
func TestStrSliceAllowedSymbols(t *testing.T) {
	err := FromStringSlice([]string{"a-a", "b@b"}).AllowedSymbols([]rune{'-', '@'}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().AllowedStrings()
func TestStrSliceAllowedStrings(t *testing.T) {
	err := FromStringSlice([]string{"aa", "bb"}).AllowedStrings([]string{"aa", "bb", "cc"}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().DeniedChars()
func TestStrSliceDeniedChars(t *testing.T) {
	err := FromStringSlice([]string{"cc", "dd"}).DeniedChars([]rune{'a', 'b'}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().DeniedSymbols()
func TestStrSliceDeniedSymbols(t *testing.T) {
	err := FromStringSlice([]string{"a-a", "b@b"}).DeniedSymbols([]rune{';', ':'}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试 FromStringSlice().DeniedStrings()
func TestStrSliceDeniedStrings(t *testing.T) {
	err := FromStringSlice([]string{"cc", "dd", "ee"}).DeniedStrings([]string{"aa", "bb"}).Error()
	if err != nil {
		t.Error(err)
		return
	}
}
