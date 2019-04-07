package rule

import (
	"strings"
	"regexp"
	"unicode"
	"bytes"
)

type StringOpt func(v string) string

// 字符串优化函数
func NewStringOpt(opt... StringOpt) Rule {
	return Rule{
		StringOpts: opt,
	}
}


//  iki: hello world => helloworld
func TrimAllString(s string) string {
	return regexp.MustCompile("\\s+").ReplaceAllString(s, "")
}

// iki: hello world => helloWorld 除去空格转驼峰式
func TrimStringAndCamel(s string) string {
	var buffer bytes.Buffer
	f := strings.FieldsFunc(s, unicode.IsSpace)
	for _, v := range f {
		buffer.WriteString(strings.Title(v))
	}
	return buffer.String()
}

// iki: hello world => hello_world 去除空格转蛇底式
func TrimStringAndSnake(s string) string {
	return strings.Join(strings.FieldsFunc(s, unicode.IsSpace), "_")
}

// iki: HelloWorld => hello_world
func SnakeString(v string) string {
	lens := len(v)
	res := make([]byte, 0, lens*2)
	j := false
	for i := 0; i < lens; i++ {
		d := v[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			res = append(res, '_')
		}
		if d != '_' {
			j = true
		}
		res = append(res, d)
	}
	return strings.ToLower(string(res[:]))
}

// iki: hello_world => helloWorld
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if !k && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || !k) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// iki: hello_world => helloWorld
func CamelCase(in string) string {
	slice := strings.Split(in, "_")
	for i := range slice {
		slice[i] = strings.Title(strings.Trim(slice[i], " "))
	}
	return strings.Join(slice, "")
}
