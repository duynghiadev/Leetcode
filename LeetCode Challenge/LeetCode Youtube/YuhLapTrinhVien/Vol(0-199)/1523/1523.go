package main

import "fmt"

// Code in video: O(n) -> O(1)
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

func main() {
	oddArray := [2]int{3, 7}
	fmt.Println("Total Odd Numbers:", countOdds(oddArray[0], oddArray[1]))
}
