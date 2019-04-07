package rule

import "strconv"

// 验证整数最小值
func minInteger(value string, min int64) bool {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false
	}
	if v < min {
		return false
	}
	return true
}

// 验证整数最大值
func maxInteger(value string, max int64) bool {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false
	}
	if v > max {
		return false
	}
	return true
}

// 验证浮点数最小值
func minFloat(value string, min float64) bool {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	if v < min {
		return false
	}
	return true
}

// 验证浮点数最大值
func maxFloat(value string, max float64) bool {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	if v > max {
		return false
	}
	return true
}
