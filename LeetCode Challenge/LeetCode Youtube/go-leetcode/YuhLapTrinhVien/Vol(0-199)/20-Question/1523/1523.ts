function countOdds(low: number, high: number): number {
  let result = 0;
  let i = low;
  let oddNumbers: number[] = []; // Array to store odd numbers

  while (i <= high) {
    if (i % 2 === 1) {
      console.log("Found Odd:", i);
      oddNumbers.push(i); // Store the odd number
      result++;
    }
    i++;
  }

  console.log("Odd numbers in range:", oddNumbers);
  return result;
}

function countOdds_1(low: number, high: number): number {
  let numOfElement = high - low + 1;
  let result = Math.floor(numOfElement / 2);

  if (numOfElement % 2 === 1 && low % 2 === 1) {
    result++;
  }
  return result;
}

// Test case
const oddArray: [number, number] = [3, 7];

console.log("Total Odd Numbers 1:", countOdds(oddArray[0], oddArray[1]));

console.log("------------------------");

console.log("Total Odd Numbers 2:", countOdds_1(oddArray[0], oddArray[1]));
