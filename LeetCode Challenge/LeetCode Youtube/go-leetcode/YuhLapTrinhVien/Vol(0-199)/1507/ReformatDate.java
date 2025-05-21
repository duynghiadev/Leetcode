import java.util.HashMap;
import java.util.Map;

public class ReformatDate {
	public static String reformatDate(String date) {
		// Map months to numbers
		Map<String, String> monthMap = new HashMap<>();
		monthMap.put("Jan", "01");
		monthMap.put("Feb", "02");
		monthMap.put("Mar", "03");
		monthMap.put("Apr", "04");
		monthMap.put("May", "05");
		monthMap.put("Jun", "06");
		monthMap.put("Jul", "07");
		monthMap.put("Aug", "08");
		monthMap.put("Sep", "09");
		monthMap.put("Oct", "10");
		monthMap.put("Nov", "11");
		monthMap.put("Dec", "12");

		// Split string into parts: Day, Month, Year
		String[] parts = date.split(" ");
		String day = parts[0];
		String month = parts[1];
		String year = parts[2];

		// Process day: remove suffix (st, nd, rd, th) and add leading zero if needed
		String dayNum = day.substring(0, day.length() - 2); // Remove "st", "nd", "rd", or "th"
		if (dayNum.length() == 1) {
			dayNum = "0" + dayNum; // Add leading zero for single-digit days
		}

		// Get month number from map
		String monthNum = monthMap.get(month);

		// Format result: YYYY-MM-DD
		return String.format("%s-%s-%s", year, monthNum, dayNum);
	}

	public static void main(String[] args) {
		// Test cases
		System.out.println(reformatDate("20th Oct 2052")); // Output: "2052-10-20"
		System.out.println(reformatDate("6th Jun 1933")); // Output: "1933-06-06"
		System.out.println(reformatDate("26th May 1960")); // Output: "1960-05-26"
	}
}