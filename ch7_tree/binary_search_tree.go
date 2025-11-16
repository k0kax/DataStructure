package ch7_tree

import (
	. "datastructure/pkg"
)

type binarySearchTree struct {
	root *TreeNode
}

func newBinarySearchTree() *binarySearchTree {
	bst := &binarySearchTree{}
	//初始化空树
	bst.root = nil
	return bst
}

// 获取根节点
func (bst *binarySearchTree) getRoot() *TreeNode {
	return bst.root
}

// 查找节点
func (bst *binarySearchTree) search(num int) *TreeNode {
	node := bst.root
	//循环查找，越过叶节点后跳出
	for node != nil {
		if node.Val.(int) < num {
			//目标节点在cur的右子树上
			node = node.Right
		} else if node.Val.(int) > num {
			//目标节点在cur的左子树
			node = node.Left
		} else {
			//未找到目标节点，跳出循环
			break
		}
	}
	//返回目标节点
	return node
}

// 插入节点
func (bst *binarySearchTree) insert(num int) {
	cur := bst.root
	//若树为空，则初始化根节点
	if cur == nil {
		bst.root = NewTreeNode(num)
		return
	}
	//待插入节点之前的节点位置
	var pre *TreeNode = nil
	//循环查找，越过叶节点后跳出
	for cur != nil {
		if cur.Val == num {
			return
		}
		pre = cur
		if cur.Val.(int) < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	//插入节点
	node := NewTreeNode(num)
	if pre.Val.(int) < num {
		pre.Right = node
	} else {
		pre.Left = node
	}
}

// 删除节点
func (bst *binarySearchTree) remove(num int) {
	cur := bst.root
	//若树为空，则提前返回
	if cur == nil {
		return
	}

	//待删除节点之前的节点位置
	var pre *TreeNode = nil
	//循环查找，越过叶节点后跳出
	for cur != nil {
		if cur.Val == num {
			break
		}
		pre = cur
		if cur.Val.(int) < num {
			//待删除节点在右子树
			cur = cur.Right
		} else {
			//待删除节点在左子树
			cur = cur.Left
		}
	}

	//无待删除节点则直接返回
	if cur == nil {
		return
	}

	//子节点数为0或1
	if cur.Left == nil || cur.Right == nil {
		var child *TreeNode = nil
		//取出待删除的节点的子节点
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}

		//删除节点cur
		if cur != bst.root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			//若删除节点为根节点，则重新制定节点
			bst.root = child
		}

		//子节点数为2
	} else {
		//获取中序遍历中待删除节点cur的下一个节点
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		//递归删除节点tmp
		bst.remove(tmp.Val.(int))
		//用tmp覆盖cur
		cur.Val = tmp.Val
	}
}

// 打印搜索二叉树
func (bst *binarySearchTree) print() {
	PrintTree(bst.root)
}
