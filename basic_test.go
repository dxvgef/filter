package filter

import "testing"

func TestBasic(t *testing.T) {
	password, err := FromString("Abc123", "密码").
		Trim().
		MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol().
		String()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(password)
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FromString("Abc123", "密码").
			Trim().
			MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol().
			String()
	}
}
