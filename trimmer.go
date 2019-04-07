package filter

// 修剪器
type Trimmer func(string) string

// NewTrimmer 创建新修剪器
func NewTrimmer(trimmer Trimmer) Rule {
	return Rule{
		trimmer: trimmer,
	}
}
