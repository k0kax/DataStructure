package ch4_array_and_linkedlist

import (
	. "datastructure/pkg"
	"fmt"
	"testing"
)

//链表结构体
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 新建链表
//
//	func NewListNode(val int) *ListNode {
//		return &ListNode{
//			Val:  val,
//			Next: nil,
//		}
//	}
func TestLinkedList(t *testing.T) {
	n0 := NewListNode(1)
	n1 := NewListNode(3)
	n2 := NewListNode(2)
	n3 := NewListNode(5)
	n4 := NewListNode(4)

	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	fmt.Println("初始化链表为:")
	PrintLinkedList(n0)

	//插入节点
	n5 := NewListNode(888)
	insertNode(n0, n5)
	fmt.Println("插入后，链表为:")
	PrintLinkedList(n0)

	//删除节点
	removeItem(n0)
	fmt.Println("删除后，链表为:")
	PrintLinkedList(n0)

	//访问节点
	node := access(n0, 4)
	fmt.Println("链表索引4处的节点为:", node)

	//查找节点
	index := findNode(n0, 3)
	fmt.Println("链表节点值为3的索引为:", index)

}
