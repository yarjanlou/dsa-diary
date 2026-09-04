package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nums []int = []int{-10, -3, 0, 5, 9}

func main() {
	fmt.Println(sortedArrayToBST(nums))
}

func sortedArrayToBST(nums []int) *TreeNode {
	var treeResult TreeNode

	// 1. find root
	var rootIndex int = len(nums) / 2

	treeResult.Val = nums[rootIndex]

	if rootIndex+1 < len(nums) {
		treeResult.Right = sortedArrayToBST(nums[rootIndex+1:])

	}

	if rootIndex-1 > 0 {

		treeResult.Left = sortedArrayToBST(nums[0 : rootIndex])
	}

	return &treeResult
}
