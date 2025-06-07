package main

import (
	"fmt"
)

func minNum(samDaily, kellyDaily, difference int) int {
	if kellyDaily <= samDaily {
		return -1
	}
	// hitung minimal hari
	diff := difference
	rateDiff := kellyDaily - samDaily

	// days = difference / rateDiff, pembulatan ke atas
	days := diff / rateDiff
	if diff%rateDiff != 0 {
		days++
	}
	return days
}

func main() {
	fmt.Println(minNum(3, 5, 1)) // Output: 1
	fmt.Println(minNum(4, 5, 1)) // Output: 2
	fmt.Println(minNum(5, 5, 10)) // Output: -1
}
