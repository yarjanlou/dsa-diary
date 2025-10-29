# 28. Find the Index of the First Occurrence in a String

**Difficulty:** Easy <br>
**Status:** Solved

---

## 🧩 Problem Description

Given two strings, `needle` and `haystack`, return the **index of the first occurrence** of `needle` in `haystack`.
Return `-1` if `needle` is not part of `haystack`.

---

## 💡 Examples

### Example 1

```text
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6. 
The first occurrence is at index 0.
```

### Example 2

```text
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" does not occur in "leetcode".
```

---

## 📐 Constraints

* `1 <= haystack.length, needle.length <= 10^4`
* `haystack` and `needle` consist of **only lowercase English letters**.


## ✅ Notes

* This is a classic **substring search problem**.
* Optimal solutions can be achieved using algorithms like **KMP**, **Rabin-Karp**, or built-in methods.

