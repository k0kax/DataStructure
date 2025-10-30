package ch5_stack_and_queue

import "container/list"

//基于链表的栈
type linkedListStack struct {

	//使用内置的list包实现栈
	data *list.List
}

//初始化栈
func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(), //新建list
	}
}

//入栈
func (s *linkedListStack) push(value int) {
	s.data.PushBack(value)
}

//出栈
func (s *linkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}

	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

//获取栈顶元素
func (s *linkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	return e.Value
}

//获取栈的长度
func (s *linkedListStack) size() int {
	return s.data.Len()
}

//判断栈是否为空
func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

//获取list用于打印
func (s *linkedListStack) toList() *list.List {
	return s.data
}
