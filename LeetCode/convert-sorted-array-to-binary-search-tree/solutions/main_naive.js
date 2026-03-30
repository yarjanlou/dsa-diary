function TreeNode(val, left, right) {
  this.val = val === undefined ? 0 : val;
  this.left = left === undefined ? null : left;
  this.right = right === undefined ? null : right;
}

const sortedArrayToBST = function (nums) {
  const rootIndex = Math.floor(nums.length / 2);
  const root = new TreeNode(nums[rootIndex]);

  if (rootIndex - 1 >= 0)
    root.left = sortedArrayToBST(nums.slice(0, rootIndex));
  if (rootIndex + 1 <= nums.length)
    root.right = sortedArrayToBST(nums.slice(rootIndex + 1));

  return root;
};