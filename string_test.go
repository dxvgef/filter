package filter

import "testing"

func TestString(t *testing.T) {
	username, err := FromString("username", "账号").Trim().MinLength(3).MaxLength(6).String()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(username)
}
