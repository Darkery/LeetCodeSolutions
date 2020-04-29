package main

func findIndex(dp []int, x int) int {
	start := 0
	end := len(dp) - 1

	for start <= end {
		mid := start + (end - start) / 2
		if dp[mid] == x {
			return -1
		}
		if dp[mid] < x {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	println("return = ", start)
	return start
}

func lengthOfLIS(nums []int) (int) {

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, 1)
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < dp[0] {
			dp[0] = nums[i]
		} else if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
		} else {
			index := findIndex(dp, nums[i])
			println("index = ", index)
			if index == -1 {
				continue
			}
			dp[index] = nums[i]
		}

		println("----------")
		for i := 0; i < len(dp); i++ {
			println("dp[", i, "] = ", dp[i])
		}

	}

	return len(dp)
}

func main() {

	//nums := []int {10,9,2,5,3,7,101,18}
	nums := []int {3,5,6,2,5,4,19,5,6,7,12}

	println("----------")

	println("lengthOfLIS = ", lengthOfLIS(nums))

	//println("----------")
	//
	//test := []int {3,5,6,7,10,18}
	//
	//println("findIndex = ", findIndex(test, 2))
}