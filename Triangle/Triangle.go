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

	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	sum := 0
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {

			x := 65535
			y := 65535
			if j - 1 >= 0 {
				x = dp[i-1][j-1]
			}

			if j < len(triangle[i-1]) {
				y = dp[i-1][j]
			}

			dp[i][j] = triangle[i][j] + min(x, y)
			if i == (len(triangle) - 1) {
				if j == 0 {
					sum = dp[i][j]
				}
				sum = min(dp[i][j], sum)
			}
			fmt.Printf("dp[%d][%d] = %d\n", i, j, dp[i][j])
		}
	}

	return sum
}

func main() {

	triangle := [][]int {{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}

	println("----------")

	println("minimumTotal = ", minimumTotal(triangle))
}