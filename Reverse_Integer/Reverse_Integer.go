package Reverse_Integer

func reverse(x int) int {
	reg := make([]int, 10)
	flag := true
	result := 0
	i := 0
	if x < 0 {
		flag = false
		x = -x
	}

	for x != 0 {
		reg[i] = x % 10
		i++
		x /= 10
	}

	for j := 0; j < i; j++ {
		result *= 10
		result += reg[j]
	}

	if result >= 2147483648 {
		result = 0
	}

	if !flag {
		result = -result
	}

	return result
}
