/*
 * 🚶 Path Crossing Checker
 * ====================
 *
 * 📝 How it works:
 * 1. Start at origin (0,0) ⭐
 * 2. Track all visited points in a set
 * 3. Move according to path (N↑, S↓, E→, W←)
 * 4. If we revisit any point → return true
 *
 * 🎯 Example: "NESWW"
 * Start: (0,0) ⭐
 * N → (0,1)  ↑
 * E → (1,1)  →
 * S → (1,0)  ↓
 * W → (0,0)  ← (visited before! → true)
 * W → (-1,0) ←
 */
function isPathCrossing(path: string): boolean {
  // Initialize set to store visited points
  const visited = new Set<string>();
  let x = 0,
    y = 0;
  visited.add(`${x},${y}`); // Mark origin as visited

  // Process each direction in the path
  for (const dir of path) {
    // Update coordinates based on direction
    switch (dir) {
      case "N": // Move north (↑)
        y++;
        break;
      case "S": // Move south (↓)
        y--;
        break;
      case "E": // Move east (→)
        x++;
        break;
      case "W": // Move west (←)
        x--;
        break;
    }

    // Check if current point was visited
    const point = `${x},${y}`;
    if (visited.has(point)) {
      return true; // Path crosses itself!
    }
    visited.add(point); // Mark point as visited
  }

  return false; // No crossing found
}

// Test cases
console.log(isPathCrossing("NES")); // Output: false
console.log(isPathCrossing("NESWW")); // Output: true
