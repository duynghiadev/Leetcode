package main

import "fmt"

func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))

	// Place each character at its correct position
	for index, value := range indices {
		// value = indices[index]
		result[value] = s[index]
		// In result dưới dạng mảng ký tự
		fmt.Print("result [")
		for i, b := range result {
			if b == 0 {
				fmt.Print("_")
			} else {
				fmt.Printf("%c", b)
			}
			if i < len(result)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
		fmt.Printf("index=%d value=%d result[value]=%c s=%s s[index]=%c\n",
			index, value, result[value], s, s[index])
	}

	return string(result)
}

func main() {
	// Example 1
	s1 := "codeleet"
	indices1 := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Printf("\n👉 Example 1:\nInput: s = \"%s\", indices = %v\nOutput: \"%s\"\n\n",
		s1, indices1, restoreString(s1, indices1))

	fmt.Println("------------------------")

	// Example 2
	s2 := "abc"
	indices2 := []int{0, 1, 2}
	fmt.Printf("\n👉 Example 2:\nInput: s = \"%s\", indices = %v\nOutput: \"%s\"\n",
		s2, indices2, restoreString(s2, indices2))
}
