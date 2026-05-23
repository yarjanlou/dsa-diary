# 404. Sum of Left Leaves

**Difficulty:** Easy <br>
**Status:** Solved

---

## 🧩 Problem Description

Given the `root` of a binary tree, return the **sum of all left leaves**.

A **leaf** is a node with no children. A **left leaf** is a leaf that is the left child of its parent node.

---

## 💡 Examples

### Example 1

```text
Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves: 9 and 15. Their sum is 24.
```

### Example 2

```text
Input: root = [1]
Output: 0
Explanation: There are no left leaves.
```

---

## 📐 Constraints

* The number of nodes in the tree is in the range `[1, 1000]`.
* `-1000 <= Node.val <= 1000`

---

## ✅ Notes

* This is a classic **binary tree traversal** problem.
* The key insight is that a node alone cannot determine if it is a left leaf — its **parent** must identify it as a left child.
* Both **recursive DFS** and **iterative BFS** strategies are valid approaches.