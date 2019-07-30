package Median_of_Two_Sorted_Arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	reg := make(map[int]int)
	i := 0
	j := 0
	k := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			reg[k] = nums1[i]
			k++
			i++
		} else {
			reg[k] = nums2[j]
			k++
			j++
		}
	}

	for i < len(nums1) {
		reg[k] = nums1[i]
		k++
		i++
	}

	for j < len(nums2) {
		reg[k] = nums2[j]
		k++
		j++
	}

	if len(reg) % 2 == 1 {
		result := (float64)(reg[len(reg) / 2])
		return result
	} else {
		result := ((float64)(reg[len(reg) / 2]) + (float64)(reg[len(reg) / 2 - 1])) /2
		return result
	}
}
