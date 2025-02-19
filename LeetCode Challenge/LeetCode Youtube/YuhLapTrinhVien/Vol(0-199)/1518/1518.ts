function numWaterBottles(numBottles: number, numExchange: number): number {
  let totalDrunk: number = numBottles;
  let emptyBottles: number = numBottles;

  while (emptyBottles >= numExchange) {
    const newBottles: number = Math.floor(emptyBottles / numExchange);
    emptyBottles = (emptyBottles % numExchange) + newBottles;
    totalDrunk += newBottles;
  }
  return totalDrunk;
}

// Main function equivalent
function main(): void {
  console.log(numWaterBottles(9, 3)); // Output: 13
  console.log(numWaterBottles(15, 4)); // Output: 19
}

main();
