package example_test

import (
	"testing"

	"github.com/dxvgef/filter"
	"github.com/dxvgef/filter/rule"
)

func TestSet(t *testing.T) {
	var reqData struct {
		ID       int8
		Money    float32
		Email    string
		Disabled bool
	}
	err := filter.Set(&reqData.ID, "13", rule.Uint)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("ID:", reqData.ID)

	err = filter.Set(&reqData.Money, "1.1314", rule.Float)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("Money:", reqData.Money)

	err = filter.Set(&reqData.Email, "dxvgef@outlook.com", rule.Email)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("Email:", reqData.Email)

	err = filter.Set(&reqData.Disabled, "1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("Disabled:", reqData.Disabled)
}
