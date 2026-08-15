func searchInsert(nums []int, target int) int {
    
  var start int = 0
  var end int = len(nums)
  var mid int = len(nums) / 2

  for start < mid {
    if target == nums[mid] {
      return mid
    }

    if target > nums[mid] {
      start = mid
    }

    if target < nums[mid] {
      end = mid
    }

    mid = (start + end) / 2
  }

  if target == nums[mid] || target < nums[mid] {
    return mid
  }

  return end
}