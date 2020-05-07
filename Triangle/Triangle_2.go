package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minimumTotal(triangle [][]int) int {

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	sum := 0
	dp := make([]int, len(triangle[len(triangle)-1]))

	dp[0] = triangle[0][0]
	x := triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			y := dp[j]

			if j == 0 {
				dp[j] = triangle[i][j] + y
			} else if j == len(triangle[i-1]) {
				dp[j] = triangle[i][j] + x
			} else {
				dp[j] = triangle[i][j] + min(x, y)
			}
			if i == (len(triangle) - 1) {
				if j == 0 {
					sum = dp[j]
				}
				sum = min(dp[j], sum)
			}
			fmt.Printf("dp[%d] = %d\n", j, dp[j])

			x = y
		}
	}

	return sum
}

func main() {

	triangle := [][]int {{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}

	println("----------")

	println("minimumTotal = ", minimumTotal(triangle))
}