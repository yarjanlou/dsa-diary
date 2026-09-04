# 387. First Unique Character in a String

Find the first non-repeating character in a string and return its index.

[LeetCode Problem 387](https://leetcode.com/problems/first-unique-character-in-a-string/)

| Field | Value |
|---|---|
| LeetCode # | 387 |
| Difficulty | Easy |
| Primary concept | Hashing & Frequency Counting |
| Secondary concepts | — |
| Status | ✅ Solved |

## Problem

Given a string `s`, find the first non-repeating character in it and return its index. If it does not exist, return `-1`.

## Examples

### Example 1

```text
Input:  s = "leetcode"
Output: 0
Explanation: The character 'l' at index 0 is the first character that does not occur at any other index.
```

### Example 2

```text
Input:  s = "loveleetcode"
Output: 2
```

### Example 3

```text
Input:  s = "aabb"
Output: -1
```

## Constraints

- `1 <= s.length <= 10^5`
- `s` consists of only lowercase English letters.
