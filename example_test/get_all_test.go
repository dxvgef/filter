package example

import (
	"testing"

	"github.com/dxvgef/filter"
	"github.com/dxvgef/filter/rule"
)

func TestGetAll(t *testing.T) {
	var reqData struct {
		ID    int8 `rule:""`
		Email string
	}

	err := filter.GetAll(
		filter.Filter("12", rule.Uint, rule.MinInt(10), rule.MaxInt(20)),
		filter.Filter("dxvgef@outlook.com", rule.Required, rule.Email),
	)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("ID:", reqData.ID)
}
