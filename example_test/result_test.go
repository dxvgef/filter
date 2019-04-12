package example_test

import (
	"testing"

	"github.com/dxvgef/filter"
	"github.com/dxvgef/filter/rule"
)

func TestString(t *testing.T) {
	s, err := filter.String("     test      ", rule.Trim)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("string:", s)

	l, err := filter.String("1二3四5六7", rule.MinLength(2), rule.MaxLength(6))
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("length:", l)

	i, err := filter.String("1", rule.Int)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("int:", i)

	f, err := filter.String("1.1314", rule.Float)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("float:", f)

	email, err := filter.String("dxvgef@outlook.com", rule.Email)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("email:", email)
}

func TestResult(t *testing.T) {
	var reqData struct {
		ID       int8
		Money    float32
		Email    string
		Disabled bool
	}
	err := filter.Result(&reqData.ID, "13", rule.Uint)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("ID:", reqData.ID)

	err = filter.Result(&reqData.Money, "1.1314", rule.Float)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("Money:", reqData.Money)

	err = filter.Result(&reqData.Email, "dxvgef@outlook.com", rule.Email)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("Email:", reqData.Email)

	err = filter.Result(&reqData.Disabled, "1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("Disabled:", reqData.Disabled)
}

func TestAllResult(t *testing.T) {
	var reqData struct {
		ID    int8
		Email string
	}
	err := filter.AllResult(
		filter.Filter(&reqData.ID, "12", rule.Uint, rule.MinInt(10), rule.MaxInt(20)),
		filter.Filter(&reqData.Email, "dxvgef@outlook.com", rule.Required, rule.Email),
	)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("ID:", reqData.ID)
	t.Log("Email:", reqData.Email)
}
