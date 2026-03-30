# 📘 Convert Sorted Array to Binary Search Tree

## 💡 Approach

To build a **height-balanced BST**, use a **divide and conquer** strategy:

* Pick the **middle element** as the root
* Recursively construct:

  * Left subtree from the left half
  * Right subtree from the right half

This guarantees the tree remains balanced.

---

## 🧩 Algorithm

1. Find the middle index of the array
2. Create a node with that value
3. Recursively repeat for:

   * Left subarray
   * Right subarray
4. Stop when the range is empty

---

## ⏱️ Time Complexity

### **O(n)**

* Each element is used exactly once to create a node
* Total operations scale linearly with input size

---

## 🧠 Space Complexity

### **O(log n)** (Recursion Stack)

* Depth of recursion equals the height of the tree
* Since the tree is balanced, height ≈ `log n`

### ⚠️ Implementation Note

* Using array slicing may introduce extra overhead depending on the language
* More optimal solutions pass indices instead of creating subarrays

---

## 📚 Summary

* Use **middle element** to ensure balance
* Apply **recursive divide & conquer**
* Efficient:

  * **Time:** O(n)
  * **Space:** O(log n)

---
