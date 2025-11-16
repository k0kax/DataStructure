package ch7_tree

import (
	"fmt"
	"testing"
)

func TestAVLTree(t *testing.T) {
	//初始化空AVL树
	tree := newAVLTree()
	//插入节点
	testInsert(tree, 1)
	testInsert(tree, 2)
	testInsert(tree, 3)
	testInsert(tree, 4)
	testInsert(tree, 5)
	testInsert(tree, 8)
	testInsert(tree, 7)
	testInsert(tree, 9)
	testInsert(tree, 10)
	testInsert(tree, 6)

	//插入重复节点
	fmt.Println("\n尝试插入重复节点 7:")
	testInsert(tree, 7)

	//删除节点
	testRemove(tree, 8)
	testRemove(tree, 5)
	testRemove(tree, 4)

	//查找节点
	node := tree.search(7)
	fmt.Printf("\n查找到的节点对象为 %#v ，节点值 = %d \n", node, node.Val)
}

// 测试插入节点
func testInsert(tree *aVlTree, val int) {
	tree.insert(val)
	fmt.Printf("\n插入节点%d后，AVL树为\n", val)
	tree.print()
}

// 测试删除节点
func testRemove(tree *aVlTree, val int) {
	tree.remove(val)
	fmt.Printf("\n删除节点%d后，AVL树为\n", val)
	tree.print()
}
