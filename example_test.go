package filter

import "testing"

func TestGetValue(t *testing.T) {
	username, err := FromString("dxvgef", "账号").
		RemoveSpace().
		MinLength(3).MaxLength(32).
		String()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(username)

	password := FromString("Abc", "密码").
		Trim().
		MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol().
		DefaultString("默认值")
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(password)
}

func TestRequired(t *testing.T) {
	age, err := FromString(" ", "年龄").
		RemoveSpace().
		MinInteger(1).Int64()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(age)
}

func TestSet(t *testing.T) {
	var ReqData struct {
		username string
		password string
		age      int16
		other    int
	}
	err := Set(
		&ReqData.username, FromString("dxvgef", "账号").
			RemoveSpace().
			MinLength(3).MaxLength(32),
	)
	if err != nil {
		t.Log(err.Error())
		return
	}

	err = Set(
		&ReqData.password, FromString("Abc123-", "密码").
			MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol(),
	)
	if err != nil {
		t.Log(err.Error())
		return
	}

	err = Set(
		&ReqData.age, FromString("100", "年龄").IsDigit().MinInteger(18),
	)
	if err != nil {
		t.Log(err.Error())
		return
	}

	err = Set(
		&ReqData.other, FromString("aaa", "其它").Silent().Default(20),
	)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("账号", ReqData.username)
	t.Log("密码", ReqData.password)
	t.Log("年龄", ReqData.age)
	t.Log("其它", ReqData.other)
}

func TestMSet(t *testing.T) {
	var ReqData struct {
		username string
		password string
		age      int16
	}
	err := MSet(
		El(&ReqData.username,
			FromString("dxvgef", "账号").
				RemoveSpace().Required().
				MinLength(3).MaxLength(32),
		),
		El(&ReqData.password,
			FromString("Abc123-", "密码").
				Required().
				MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol(),
		),
		El(&ReqData.age,
			FromString("", "年龄").
				RemoveSpace().Required().
				IsDigit().MinInteger(18)),
	)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("账号", ReqData.username)
	t.Log("密码", ReqData.password)
	t.Log("年龄", ReqData.age)
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FromString("Abc123", "密码").
			Trim().
			MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol().
			String()
	}
}
