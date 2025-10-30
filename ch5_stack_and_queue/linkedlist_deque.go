package ch5_stack_and_queue

import "container/list"

//使用链表实现双向队列 使用list
//比数组好用些
type linkedListDeque struct {
	data *list.List
}

//初始化队列
func newLinkedListDeque() *linkedListDeque {
	return &linkedListDeque{
		data: list.New(),
	}
}

//元素队首入队
func (s *linkedListDeque) pushFirst(value any) {
	s.data.PushFront(value)
}

//元素队尾入队
func (s *linkedListDeque) pushLast(value any) {
	s.data.PushBack(value)
}

//队首元素出队
func (s *linkedListDeque) popFirst() any {
	if s.isEmpty() {
		return nil
	}
	num := s.data.Front()
	s.data.Remove(num)
	return num.Value
}

//队尾元素出队
func (s *linkedListDeque) popLast() any {
	if s.isEmpty() {
		return nil
	}
	num := s.data.Back()
	s.data.Remove(num)
	return num.Value
}

//访问队首元素
func (s *linkedListDeque) peekFirst() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Front().Value
}

//访问队尾元素
func (s *linkedListDeque) peekLast() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Back().Value
}

//获取队列长度
func (s *linkedListDeque) size() int {
	return s.data.Len()
}

//判断队列是否为空
func (s *linkedListDeque) isEmpty() bool {
	return s.data.Len() == 0
}

//获取list 用于打印
func (s *linkedListDeque) toList() *list.List {
	return s.data
}
