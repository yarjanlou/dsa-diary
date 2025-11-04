# 🧩 Problem 26: Remove Duplicates from Sorted Array

**Difficulty:** Easy
**Topics:** Array, Two Pointers

---

## 🧠 Description

Given an integer array `nums` **sorted in non-decreasing order**, remove the duplicates **in-place** such that each unique element appears only once.
The relative order of the elements should remain the same.

After removing duplicates, return the number of unique elements `k`.

The first `k` elements of `nums` should contain the unique numbers in sorted order.
The remaining elements beyond index `k - 1` can be ignored.

---

## 🧪 Example

### Example 1

```text
Input:  nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: 
Your function should return k = 2, 
with the first two elements of nums being [1, 2].
It does not matter what values are beyond the returned k.
```

### Example 2

```text
Input:  nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: 
Your function should return k = 5,
with the first five elements of nums being [0,1,2,3,4].
The rest can be ignored.
```

---

## 🧩 Custom Judge

Your implementation will be tested with the following logic:

```java
int[] nums = [...];           // Input array
int[] expectedNums = [...];   // Expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```

If all assertions pass, your solution will be accepted ✅.

---

## ⚙️ Constraints

* `1 <= nums.length <= 3 * 10^4`
* `-100 <= nums[i] <= 100`
* `nums` is sorted in **non-decreasing order**

---

