# 110. Balanced Binary Tree

Determine whether a binary tree is height-balanced.

[LeetCode Problem 110](https://leetcode.com/problems/balanced-binary-tree/)

| Field | Value |
|---|---|
| LeetCode # | 110 |
| Difficulty | Easy |
| Primary concept | Binary Tree DFS |
| Secondary concepts | — |
| Status | ✅ Solved |

## Problem

A binary tree is considered balanced if, for every node, the difference between the heights of its left and right subtrees is at most 1.

## Examples

### Example 1

```text
Input:  root = [3,9,20,null,null,15,7]
Output: true
```

### Example 2

```text
Input:  root = [1,2,2,3,3,null,null,4,4]
Output: false
```

### Example 3

```text
Input:  root = []
Output: true
```

## Constraints

- The number of nodes is in the range `[0, 5000]`
- `-10^4 <= Node.val <= 10^4`
