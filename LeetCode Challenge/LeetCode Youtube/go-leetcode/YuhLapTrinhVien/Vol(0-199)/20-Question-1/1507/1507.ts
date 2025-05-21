function reformatDate(date: string): string {
  // Map months to numbers
  const monthMap: Record<string, string> = {
    Jan: "01",
    Feb: "02",
    Mar: "03",
    Apr: "04",
    May: "05",
    Jun: "06",
    Jul: "07",
    Aug: "08",
    Sep: "09",
    Oct: "10",
    Nov: "11",
    Dec: "12",
  };

  // Split string into parts: Day, Month, Year
  const parts: string[] = date.split(" ");
  const day: string = parts[0];
  const month: string = parts[1];
  const year: string = parts[2];

  // Process day: remove suffix (st, nd, rd, th) and add leading zero if needed
  let dayNum: string = day.slice(0, day.length - 2); // Remove "st", "nd", "rd", or "th"
  if (dayNum.length === 1) {
    dayNum = "0" + dayNum; // Add leading zero for single-digit days
  }

  // Get month number from map
  const monthNum: string = monthMap[month];

  // Format result: YYYY-MM-DD
  return `${year}-${monthNum}-${dayNum}`;
}

// Test cases
console.log(reformatDate("20th Oct 2052")); // Output: "2052-10-20"
console.log(reformatDate("6th Jun 1933")); // Output: "1933-06-06"
console.log(reformatDate("26th May 1960")); // Output: "1960-05-26"
