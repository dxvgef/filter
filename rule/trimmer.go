package rule

import "strings"

// 修剪器
type Trimmer func(string) string

// NewTrimmer 创建新修剪器
func NewTrimmer(trimmer Trimmer) Rule {
	return Rule{
		Trimmer: trimmer,
	}
}

// 去除前后的空格
func trim(value string) string {
	return strings.TrimSpace(value)
}

// 去除所有的空格
func removeSpace(value string) string {
	return strings.ReplaceAll(value, " ", "")
}

// 蛇形转驼峰: hello_world => helloWorld
func snakeCaseToCamelCase(value string) string {
	value = strings.ReplaceAll(value, " ", "")
	slice := strings.Split(value, "_")
	for i := range slice {
		if i > 0 {
			slice[i] = strings.Title(slice[i])
		}
	}
	return strings.Join(slice, "")
}

// 蛇形转帕斯卡: hello_world => HelloWorld
func snakeCaseToPascalCase(value string) string {
	value = strings.ReplaceAll(value, " ", "")
	slice := strings.Split(value, "_")
	for i := range slice {
		slice[i] = strings.Title(slice[i])
	}
	return strings.Join(slice, "")
}

// 驼峰(含帕斯卡)转蛇形 helloWorld/HelloWorld => hello_world
func camelCaseToSnakeCase(str string) string {
	strLen := len(str)
	result := make([]byte, 0, strLen*2)
	j := false
	for i := 0; i < strLen; i++ {
		char := str[i]
		if i > 0 && char >= 'A' && char <= 'Z' && j {
			result = append(result, '_')
		}
		if char != '_' {
			j = true
		}
		result = append(result, char)
	}
	return strings.ToLower(string(result[:]))
}
