package main

import (
	"fmt"
	"sort"
)

func getMaxScore(arr []int, k int) int {
	
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	maxScore := 0
	used := make(map[int]int) 

	i := 0
	for pick := 0; pick < k && i < len(arr); {
		val := arr[i]
		
		count := countOccurrences(arr, val)
		remaining := k - pick

		take := min(count-used[val], remaining)
		if take > 0 {
			maxScore += val * take
			pick += take
			used[val] += take
		}
		i += count
	}
	return maxScore
}

func countOccurrences(arr []int, val int) int {
	count := 0
	for _, v := range arr {
		if v == val {
			count++
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func main() {
	arr := []int{10, 10, 9, 8, 7, 3, 1}
	k := 4

	fmt.Println("Max score:", getMaxScore(arr, k)) 
}




