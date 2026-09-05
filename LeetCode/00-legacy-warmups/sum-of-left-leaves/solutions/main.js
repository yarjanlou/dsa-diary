/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

/**
 * @param {TreeNode} root
 * @return {number}
 */

var sumOfLeftLeaves = function (root) {
  if (root === null) return 0;

  const left = root.left;
  const right = root.right;

  if (left === null && right === null) return 0;

  if (left?.left === null && left?.right === null)
    return left.val + sumOfLeftLeaves(right);
  else {
    return sumOfLeftLeaves(left) + sumOfLeftLeaves(right);
  }
};
