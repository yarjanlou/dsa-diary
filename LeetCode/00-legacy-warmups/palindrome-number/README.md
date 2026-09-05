# 9. Palindrome Number

Determine whether an integer reads the same forward and backward.

[LeetCode Problem 9](https://leetcode.com/problems/palindrome-number/)

| Field | Value |
|---|---|
| LeetCode # | 9 |
| Difficulty | Easy |
| Primary concept | Math (Digit Manipulation) |
| Secondary concepts | — |
| Status | ✅ Solved |

## Problem

Given an integer `x`, return `true` if `x` is a palindrome, and `false` otherwise. A negative number is never a palindrome, because of its minus sign.

## Examples

### Example 1

```text
Input:  x = 121
Output: true
Explanation: 121 reads the same backward.
```

### Example 2

```text
Input:  x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, 121-. Not the same.
```

### Example 3

```text
Input:  x = 10
Output: false
Explanation: Reads 01 backward, which is not 10.
```

## Constraints

- `-2^31 <= x <= 2^31 - 1`
