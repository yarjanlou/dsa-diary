# 🧩 Problem 35: Search Insert Position

**Difficulty:** Easy  
**Topics:** Array, Binary Search

---

## 🧠 Description

Given a sorted array of distinct integers nums and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order to keep the array sorted.

You must write an algorithm with O(log n) runtime complexity.

---

## 🧪 Examples

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

---

## 🧭 Approaches

1. Linear scan (Brute force)
   - Iterate through nums and return the first index i where nums[i] >= target.
   - Time: O(n), Space: O(1).
   - Simple but does not meet the preferred O(log n) requirement for large inputs.

2. Binary search (Recommended)
   - Use a classic binary search to find the lowest index where nums[index] >= target.
   - Maintain left and right pointers and narrow the search range until left == right. The final position for insertion is `left`.
   - Time: O(log n), Space: O(1).

Pseudo-steps for binary search:
- Initialize left = 0, right = nums.length
- While left < right:
  - mid = left + (right - left) // 2
  - if nums[mid] < target: left = mid + 1
  - else: right = mid
- Return left

---

## ⚙️ Constraints

- 1 <= nums.length <= 10^4
- -10^4 <= nums[i] <= 10^4
- nums contains distinct values sorted in ascending order
- -10^4 <= target <= 10^4

---

## ✅ Notes

- The returned index is always between 0 and nums.length (inclusive). If the target is larger than all elements, it returns nums.length.
- The binary search variant above returns the leftmost insertion point, which is suitable for arrays with distinct values.

