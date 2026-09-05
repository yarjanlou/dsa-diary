# 108. Convert Sorted Array to Binary Search Tree

Convert a sorted array into a height-balanced binary search tree.

[LeetCode Problem 108](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)

| Field | Value |
|---|---|
| LeetCode # | 108 |
| Difficulty | Easy |
| Primary concept | Divide & Conquer |
| Secondary concepts | — |
| Status | ✅ Solved |

## Problem

Given an integer array `nums` where the elements are sorted in ascending order, convert it into a height-balanced Binary Search Tree (BST).

## Examples

### Example 1

```text
Input:  nums = [-10, -3, 0, 5, 9]
Output: [0, -3, 9, -10, null, 5]
```

### Example 2

```text
Input:  nums = [1, 3]
Output: [3, 1]
```

## Constraints

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` is sorted in strictly increasing order
