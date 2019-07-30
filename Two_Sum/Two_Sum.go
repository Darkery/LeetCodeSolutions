package Two_Sum

func twoSum(nums []int, target int) []int {

	result := make([]int, 2)
	map1 := make(map[int]int)

	for i := 0; i < len(nums); i++{

		index, has_index := map1[nums[i]]

		if has_index{
			result[0] = index
			result[1] = i
			return result
		}else{
			map1[target-nums[i]] = i
		}
	}
	return nil
}
