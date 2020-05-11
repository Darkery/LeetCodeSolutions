package main

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxProduct(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	product := nums[0]
	temp_max, temp_min := 1, 1

	for i := 0; i < len(nums); i++ {

		if nums[i] < 0 {
			temp_max, temp_min = temp_min, temp_max
		}
		temp_max = max(nums[i], nums[i]*temp_max)
		temp_min = min(nums[i], nums[i]*temp_min)

		product = max(product, temp_max)
	}
	
	return product
}

func main() {

	triangle := []int {-2, 3, -4}

	println("----------")

	println("maxProduct = ", maxProduct(triangle))
}