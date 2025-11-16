package ch7_tree

import (
	. "datastructure/pkg"
)

var nums []any

// 前序遍历
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	//访问优先级：根节点->左子树->右子树
	nums = append(nums, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

// 中序遍历
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	//访问优先级：左子树->根节点->右子树
	preOrder(node.Left)
	nums = append(nums, node.Val)
	preOrder(node.Right)
}

// 后序遍历
func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	//访问优先级：左子树->右子树->根节点
	preOrder(node.Left)
	preOrder(node.Right)
	nums = append(nums, node.Val)
}
