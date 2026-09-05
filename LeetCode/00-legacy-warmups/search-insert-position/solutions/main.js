/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */

const searchInsert = (nums, target) => {
  let left = 0;
  let right = nums.length - 1;

  while (left <= right) {
    root = Math.floor((left + right) / 2);

    if (nums[root] === target) return root;
    else if (target > nums[root]) left = root + 1;
    else if (target < nums[root]) right = root - 1;
  }

  return left;
};
