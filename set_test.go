package filter

import "testing"

func TestSet(t *testing.T) {
	var ReqData struct {
		password string
		age      int16
	}
	err := Set(
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
	t.Log("密码", ReqData.password)
	t.Log("年龄", ReqData.age)
}

func TestMSet(t *testing.T) {
	var ReqData struct {
		password string
		age      int16
	}
	err := MSet(
		El(&ReqData.password,
			FromString("Abc123-", "密码").
				MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol(),
		),
		El(&ReqData.age,
			FromString("3", "年龄").
				IsDigit().MinInteger(18)),
	)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("密码", ReqData.password)
	t.Log("年龄", ReqData.age)
}
