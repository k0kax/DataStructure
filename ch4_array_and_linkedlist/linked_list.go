package ch4_array_and_linkedlist

import (
	. "datastructure/pkg"
)

// 链表n0插入节点  n0之后插入节点P
func insertNode(n0 *ListNode, P *ListNode) {
	n1 := n0.Next
	P.Next = n1
	n0.Next = P
}

// 删除链表的结点 n0之后的首个节点
func removeItem(n0 *ListNode) {
	if n0.Next == nil {
		return
	}
	P := n0.Next
	n1 := P.Next
	n0.Next = n1
	P.Next = nil
}

// 访问链表索引为index的节点
func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next //不断后移
	}
	return head
}

// 查找值为target的节点的索引
func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next //不断后移
		index++          //索引自增一
	}
	return -1
}
