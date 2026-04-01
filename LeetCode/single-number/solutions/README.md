# 📌 Single Number Problem Solutions

This document outlines two distinct solutions to the **LeetCode "Single Number"** problem, which requires identifying the number that appears **exactly once** in an array where all other numbers appear **twice**.

---

## 🧠 Problem Overview

Given an array of integers where every element appears **exactly twice** except for one element that appears **once**, return the unique number.

**Example:**
```js
Input: [2, 2, 1]
Output: 1
```

---

## 🔍 Solution 1: Using a Hash Map (Object)

### ✅ Approach
- Use an object to count the frequency of each number.
- Iterate through the array, updating the count for each number.
- Finally, iterate through the object to find the number with a count of `1`.


### 📊 Time/Space Complexity
- **Time Complexity**: `O(n)` — Two passes over the array.
- **Space Complexity**: `O(n)` — The object stores up to `n` unique numbers.

### 📝 Notes
- **Readability**: Easy to understand for beginners.
- **Drawback**: Uses extra memory for the object.

---

## 🔁 Solution 2: Using XOR Bitwise Operation (Optimal)

### ✅ Approach
- Use the **XOR bitwise operation**:
  - `a ^ a = 0` (any number XOR itself is 0)
  - `a ^ 0 = a` (any number XOR 0 is the number itself)
- XOR all the numbers in the array. The duplicates will cancel out, leaving the unique number.


### 📊 Time/Space Complexity
- **Time Complexity**: `O(n)` — One pass over the array.
- **Space Complexity**: `O(1)` — No extra memory used.

### 📝 Notes
- **Efficiency**: Fast and uses no extra memory.
- **Best Practice**: Preferred for large inputs or performance-critical code.

---

## 📌 Key Takeaways

| Feature               | Hash Map Solution       | XOR Solution            |
|----------------------|-------------------------|-------------------------|
| **Time Complexity**  | `O(n)`                  | `O(n)`                  |
| **Space Complexity** | `O(n)`                  | `O(1)`                  |
| **Readability**      | ✅ Easy to understand   | ✅ Easy to understand   |
| **Performance**      | ⚠️ Moderate             | ✅ Optimal              |
| **Recommended Use**  | Small arrays or learning | All cases              |

---

## 🧪 Example Usage

```js
console.log(singleNumber([2, 2, 1])); // Output: 1
console.log(singleNumber([4, 1, 2, 1, 2])); // Output: 4
```

---

## 📚 References

- [LeetCode - Single Number Problem](https://leetcode.com/problems/single-number/)
- [Bitwise XOR Operation](https://en.wikipedia.org/wiki/Exclusive_or)