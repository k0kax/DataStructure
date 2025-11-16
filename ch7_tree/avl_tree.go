package ch7_tree

import (
	. "datastructure/pkg"
)

// 平衡二叉树AVL
type aVlTree struct {
	//根节点
	root *TreeNode
}

// 新建AVL树
func newAVLTree() *aVlTree {
	return &aVlTree{root: nil}
}

// 获取节点的高度
func (t *aVlTree) height(node *TreeNode) int {
	//空节点高度为-1，叶节点高度为0
	if node != nil {
		return node.Height
	}
	return -1
}

// 更新节点高度
func (t *aVlTree) updateHeight(node *TreeNode) {
	if node == nil {
		return
	}
	lh := t.height(node.Left)
	rh := t.height(node.Right)
	//节点高度等于最高子树高度 +1
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

// 获取平衡因子
func (t *aVlTree) balanceFactor(node *TreeNode) int {
	//空节点平衡因子
	if node == nil {
		return 0
	}
	//节点平衡因子 = 左子树高度 - 右子树高度
	return t.height(node.Left) - t.height(node.Right)
}

// 右旋操作
func (t *aVlTree) rightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right
	//以child为原点将node向右旋转
	child.Right = node
	node.Left = grandChild
	//更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)
	//返回旋转后子树的根节点
	return child
}

// 左旋操作
func (t *aVlTree) leftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left
	//以child为原点，将node向左旋转
	child.Left = node
	node.Right = grandChild
	//更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)
	//返回旋转后子树的根节点
	return child
}

// 通过旋转，重新恢复平衡
func (t *aVlTree) rotate(node *TreeNode) *TreeNode {
	//获取节点node的平衡因子
	bf := t.balanceFactor(node)
	//左偏树
	if bf > 1 {
		if t.balanceFactor(node.Left) >= 0 {
			//右旋
			return t.rightRotate(node)
		} else {
			//先左旋后右旋
			node.Left = t.leftRotate(node.Left)
			return t.rightRotate(node)
		}
	}
	//右偏树
	if bf < -1 {
		if t.balanceFactor(node.Right) <= 0 {
			//左旋
			return t.leftRotate(node)
		} else {
			//先右旋后左旋
			node.Right = t.rightRotate(node.Right)
			return t.leftRotate(node)
		}
	}
	//平衡树，无需旋转，直接返回
	return node
}

// 插入节点
func (t *aVlTree) insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

// 递归插入节点（辅助函数）
func (t *aVlTree) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewTreeNode(val)
	}
	//1.查找插入位置并插入节点
	if val < node.Val.(int) {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val.(int) {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		//重复节点不插入，直接返回
		return node
	}
	//更新节点高度
	t.updateHeight(node)
	//2.执行旋转操作，使树恢复平衡
	node = t.rotate(node)
	//返回子树的根节点
	return node
}

// 删除节点
func (t *aVlTree) remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

// 递归删除节点（辅助函数）
func (t *aVlTree) removeHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	//1.查找结点并删除
	if val < node.Val.(int) {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val.(int) {
		node.Right = t.removeHelper(node.Right, val)
	} else {
		//子节点数量为0或1
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				//子节点数量=1，直接删除node,并返回
				return nil
			} else {
				//子节点数量=1，直接删除node
				node = child
			}
		} else {
			//子节点数量=2，则将中序遍历的下一个节点删除（右子树最左节点），并用该节点替换当前节点
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			node.Right = t.removeHelper(node.Right, temp.Val.(int))
			node.Val = temp.Val
		}
	}

	// 如果节点已经被删除（返回nil），则无需更新高度和旋转
	if node == nil {
		return nil
	}

	//更新节点高度
	t.updateHeight(node)
	//2.执行旋转操作，使树重新平衡
	node = t.rotate(node)
	//返回子树根节点
	return node
}

// 查找节点
func (t *aVlTree) search(num int) *TreeNode {
	node := t.root
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

// 打印树
func (t *aVlTree) print() {
	PrintTree(t.root)
}
