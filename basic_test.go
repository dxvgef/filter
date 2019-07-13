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

func TestSetRequired(t *testing.T) {
	var username string
	err := Set(&username, FromString("   ", "账号").
		Trim().
		Required("username不能是零值").
		MinLength(6))
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(username)
}

func TestRequired(t *testing.T) {
	username, err := FromString("0", "账号").
		Trim().
		Required("username不能是零值").
		MinInteger(1).Int64()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(username)
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FromString("Abc123", "密码").
			Trim().
			MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol().
			String()
	}
}
