# 367. Valid Perfect Square

Determine whether a positive integer is a perfect square, without using built-in square root functions.

[LeetCode Problem 367](https://leetcode.com/problems/valid-perfect-square/)

| Field | Value |
|---|---|
| LeetCode # | 367 |
| Difficulty | Easy |
| Primary concept | Binary Search |
| Secondary concepts | Math (Newton's Method) |
| Status | ✅ Solved |

## Problem

A number is a perfect square if it can be expressed as `k * k` for some integer `k`. Given a positive integer `num`, return `true` if it is a perfect square, and `false` otherwise.

Built-in functions such as `sqrt` must not be used — the solution should rely on pure logic or math techniques.

## Examples

### Example 1

```text
Input: num = 16
Output: true
Explanation: 4 * 4 = 16, so it is a perfect square.
```

### Example 2

```text
Input: num = 14
Output: false
Explanation: No integer k exists such that k * k = 14.
```

## Constraints

- `1 <= num <= 2^31 - 1`
