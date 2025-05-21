/*
 * 🚶 Path Crossing Checker
 * ====================
 *
 * 📝 How it works:
 * 1. Start at origin (0,0) ⭐
 * 2. Track visited points in a set 📍
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
function isPathCrossing(path: string): boolean {
  // Initialize set to store visited points 📍
  const visited = new Set<string>();
  let x = 0,
    y = 0;
  // Mark origin as visited ⭐
  visited.add(`${x},${y}`);

  // Process each direction in the path 🚶
  for (const direction of path) {
    // Update coordinates based on direction ➡️⬅️⬆️⬇️
    switch (direction) {
      case "N":
        y++; // Move north ⬆️
        break;
      case "S":
        y--; // Move south ⬇️
        break;
      case "E":
        x++; // Move east ➡️
        break;
      case "W":
        x--; // Move west ⬅️
        break;
    }

    // Check if current point was visited before 🔍
    const point = `${x},${y}`;
    if (visited.has(point)) {
      return true; // Path crosses! ✅
    }
    // Mark current point as visited 📍
    visited.add(point);
  }

  // No crossing found 🚫
  return false;
}

// Test cases 🧪
console.log(isPathCrossing("NES")); // Output: false
console.log(isPathCrossing("NESWW")); // Output: true
