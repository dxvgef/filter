package example_test

import (
	"testing"

	"github.com/dxvgef/filter"
	"github.com/dxvgef/filter/rule"
)

func TestGet(t *testing.T) {
	s, err := filter.Get("     test      ", rule.Trim)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("string:", s)

	l, err := filter.Get("1二3四5六7", rule.MinLength(2), rule.MaxLength(6))
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("length:", l)

	i, err := filter.Get("1", rule.Int)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("int:", i)

	f, err := filter.Get("1.1314", rule.Float)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("float:", f)

	email, err := filter.Get("dxvgef@outlook.com", rule.Email)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("email:", email)
}
