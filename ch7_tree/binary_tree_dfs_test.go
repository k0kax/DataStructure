package ch7_tree

import (
	. "datastructure/pkg"
	"fmt"
	"testing"
)

func TestPreInPostOrderTraversal(t *testing.T) {
	//初始化二叉树
	root := SliceToTree([]any{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println("\n初始化二叉树: ")
	PrintTree(root)

	//前序遍历
	nums = nil
	preOrder(root)
	fmt.Println("\n前序遍历的节点打印序列 =", nums)
	//中序遍历
	nums = nil
	inOrder(root)
	fmt.Println("\n中序遍历的节点打印序列 =", nums)
	//后序遍历
	nums = nil
	postOrder(root)
	fmt.Println("\n后序遍历的节点打印序列 =", nums)
}
