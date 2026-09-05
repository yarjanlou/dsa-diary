# 66. Plus One

Add one to a large integer represented as an array of digits.

[LeetCode Problem 66](https://leetcode.com/problems/plus-one/)

| Field | Value |
|---|---|
| LeetCode # | 66 |
| Difficulty | Easy |
| Primary concept | Array Simulation |
| Secondary concepts | Math |
| Status | ✅ Solved |

## Problem

You are given a large integer represented as an integer array `digits`. Each `digits[i]` is the `i`th digit of the integer, ordered from most significant to least significant. The integer does not contain any leading zeros.

Add one to the integer and return the resulting array of digits.

## Examples

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

## Constraints

- `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`
- `digits` does not contain leading zeros
