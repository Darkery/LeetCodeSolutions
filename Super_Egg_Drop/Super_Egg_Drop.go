package main

import (
	"fmt"
	"math"
)

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

func get_index(K int, N int) string {
	return fmt.Sprintf("%d%d", K, N)
}

func dp(K int, N int, countMap *map[string]int) int {
	index := get_index(K, N)
	count, has_count := (*countMap)[index]
	if !has_count {
		count = math.MaxInt32
		if N == 0{
			count = 0
		} else if K == 1 {
			count = N
		} else {
			start := 1
			end := N
			for start <= end {
				x := (end - start) / 2 + start
				t1 := dp(K-1, x-1, countMap)
				t2 := dp(K, N-x, countMap)
				count = min(count, max(t1, t2) + 1)
				if t1 < t2 {
					start = x + 1
				} else {
					end = x - 1
				}
			}
			(*countMap)[index] = count
		}
	}
	return count
}

func superEggDrop(K int, N int) int {
	countMap := make(map[string]int)
	return dp(K, N, &countMap)
}

func main() {
	println("----------")
	println("superEggDrop = ", superEggDrop(2, 2))
}