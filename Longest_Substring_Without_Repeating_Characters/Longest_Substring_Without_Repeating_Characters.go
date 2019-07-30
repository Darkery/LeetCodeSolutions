package Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max, cur := 1, 0
	for i := 1; i < len(s); i++ {
		j := cur
		for j < i {
			if s[j] == s[i] {
				break
			}
			j++
		}
		if i == j {
			n := j - cur + 1
			if max < n {
				max = n
			}
		} else {
			cur = j + 1
		}
	}
	return max
}
