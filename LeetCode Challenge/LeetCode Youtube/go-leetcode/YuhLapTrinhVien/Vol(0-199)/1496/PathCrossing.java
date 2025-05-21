import java.util.HashSet;
import java.util.Set;

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
public class PathCrossing {
	public static boolean isPathCrossing(String path) {
		// Initialize set to store visited points
		Set<String> visited = new HashSet<>();
		int x = 0, y = 0;
		visited.add(x + "," + y); // Mark origin as visited

		// Process each direction in the path
		for (char direction : path.toCharArray()) {
			// Update coordinates based on direction
			switch (direction) {
				case 'N': // Move north (↑)
					y++;
					break;
				case 'S': // Move south (↓)
					y--;
					break;
				case 'E': // Move east (→)
					x++;
					break;
				case 'W': // Move west (←)
					x--;
					break;
			}

			// Check if current point was visited
			String point = x + "," + y;
			if (visited.contains(point)) {
				return true; // Path crosses itself!
			}
			visited.add(point); // Mark point as visited
		}

		return false; // No crossing found
	}

	public static void main(String[] args) {
		// Test cases
		System.out.println(isPathCrossing("NES")); // Output: false
		System.out.println(isPathCrossing("NESWW")); // Output: true
	}
}