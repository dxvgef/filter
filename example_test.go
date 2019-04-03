package filter

import (
	"testing"

	"filter/rule"
)

func TestEmail(t *testing.T) {
	value := "dxvgef@outlook.com"
	err := Result(value, rule.Email)
	if err != nil {
		t.Log(err.Error())
	}
}

func TestTrim(t *testing.T) {
	value := "    127.0.0.1      "
	err := Result(value, rule.Trim, rule.IP.Error("请填写正确的IP地址"))
	if err != nil {
		t.Log(err.Error())
	}
}

func TestCustom(t *testing.T) {
	value := "    127.0.0.1      "
	var f = func(value string) bool {
		return true
	}
	customRule := rule.NewValidator(f, "默认消息")
	err := Result(value, rule.Trim, customRule)
	if err != nil {
		t.Log(err.Error())
	}
}
