package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func quickSortOuter(arr *[][]int, start int, end int)  {
	if start >= end {
		return
	}
	i, j := start, end
	for i < j {

		for i < j && (*arr)[i][0] <= (*arr)[j][0]  {
			j--
		}
		if i < j {
			(*arr)[i][0], (*arr)[j][0] = (*arr)[j][0], (*arr)[i][0]
			(*arr)[i][1], (*arr)[j][1] = (*arr)[j][1], (*arr)[i][1]
		}
		for i < j && (*arr)[i][0] <= (*arr)[j][0]  {
			i++
		}
		if i < j {
			(*arr)[j][0], (*arr)[i][0] = (*arr)[i][0], (*arr)[j][0]
			(*arr)[j][1], (*arr)[i][1] = (*arr)[i][1], (*arr)[j][1]
		}
	}

	quickSortOuter(arr, start, i-1)
	quickSortOuter(arr, i+1, end)
}

func quickSortInner(arr *[][]int, start int, end int)  {
	if start >= end {
		return
	}
	i, j := start, end
	for i < j {

		for i < j && (*arr)[i][1] >= (*arr)[j][1]  {
			j--
		}
		if i < j {
			(*arr)[i][1], (*arr)[j][1] = (*arr)[j][1], (*arr)[i][1]
		}
		for i < j && (*arr)[i][1] >= (*arr)[j][1]  {
			i++
		}
		if i < j {
			(*arr)[j][1], (*arr)[i][1] = (*arr)[i][1], (*arr)[j][1]
		}
	}

	quickSortInner(arr, start, i-1)
	quickSortInner(arr, i+1, end)
}
	
func sort_outer(envelopes *[][]int) {
	quickSortOuter(envelopes, 0, len(*envelopes)-1)
}

func sort_inner(envelopes *[][]int) {
	for i := 0; i < len(*envelopes); i++ {
		temp := i
		for j := i; j < len(*envelopes); j++ {
			if (*envelopes)[i][0] != (*envelopes)[j][0] {
				temp = j - 1
				break
			}
			if j == len(*envelopes) - 1 {
				temp = j
				break
			}
		}
		quickSortInner(envelopes, i, temp)
		i = temp
	}
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 || len(envelopes) == 1 {
		return len(envelopes)
	}

	sort_outer(&envelopes)
	sort_inner(&envelopes)

	dp := make([]int, len(envelopes))
	result := 0
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}

		}
		result = max(result, dp[i])
	}
	return result
}

func print_envelopes(envelopes *[][]int) {
	for i := 0; i < len(*envelopes); i++ {
		fmt.Printf("env[%d][0] = %d\n", i, (*envelopes)[i][0])
		fmt.Printf("env[%d][1] = %d\n", i, (*envelopes)[i][1])
	}
}

func main() {
	envelopes := [][]int {{1,15},{7,18},{7,6},{7,100},{2,200},{17,30},{17,45},{3,5},{7,8},{3,6},{3,10},{7,20},{17,3},{17,45}}
	//envelopes := [][]int {{2,100},{3,200},{4,300},{5,500},{5,400},{5,250},{6,370},{6,360},{7,380}}
	//envelopes := [][]int {{46,89},{50,53},{52,68},{72,45},{77,81}}
	//envelopes := [][]int {{5,4},{6,4},{6,7},{2,3}}
	//envelopes := [][]int {{1,1},{1,1},{1,1}}

	sort_outer(&envelopes)
	sort_inner(&envelopes)
	println("------ sort ------")
	print_envelopes(&envelopes)

	println("------------------")
	println("maxEnvelopes = ", maxEnvelopes(envelopes))
}