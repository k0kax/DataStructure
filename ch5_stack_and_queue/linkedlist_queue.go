package ch5_stack_and_queue

//使用链表实现队列
import "container/list"

type linkedListQueue struct {
	data *list.List
}

//初始化队列
func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		data: list.New(),
	}
}

//入队
func (s *linkedListQueue) push(value any) {
	s.data.PushBack(value)
}

//出队
func (s *linkedListQueue) pop() any {
	if s.isEmpty() { //队列非空
		return nil
	}
	v := s.data.Front()
	s.data.Remove(v)
	return v.Value
}

//访问队首元素
func (s *linkedListQueue) peek() any {
	if s.isEmpty() { //队列非空
		return nil
	}
	peek := s.data.Front()
	return peek.Value
}

//获取队列长度
func (s *linkedListQueue) size() int {
	return s.data.Len()
}

//判断是否队空
func (s *linkedListQueue) isEmpty() bool {
	return s.data.Len() == 0
}

//获取list用于打印
func (s *linkedListQueue) toList() *list.List {
	return s.data
}
