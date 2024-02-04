package filter

import (
	"strings"
)

// 字母转为大写
func (self *Str) ToUpper() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.ToUpper(self.currentValue)
	return self
}

// 字母转为小写
func (self *Str) ToLower() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	self.currentValue = strings.ToLower(self.currentValue)
	return self
}

// SnakeCaseToCamelCase 蛇形转驼峰: hello_world => helloWorld
func (self *Str) SnakeCaseToCamelCase() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	slice := strings.Split(self.currentValue, "_")
	for i := range slice {
		if i > 0 {
			slice[i] = strings.Title(slice[i])
		}
	}
	self.currentValue = strings.Join(slice, "")
	return self
}

// SnakeCaseToPascalCase 蛇形转帕斯卡: hello_world => HelloWorld
func (self *Str) SnakeCaseToPascalCase() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	slice := strings.Split(self.currentValue, "_")
	for i := range slice {
		slice[i] = strings.Title(slice[i])
	}
	self.currentValue = strings.Join(slice, "")
	return self
}

// CamelCaseToSnakeCase 驼峰(含帕斯卡)转蛇形 helloWorld/HelloWorld => hello_world
func (self *Str) CamelCaseToSnakeCase() *Str {
	if self.err != nil || self.currentValue == "" {
		return self
	}
	strLen := len(self.currentValue)
	result := make([]byte, 0, strLen*2)
	j := false
	for i := 0; i < strLen; i++ {
		char := self.currentValue[i]
		if i > 0 && char >= 'A' && char <= 'Z' && j {
			result = append(result, '_')
		}
		if char != '_' {
			j = true
		}
		result = append(result, char)
	}
	self.currentValue = strings.ToLower(string(result))
	return self
}
