package filter

import (
	"strings"
)

// 字母转为大写
func (self *Str) ToUpper() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = strings.ToUpper(self.value)
	return self
}

// 字母转为小写
func (self *Str) ToLower() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	self.value = strings.ToLower(self.value)
	return self
}

// SnakeCaseToCamelCase 蛇形转驼峰: hello_world => helloWorld
func (self *Str) SnakeCaseToCamelCase() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	slice := strings.Split(self.value, "_")
	for i := range slice {
		if i > 0 {
			slice[i] = strings.Title(slice[i])
		}
	}
	self.value = strings.Join(slice, "")
	return self
}

// SnakeCaseToPascalCase 蛇形转帕斯卡: hello_world => HelloWorld
func (self *Str) SnakeCaseToPascalCase() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	slice := strings.Split(self.value, "_")
	for i := range slice {
		slice[i] = strings.Title(slice[i])
	}
	self.value = strings.Join(slice, "")
	return self
}

// CamelCaseToSnakeCase 驼峰(含帕斯卡)转蛇形 helloWorld/HelloWorld => hello_world
func (self *Str) CamelCaseToSnakeCase() StrType {
	if self.err != nil || self.value == "" {
		return self
	}
	strLen := len(self.value)
	result := make([]byte, 0, strLen*2)
	j := false
	for i := 0; i < strLen; i++ {
		char := self.value[i]
		if i > 0 && char >= 'A' && char <= 'Z' && j {
			result = append(result, '_')
		}
		if char != '_' {
			j = true
		}
		result = append(result, char)
	}
	self.value = strings.ToLower(string(result))
	return self
}
