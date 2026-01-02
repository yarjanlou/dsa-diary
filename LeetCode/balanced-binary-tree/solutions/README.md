# Balanced Binary Tree — Explanation


## Key Insight

A simple idea would be:

* compute the height of every subtree
* check balance at each node

---

## Efficient Idea (Bottom-Up)

We compute the height **while** checking balance.

For each node:

1. Recursively get the height of the left subtree
2. Recursively get the height of the right subtree
3. If either subtree is already unbalanced → propagate a special signal (for example, `-1`) upward
4. If the height difference is greater than 1 → mark as unbalanced
5. Otherwise, return the node’s height:

```
height = max(left, right) + 1
```

### Why this works

* Each node is processed once
* As soon as we detect imbalance, we **stop doing unnecessary work**
* The “signal” bubbles up and prevents deeper recursion from mattering

This converts the entire problem into a **single depth-first traversal**.

---

## Time & Space Complexity

### Time Complexity — **O(n)**

* Each node is visited exactly once
* Balance and height are computed in the same pass

Where `n` = number of nodes.

---

### Space Complexity — **O(h)**

* `h` is the height of the tree
* space comes from the recursion stack

Worst case (skewed tree): **O(n)**
Best / average case (balanced tree): **O(log n)**

---

## Summary

| Concept   | Explanation                                    |
| --------- | ---------------------------------------------- |
| Strategy  | DFS (bottom-up)                                |
| Key Trick | Return height AND imbalance signal in one pass |
| Time      | **O(n)**                                       |
| Space     | **O(h)** recursion                             |

