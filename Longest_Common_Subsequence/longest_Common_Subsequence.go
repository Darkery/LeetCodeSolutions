package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {

	if text1 == "" || text2 == "" {
		return 0
	}

	arr_1 := []byte(text1)
	arr_2 := []byte(text2)

	dp := make([][]int, len(arr_1)+1)
	for i := range dp {
		dp[i] = make([]int, len(arr_2)+1)
	}

	for i := 1; i <= len(arr_1); i++ {
		for j := 1; j <= len(arr_2); j++ {

			max := 0
			if dp[i][j-1] > dp[i-1][j] {
				max = dp[i][j-1]
			} else {
				max = dp[i-1][j]
			}

			if arr_1[i-1] == arr_2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max
			}

			fmt.Printf("dp[%d][%d] = %d\n", i, j, dp[i][j])
		}
	}
	return dp[len(arr_1)][len(arr_2)]
}

func main() {

	a := "bfff"
	b := "abcb"

	println("----------")

	println("longestCommonSubsequence = ", longestCommonSubsequence(a, b))
}