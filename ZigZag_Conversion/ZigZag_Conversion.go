package ZigZag_Conversion

import "strings"

func convert(s string, numRows int) string {
	if numRows < 2 || numRows >= len(s) {
		return s
	}
	i := -1
	flag := true
	str := make([]string, numRows)

	for _, j := range s {
		if i <= 0 {
			flag = true
		}
		if i >= numRows - 1 {
			flag = false
		}

		if flag {
			i++
		} else {
			i--
		}

		str[i] += string(j)
	}

	return strings.Join(str, "")
}
