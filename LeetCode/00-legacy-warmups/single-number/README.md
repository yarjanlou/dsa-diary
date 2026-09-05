# 136. Single Number

Find the element that appears once in an array where every other element appears twice.

[LeetCode Problem 136](https://leetcode.com/problems/single-number/)

| Field | Value |
|---|---|
| LeetCode # | 136 |
| Difficulty | Easy |
| Primary concept | Bit Manipulation (XOR) |
| Secondary concepts | Hashing & Frequency Counting |
| Status | ✅ Solved |

## Problem

Given a non-empty array of integers `nums`, every element appears twice except for one. Find that single one.

The solution must run in linear time and use only constant extra space.

## Examples

### Example 1

```text
Input: nums = [2, 2, 1]
Output: 1
```

### Example 2

```text
Input: nums = [4, 1, 2, 1, 2]
Output: 4
```

### Example 3

```text
Input: nums = [1]
Output: 1
```

## Constraints

- `1 <= nums.length <= 3 * 10^4`
- `-3 * 10^4 <= nums[i] <= 3 * 10^4`
- Each element in the array appears twice except for one element which appears only once.
