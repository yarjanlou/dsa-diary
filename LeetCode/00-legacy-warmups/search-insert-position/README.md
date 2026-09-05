# 35. Search Insert Position

Return the index of `target` in a sorted array, or the index where it would be inserted to keep the array sorted.

[LeetCode Problem 35](https://leetcode.com/problems/search-insert-position/)

| Field | Value |
|---|---|
| LeetCode # | 35 |
| Difficulty | Easy |
| Primary concept | Binary Search |
| Secondary concepts | — |
| Status | ✅ Solved |

## Problem

Given a sorted array of distinct integers `nums` and a target value, return the index if the target is found. If not, return the index where it would be inserted, in order to keep the array sorted.

The algorithm must run in `O(log n)` time.

## Examples

### Example 1

```text
Input: nums = [1,3,5,6], target = 5
Output: 2
Explanation: 5 is found at index 2.
```

### Example 2

```text
Input: nums = [1,3,5,6], target = 2
Output: 1
Explanation: 2 would be inserted at index 1 to maintain sorted order.
```

### Example 3

```text
Input: nums = [1,3,5,6], target = 7
Output: 4
Explanation: 7 would be inserted at the end (index 4).
```

### Example 4

```text
Input: nums = [1,3,5,6], target = 0
Output: 0
Explanation: 0 would be inserted at the beginning (index 0).
```

## Constraints

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` contains distinct values sorted in ascending order
- `-10^4 <= target <= 10^4`

## Notes

- The returned index is always between `0` and `nums.length` (inclusive). If the target is larger than all elements, it returns `nums.length`.
- The binary search implementation returns the leftmost insertion point, which is suitable for arrays with distinct values.
