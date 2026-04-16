# Reverse String (In-Place)

## 📌 Description

This function reverses a string represented as a slice of bytes **in-place** without using any additional memory.

It uses a **two-pointer approach**, swapping elements from the beginning and the end of the slice until reaching the middle.

---

## ⚙️ How It Works

* Initialize the size of the slice
* Iterate from the start to the middle of the slice
* Swap:

  * First element ↔ Last element
  * Second element ↔ Second-last element
  * Continue until the middle is reached

This ensures the array is reversed **without extra space**.

---

## ⏱️ Time Complexity

**O(n)**

* The loop runs `n/2` times
* Constant work per iteration
* Overall linear time complexity

---

## 💾 Space Complexity

**O(1)**

* No additional memory is used
* Reversal is done **in-place**

---

## ✅ Key Advantages

* Memory efficient (no extra allocation)
* Simple and readable
* Optimal time complexity

---

## 🧪 Example

```
Input:  ['h','e','l','l','o']
Output: ['o','l','l','e','h']
```
