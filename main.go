package main

import (
	"fmt"
)

func minSubsequenceLength(n int, arr []int) interface{} {
	required := 26
	freq := make(map[int]int)
	left := 0
	minLen := n + 1
	unique := 0

	for right := 0; right < n; right++ {
		num := arr[right]
		if freq[num] == 0 {
			unique++
		}
		freq[num]++

		for unique == required {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
			}

			leftNum := arr[left]
			if 1 <= leftNum && leftNum <= 26 {
				freq[leftNum]--
				if freq[leftNum] == 0 {
					unique--
				}
			}
			left++
		}
	}

	if minLen <= n {
		return minLen
	}
	return "NONE"
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	result := minSubsequenceLength(n, arr)
	fmt.Println(result)
}
