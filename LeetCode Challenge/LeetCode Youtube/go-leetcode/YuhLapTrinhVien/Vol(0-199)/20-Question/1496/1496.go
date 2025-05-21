package main

import "fmt"

/*
 * 🚶 Path Crossing Checker
 * ====================
 *
 * 📝 How it works:
 * 1. Start at origin (0,0) ⭐
 * 2. Track visited points in a map 📍
 * 3. Move according to path (N,S,E,W) ➡️⬅️⬆️⬇️
 * 4. If we revisit any point → return true ✅
 *
 * 🎯 Example: "NESWW"
 * Start: (0,0) ⭐
 * N → (0,1)  ⬆️
 * E → (1,1)  ➡️
 * S → (1,0)  ⬇️
 * W → (0,0)  ⬅️ (visited before! → true ✅)
 * W → (-1,0) ⬅️
 */
func isPathCrossing(path string) bool {
	// Initialize map to store visited points 📍
	visited := make(map[[2]int]bool)
	x, y := 0, 0
	// Mark origin as visited ⭐
	visited[[2]int{x, y}] = true

	// Process each direction in the path 🚶
	for _, direction := range path {
		// Update coordinates based on direction ➡️⬅️⬆️⬇️
		switch direction {
		case 'N':
			y++ // Move north ⬆️
		case 'S':
			y-- // Move south ⬇️
		case 'E':
			x++ // Move east ➡️
		case 'W':
			x-- // Move west ⬅️
		}

		// Check if current point was visited before 🔍
		point := [2]int{x, y}
		if visited[point] {
			return true // Path crosses! ✅
		}
		// Mark current point as visited 📍
		visited[point] = true
	}

	// No crossing found 🚫
	return false
}

func main() {
	// Test cases 🧪
	fmt.Println(isPathCrossing("NES"))   // Output: false
	fmt.Println(isPathCrossing("NESWW")) // Output: true
}
