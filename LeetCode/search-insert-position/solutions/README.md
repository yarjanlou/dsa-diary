# 🧩 Problem 35: Search Insert Position

## 🧠 Solution Overview

The solution uses **binary search** to find the insertion position of the target.

Instead of searching only for an exact match, we find the **first index where `nums[index] >= target`**.

The search range is narrowed until `left` and `right` meet. At that point, `left` is the correct insertion index.

---

## ⚙️ How It Works

1. Initialize `left = 0` and `right = nums.length`.
2. While `left < right`:

   * Calculate the middle index.
   * If `nums[mid] < target`, the target must be to the right of `mid`, so set `left = mid + 1`.
   * Otherwise, `mid` could be the insertion position, so set `right = mid`.
3. Return `left`.

This also handles cases where the target is:

* Smaller than every element → returns `0`
* Equal to an existing element → returns its index
* Between two elements → returns the index where it should be inserted
* Larger than every element → returns `nums.length`

---

## ⏱️ Time Complexity

**O(log n)**

Binary search eliminates approximately half of the remaining search space on each iteration.

---

## 💾 Space Complexity

**O(1)**

The algorithm only uses a constant number of variables regardless of the input size.

---

## 🚀 Key Points

* Uses binary search to satisfy the required **O(log n)** runtime.
* Finds the first position where `nums[index] >= target`.
* The final value of `left` is the correct insertion index.
* No additional data structures are required.
