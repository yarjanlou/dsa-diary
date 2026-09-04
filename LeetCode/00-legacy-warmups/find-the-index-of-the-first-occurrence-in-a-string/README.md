# 28. Find the Index of the First Occurrence in a String

Return the index of the first occurrence of `needle` in `haystack`, or `-1` if it does not occur.

[LeetCode Problem 28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/)

| Field | Value |
|---|---|
| LeetCode # | 28 |
| Difficulty | Easy |
| Primary concept | String Matching |
| Secondary concepts | — |
| Status | ✅ Solved |

## Problem

Given two strings, `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`. Return `-1` if `needle` is not part of `haystack`.

## Examples

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

## Constraints

- `1 <= haystack.length, needle.length <= 10^4`
- `haystack` and `needle` consist of only lowercase English letters.

## Notes

This is a classic substring search problem. Optimal solutions can be achieved using algorithms like KMP, Rabin-Karp, or built-in methods.
