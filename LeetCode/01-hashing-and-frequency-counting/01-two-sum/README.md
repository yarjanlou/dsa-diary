# 1. Two Sum

Given an array of integers, return the indices of the two numbers that add up to a specific target.

[LeetCode Problem 1](https://leetcode.com/problems/two-sum/)

| Field | Value |
|---|---|
| LeetCode # | 1 |
| Difficulty | Easy |
| Primary concept | Hashing & Frequency Counting |
| Secondary concepts | — |
| Progression order | 3 / 5 |
| Status | ✅ Solved |

## Problem

Given an array of integers `nums` and an integer `target`, return the indices of the two numbers such that they add up to `target`.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## Examples

### Example 1

```text
Input:  nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
Explanation: nums[0] + nums[1] == 9, so we return [0, 1].
```

### Example 2

```text
Input:  nums = [3, 2, 4], target = 6
Output: [1, 2]
```

### Example 3

```text
Input:  nums = [3, 3], target = 6
Output: [0, 1]
```

## Constraints

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- Only one valid answer exists.
