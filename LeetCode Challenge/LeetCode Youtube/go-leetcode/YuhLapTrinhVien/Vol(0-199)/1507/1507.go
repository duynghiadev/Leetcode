package main

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {
	// Map months to numbers
	monthMap := map[string]string{
		"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04",
		"May": "05", "Jun": "06", "Jul": "07", "Aug": "08",
		"Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
	}

	// Split string into parts: Day, Month, Year
	parts := strings.Split(date, " ")
	day := parts[0]
	month := parts[1]
	year := parts[2]

	// Process day: remove suffix (st, nd, rd, th) and add leading zero if needed
	dayNum := day[:len(day)-2] // Remove "st", "nd", "rd", or "th"
	if len(dayNum) == 1 {
		dayNum = "0" + dayNum // Add leading zero for single-digit days
	}

	// Get month number from map
	monthNum := monthMap[month]

	// Format result: YYYY-MM-DD
	return fmt.Sprintf("%s-%s-%s", year, monthNum, dayNum)
}

func main() {
	// Test cases
	fmt.Println(reformatDate("20th Oct 2052")) // Output: "2052-10-20"
	fmt.Println(reformatDate("6th Jun 1933"))  // Output: "1933-06-06"
	fmt.Println(reformatDate("26th May 1960")) // Output: "1960-05-26")
}
