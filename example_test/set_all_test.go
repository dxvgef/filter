package example

import (
	"testing"

	"github.com/dxvgef/filter"
	"github.com/dxvgef/filter/rule"
)

func TestSetAll(t *testing.T) {
	var reqData struct {
		ID    int8
		Email string
	}

	err := filter.SetAll(
		filter.SetFilter(&reqData.ID, "12", rule.Uint, rule.MinInt(10), rule.MaxInt(20)),
		filter.SetFilter(&reqData.Email, "dxvgef@outlook.com", rule.Required, rule.Email),
	)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("ID:", reqData.ID)
	t.Log("Email:", reqData.Email)
}
