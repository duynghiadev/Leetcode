import java.util.Arrays;

public class ShuffleString {

	public static String restoreString(String s, int[] indices) {
		char[] result = new char[s.length()];

		// Place each character at its correct position
		for (int index = 0; index < indices.length; index++) {
			int value = indices[index];
			result[value] = s.charAt(index);

			System.out.print("result [");
			for (int i = 0; i < result.length; i++) {
				if (result[i] == 0) {
					System.out.print("_");
				} else {
					System.out.print(result[i]);
				}
				if (i < result.length - 1) {
					System.out.print(" ");
				}
			}
			System.out.println("]");

			System.out.printf("index=%d value=%d result[value]=%c s=%s s[index]=%c%n",
					index, value, result[value], s, s.charAt(index));
		}

		return new String(result);
	}

	public static void main(String[] args) {
		// Example 1
		String s1 = "codeleet";
		int[] indices1 = { 4, 5, 6, 7, 0, 2, 1, 3 };
		System.out.printf("\n👉 Example 1:\nInput: s = \"%s\", indices = %s\nOutput: \"%s\"\n\n",
				s1, Arrays.toString(indices1), restoreString(s1, indices1));

		System.out.println("------------------------");

		// Example 2
		String s2 = "abc";
		int[] indices2 = { 0, 1, 2 };
		System.out.printf("\n👉 Example 2:\nInput: s = \"%s\", indices = %s\nOutput: \"%s\"\n",
				s2, Arrays.toString(indices2), restoreString(s2, indices2));
	}
}