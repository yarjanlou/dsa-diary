# 387. First Unique Character in a String

## 🧠 Solution Overview

To solve this problem, we use a **two-pass approach with a hash map**:

1. **First pass:** Count the frequency of each character.
2. **Second pass:** Find the first character with frequency `1` and return its index.

This ensures we efficiently identify the first non-repeating character.

---

## ⚙️ How It Works

* First loop:

  * Count how many times each character appears.
* Second loop:

  * Return the index of the first character that appears exactly once.
* If no such character exists, return `-1`.

---

## ⏱️ Time Complexity

* **O(n)**

  * First loop: `O(n)`
  * Second loop: `O(n)`
  * Total: `O(n)`

---

## 💾 Space Complexity

* **O(1)**

  * The map stores at most 26 lowercase English letters.
  * Even though we use a map, the space is constant due to input constraints.

---

## ✅ Key Points

* Efficient and clean solution using hashing.
* Avoids nested loops (which would be `O(n^2)`).
* Works well for large inputs up to `10^5`.

---

## 🚀 Summary

This is a classic example of trading **space for time**:

* We use extra memory (hash map) to achieve linear time performance.
* A very common pattern in string and array problems.

