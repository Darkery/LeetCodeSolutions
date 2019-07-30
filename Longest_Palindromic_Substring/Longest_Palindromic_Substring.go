package Longest_Palindromic_Substring


func longestPalindrome(s string) string {
	max := 0
	result := ""
	dp := [1000][1000]int{}
	if len(s) < 2 {
		max = len(s)
		result = s
	} else {
		for i := 0; i < len(s); i++ {
			dp[i][i] = 1
			if len(s[i:i+1]) > max {
				max = len(s[i:i+1])
				result = s[i:i+1]
			}
			if i + 1 >= len(s) {
				continue
			}
			if s[i] == s[i+1] {
				dp[i][i+1] = 1
				if len(s[i:i+2]) > max {
					max = len(s[i:i+2])
					result = s[i:i+2]
				}
			} else {
				dp[i][i+1] = 0
			}
		}

		for i := 0; i < len(s); i++ {
			for j := i + 2; j < len(s); j++{
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
					if dp[i][j] == 1 {
						if len(s[i:j+1]) > max {
							max = len(s[i:j+1])
							result = s[i:j+1]
						}
					}
				} else {
					dp[i][j] = 0
				}
			}

		}
	}

	return result
}
