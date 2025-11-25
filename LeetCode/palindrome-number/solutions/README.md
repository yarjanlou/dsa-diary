## ⚙️ How It Works

1. If the number is negative, return `false` immediately — negative numbers cannot be palindromes.
2. Store the original value of `x` in `number` so you can compare it later.
2. Initialize `reverseNumber` as `0`.
3. While `x` is greater than `0`:

   * Extract the last digit using `x % 10`.
   * Append it to reverseNumber by doing:
     `reverseNumber = reverseNumber * 10 + (x % 10)`
   * Remove the last digit from x using integer division:
     `x = Math.floor(x / 10)` (or `x /= 10` in Go).
4. After the loop, compare the reversed number with the original:
   If `reverseNumber === number`, it's a palindrome.

---

## ⏱️ Time and Space Complexity

| Complexity Type | Explanation                                                                    |
| --------------- | ------------------------------------------------------------------------------ |
| **Time**        | `O(log₁₀ n)` — You process each digit once while reversing the number.         |                               |
| **Space**       | `O(1)` — Only a few variables are used; no extra data structures required.     |
