package filter

import "strings"

func ReplaceAll(str, old, new string) string {
	count := strings.Count(str, old)
	for i := 0; i < count; i++ {
		str = strings.ReplaceAll(str, old, new)
	}
	return str
}
