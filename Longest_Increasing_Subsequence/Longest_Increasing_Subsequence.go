package main

func lengthOfLIS(nums []int) (int) {

	if len(nums) == 0 {
		return 0
	}

	length := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if  dp[i] < dp[j] + 1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if length < dp[i] {
			length = dp[i]
		}
	}

	return length
}

func main() {

	nums := []int {10,9,2,5,3,7,101,18}

	println("----------")

	println("lengthOfLIS = ", lengthOfLIS(nums))
}