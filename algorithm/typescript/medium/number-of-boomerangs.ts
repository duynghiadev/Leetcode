export function getDis(p1: number[], p2: number[]): number {
  return (p1[0] - p2[0]) ** 2 + (p1[1] - p2[1]) ** 2;
}

export function numberOfBoomerangs(points: number[][]): number {
  let res = 0;
  for (let i = 0; i < points.length; i++) {
    const map = new Map<number, number>();
    for (let j = 0; j < points.length; j++) {
      if (i === j) continue; // Skip the same point
      const dis = getDis(points[i], points[j]); // Get the distance squared
      map.set(dis, (map.get(dis) || 0) + 1); // Increment count of the distance
    }
    for (const v of map.values()) {
      res += v * (v - 1); // For each distance, calculate the number of boomerangs
    }
  }
  return res;
}

// Test the function
const points = [
  [0, 0],
  [1, 0],
  [2, 0],
];
console.log(numberOfBoomerangs(points)); // Output should be 2
