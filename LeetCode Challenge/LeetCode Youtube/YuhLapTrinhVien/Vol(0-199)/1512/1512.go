package main

import (
	"fmt"
)

// time complexity: O(n^2) (nested loop)
func numIdenticalPairs(nums []int) int {
	count := 0
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

// time complexity: O(n) (single pass, using map)
func numIdenticalPairs_1(nums []int) int {
	freq := make(map[int]int)
	count := 0

	for _, num := range nums {
		count += freq[num]
		freq[num]++
	}
	return count
}

// code in video: O(n)
func numIdenticalPairs_2(nums []int) int {
	result := 0
	numberCountMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		count, isExist := numberCountMap[nums[i]]
		if isExist {
			result += count
			numberCountMap[nums[i]] = count + 1
		} else {
			numberCountMap[nums[i]] = 1
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	// fmt.Println("Total Good Pairs O(n^2):", numIdenticalPairs(nums))
	// fmt.Println("Total Good Pairs O(n):", numIdenticalPairs_1(nums))
	fmt.Println("Total Good Pairs using map:", numIdenticalPairs_2(nums))
}
