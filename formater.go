package filter

import (
	"strings"
)

// Trim 去除前后空格
func (obj *Object) Trim() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = strings.TrimSpace(obj.rawValue)
	return obj
}

// 去除所有的空格
func (obj *Object) RemoveSpace() *Object {
	if obj.rawValue == "" {
		return obj
	}
	obj.rawValue = strings.ReplaceAll(obj.rawValue, " ", "")
	return obj
}

// ReplaceAll 替换所有
func (obj *Object) ReplaceAll(old, new string) *Object {
	if obj.rawValue == "" {
		return obj
	}
	count := strings.Count(obj.rawValue, old)
	for i := 0; i < count; i++ {
		obj.rawValue = strings.ReplaceAll(obj.rawValue, old, new)
	}
	return obj
}

// SnakeCaseToCamelCase 蛇形转驼峰: hello_world => helloWorld
func (obj *Object) SnakeCaseToCamelCase() *Object {
	if obj.rawValue == "" {
		return obj
	}
	slice := strings.Split(obj.rawValue, "_")
	for i := range slice {
		if i > 0 {
			slice[i] = strings.Title(slice[i])
		}
	}
	obj.rawValue = strings.Join(slice, "")
	return obj
}

// SnakeCaseToPascalCase 蛇形转帕斯卡: hello_world => HelloWorld
func (obj *Object) SnakeCaseToPascalCase() *Object {
	if obj.rawValue == "" {
		return obj
	}
	slice := strings.Split(obj.rawValue, "_")
	for i := range slice {
		slice[i] = strings.Title(slice[i])
	}
	obj.rawValue = strings.Join(slice, "")
	return obj
}

// CamelCaseToSnakeCase 驼峰(含帕斯卡)转蛇形 helloWorld/HelloWorld => hello_world
func (obj *Object) CamelCaseToSnakeCase() *Object {
	if obj.rawValue == "" {
		return obj
	}
	strLen := len(obj.rawValue)
	result := make([]byte, 0, strLen*2)
	j := false
	for i := 0; i < strLen; i++ {
		char := obj.rawValue[i]
		if i > 0 && char >= 'A' && char <= 'Z' && j {
			result = append(result, '_')
		}
		if char != '_' {
			j = true
		}
		result = append(result, char)
	}
	obj.rawValue = strings.ToLower(string(result))
	return obj
}
