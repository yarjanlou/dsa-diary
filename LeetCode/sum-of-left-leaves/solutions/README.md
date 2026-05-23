## 🧩 Binary Tree Traversal (`sumOfLeftLeaves`)

### 📖 Overview

This project implements the **Sum of Left Leaves** problem — identifying and summing every leaf node that appears as a **left child** in a binary tree.

It contains two distinct recursive DFS strategies written in **JavaScript** and **Go**.

---

### 🚀 Features

- ✅ Simple and clear implementation
- ✅ Solutions in Go and JavaScript
- ✅ No third-party libraries
- 🧠 Two recursive DFS approaches: parent-aware and flag-based

---

### 🧮 Algorithm 1: Parent-Aware Recursion (JavaScript)

**Description:**
At each node, the algorithm inspects the **left child directly** to determine whether it is a leaf (both children are `null`). If so, its value is added to the result of recursing into the right subtree. Otherwise, the algorithm recurses into both subtrees independently.

This approach delegates the leaf-detection responsibility to the **parent node**, allowing a clean recursive structure without passing extra state down the call stack.

**Base cases:**
- `root === null` → return `0`
- `root` is itself a leaf (no children) → return `0` (it has no parent context here, so it cannot be a left leaf)

**Time Complexity:** `O(n)`
**Space Complexity:** `O(h)`

Where:

- `n` → number of nodes in the tree
- `h` → height of the tree (call stack depth)

---

### 🧮 Algorithm 2: DFS with `isLeft` Flag (Go)

**Description:**
A helper function `dfs(node, isLeft)` propagates a boolean flag indicating whether the current node was reached as a **left child**. When a leaf is encountered and `isLeft` is `true`, its value is returned. Otherwise, the recursion continues into both children, passing `true` for the left branch and `false` for the right.

This approach is arguably more explicit — the "left" context travels with each node through the recursion, making the logic straightforward to trace.

**Base cases:**
- `root == nil` → return `0`
- `root` is a leaf and `isLeft == true` → return `root.Val`

**Time Complexity:** `O(n)`
**Space Complexity:** `O(h)`

Where:

- `n` → number of nodes in the tree
- `h` → height of the tree (call stack depth)

---

### 📊 Comparison

Both algorithms share the same asymptotic complexity and visit every node exactly once. The difference is stylistic:

| Approach | Language | Left-context tracking |
|---|---|---|
| Parent-aware recursion | JavaScript | Parent inspects its left child |
| `isLeft` flag propagation | Go | Flag passed down the call stack |

- **Parent-aware** → slightly more implicit; avoids extra function arguments
- **Flag-based** → more explicit; easier to extend (e.g. summing right leaves instead)