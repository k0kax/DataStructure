package ch7_tree

import (
	. "datastructure/pkg"
	"fmt"
	"testing"
)

func TestArrayBinaryTree(t *testing.T) {
	//初始化二叉树
	arr := []any{1, 2, 3, 4, nil, 6, 7, 8, 9, nil, nil, 12, nil, nil, 15}
	root := SliceToTree(arr)
	fmt.Println("\n初始化二叉树")
	fmt.Println("二叉树的数组表示:")
	fmt.Println(arr)
	fmt.Println("二叉树的链表表示:")
	PrintTree(root)

	//数组表示下的二叉树
	abt := newArrayBinaryTree(arr)

	//访问节点
	i := 1
	l := abt.left(i)
	r := abt.right(i)
	p := abt.parent(i)
	fmt.Println("\n当前节点的索引为", i, ",值为", abt.val(i))
	fmt.Println("其左子节点的索引为", l, ",值为", abt.val(l))
	fmt.Println("其右子节点的索引为", r, ",值为", abt.val(r))
	fmt.Println("其父节点的索引为", p, ",值为", abt.val(p))

	//遍历树
	res := abt.levelOrder()
	fmt.Println("\n层序遍历为：", res)
	res = abt.preOrder()
	fmt.Println("\n前序遍历为：", res)
	res = abt.postOrder()
	fmt.Println("\n后序遍历为：", res)
}
