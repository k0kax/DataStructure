package ch7_tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := newBinarySearchTree()
	nums := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	//该数列会生成一个完美二叉树
	for _, num := range nums {
		bst.insert(num)
	}
	fmt.Println("\n初始化的二叉树：")
	bst.print()

	//获取根节点
	node := bst.getRoot()
	fmt.Println("\n二叉搜索树的根节点为：", node.Val)

	//查找节点
	node = bst.search(7)
	fmt.Println("\n二叉搜索树的查到的节点对象为为：", node, ",节点值为 =", node.Val)

	//插入节点
	bst.insert(16)
	fmt.Println("\n二叉搜索树插入节点后 16 为：")
	bst.print()

	//删除节点
	bst.remove(1)
	fmt.Println("\n删除叶节点1后的二叉树为：")
	bst.print()
	bst.remove(2)
	fmt.Println("\n删除节点2后的二叉树为：")
	bst.print()
	bst.remove(4)
	fmt.Println("\n删除节点4后的二叉树为：")
	bst.print()
	bst.remove(8)
	fmt.Println("\n删除根节点8后的二叉树为：")
	bst.print()

}
