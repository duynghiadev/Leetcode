package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	totalDrunk := numBottles
	emptyBottles := numBottles

	for emptyBottles >= numExchange {
		newBottles := emptyBottles / numExchange
		emptyBottles = emptyBottles%numExchange + newBottles
		totalDrunk += newBottles
	}
	return totalDrunk
}

func main() {
	fmt.Println(numWaterBottles(9, 3))
	fmt.Println(numWaterBottles(15, 4))
}
