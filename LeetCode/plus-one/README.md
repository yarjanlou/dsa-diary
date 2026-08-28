# 🧩 Problem 66: Plus One

**Difficulty:** Easy  
**Topics:** Array, Math

---

## 🧠 Description

You are given a large integer represented as an integer array `digits`. Each
`digits[i]` is the `i`th digit of the integer, ordered from most significant to
least significant.

The integer does not contain any leading zeros. Add one to the integer and
return the resulting array of digits.

---

## 🧪 Examples

### Example 1

```text
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents 123. Adding one produces 124.
```

### Example 2

```text
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents 4321. Adding one produces 4322.
```

### Example 3

```text
Input: digits = [9]
Output: [1,0]
Explanation: The array represents 9. Adding one produces 10.
```

---

## ⚙️ Constraints

- `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`
- `digits` does not contain leading zeros

---

## 💡 Approach

Process the digits from right to left, just as in manual addition:

1. If the current digit is less than `9`, increment it and return the array.
2. If the current digit is `9`, change it to `0` and carry one to the next
   digit on the left.
3. If every digit was `9`, add a new leading `1`.

For example, `[1,9,9]` becomes `[2,0,0]`, while `[9,9]` becomes `[1,0,0]`.

---

## ⏱️ Complexity

- **Time:** O(n) in the worst case, when every digit is `9`
- **Space:** O(1) auxiliary space; an additional digit is required only when
  the input contains all `9`s
