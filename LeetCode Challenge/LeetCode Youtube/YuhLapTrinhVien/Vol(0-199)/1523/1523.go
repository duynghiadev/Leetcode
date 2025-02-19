package main

import "fmt"

// Type 1: Code in video: O(n) -> O(1)
func countOdds(low int, high int) int {
	result := 0
	i := low
	oddNumbers := []int{} // Slice to store odd numbers

	for i <= high {
		if i%2 == 1 {
			fmt.Println("Found Odd:", i)
			oddNumbers = append(oddNumbers, i) // Store the odd number
			result++
		}
		i++
	}

	fmt.Println("Odd numbers in range:", oddNumbers)
	return result
}

// Type 2: Code in video: faster than type 1.
// Time complexity: O(1)
func countOdds_1(low int, high int) int {
	result := 0
	numOfElement := high - low + 1
	result = numOfElement / 2

	if numOfElement%2 == 1 && low%2 == 1 {
		result++
	}
	return result
}

func main() {
	oddArray := [2]int{3, 7}

	fmt.Println("Total Odd Numbers 1:", countOdds(oddArray[0], oddArray[1]))

	fmt.Println("------------------------")

	fmt.Println("Total Odd Numbers 2:", countOdds_1(oddArray[0], oddArray[1]))
}
