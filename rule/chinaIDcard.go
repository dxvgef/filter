package rule

import "strconv"

// 中国大陆地区身份证号码
func chinaIDCard(str string) bool {
	return verifyID(checkID(str[:17]), byteToInt(str[17:]))
}

func byteToInt(x string) int {
	if x == "X" {
		return 88
	}

	res, _ := strconv.Atoi(x)

	return res
}

func checkID(id string) int { // len(id)= 17
	arry := make([]int, 17)

	for i := 0; i < 17; i++ {
		arry[i], _ = strconv.Atoi(string(id[i]))
	}
	var wi [17]int = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		res += arry[i] * wi[i]
	}
	return (res % 11)
}

func verifyID(verify int, id_v int) bool {
	var temp int
	var i int
	a18 := [11]int{1, 0, 88 /* 'X' */, 9, 8, 7, 6, 5, 4, 3, 2}

	for i = 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			break
		}
	}
	if temp == id_v {
		return true
	}

	return false
}
