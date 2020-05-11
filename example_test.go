package filter

import (
	"testing"
)

func TestError(t *testing.T) {
	t.Log(FromString("    ", "账号").
		RemoveSpace().
		IsChineseMobile().
		Error(),
	)
}

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
		MinLength(6).MaxLength(32).MustHasLetter().MustHasUpper().MustHasDigit().MustHasSymbol().
		DefaultString("默认值")
	t.Log(password)
}

func TestRequired(t *testing.T) {
	age, err := FromString(" ", "年龄").
		RemoveSpace().
		Required().
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
			RemoveSpace().Required().
			MinLength(3).MaxLength(32),
	)
	if err != nil {
		t.Log(err.Error())
		return
	}

	err = Set(
		&ReqData.password, FromString("Abc123-", "密码").
			MinLength(6).MaxLength(32).MustHasLetter().MustHasUpper().MustHasDigit().MustHasSymbol(),
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
		&ReqData.other, FromString("25", "其它").MinInteger(18).Default(20),
	)
	if err != nil {
		t.Log(err.Error())
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
		columns  []string
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
				MinLength(6).MaxLength(32).MustHasLetter().MustHasUpper().MustHasDigit().MustHasSymbol(),
		),
		El(&ReqData.columns,
			FromString("abc", "字段").
				RemoveSpace().EnumString([]string{"abc", "def"}).Separator(",")),
	)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("账号", ReqData.username)
	t.Log("密码", ReqData.password)
	t.Log("字段", ReqData.columns)
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FromString("Abc123", "密码").
			Trim().
			MinLength(6).MaxLength(32).MustHasLetter().MustHasUpper().MustHasDigit().MustHasSymbol().
			String()
	}
}
