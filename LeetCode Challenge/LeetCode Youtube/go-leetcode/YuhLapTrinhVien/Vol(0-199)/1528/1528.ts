/**
 * Hàm restoreString nhận vào một chuỗi s và một mảng indices,
 * trả về chuỗi đã được sắp xếp lại theo indices.
 */
function restoreString(s: string, indices: number[]): string {
  const result: string[] = new Array(s.length);

  // Đặt mỗi ký tự vào đúng vị trí của nó
  for (let i = 0; i < s.length; i++) {
    result[indices[i]] = s[i];
  }

  return result.join("");
}

// Hàm main để kiểm tra kết quả
function main(): void {
  // Ví dụ 1
  const s1: string = "codeleet";
  const indices1: number[] = [4, 5, 6, 7, 0, 2, 1, 3];
  console.log(
    `Example 1:\nInput: s = "${s1}", indices = [${indices1}]\nOutput: "${restoreString(
      s1,
      indices1
    )}"`
  );

  // Ví dụ 2
  const s2: string = "abc";
  const indices2: number[] = [0, 1, 2];
  console.log(
    `Example 2:\nInput: s = "${s2}", indices = [${indices2}]\nOutput: "${restoreString(
      s2,
      indices2
    )}"`
  );
}

// Gọi hàm main
main();
