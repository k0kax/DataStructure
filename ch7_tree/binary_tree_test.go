package ch7_tree

import (
	. "datastructure/pkg"
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	//初始化节点
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)

	//构建节点间的引用（指针）
	n1.Left = n2
	n1.Right = n3
	n2.Right = n5
	n2.Left = n4
	fmt.Println("初始化二叉树")
	PrintTree(n1)

	//插入节点
	p := NewTreeNode(0)
	n1.Left = p
	p.Left = n2
	fmt.Println("插入节点P后")
	PrintTree(n1)

	//删除节点
	n1.Left = n2
	fmt.Println("删除节点P后")
	PrintTree(n1)
}
