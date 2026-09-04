package main

import (
	"fmt"
	"math"
)


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if calHeight(root) == -1 {
		return  false
	}

	return true
}

func calHeight(root *TreeNode) float64 {
    fmt.Println(root)
	if root == nil {
		return 0
	}

	var leftHeight float64 = calHeight(root.Left)
	var rightHeight float64 = calHeight(root.Right)


	if math.Abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return math.Max(leftHeight, rightHeight) + 1
}

func main() {
	isBalanced(&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}})
}
