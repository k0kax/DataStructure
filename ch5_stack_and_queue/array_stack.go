package ch5_stack_and_queue

//基于数组的栈
type arrayStack struct {
	data []int //切片模拟栈
}

//初始化栈
func newArrayStack() *arrayStack {
	return &arrayStack{
		//设置栈的长度为0，容量为16
		data: make([]int, 0, 16),
	}
}

//栈的长度
func (s *arrayStack) size() int {
	return len(s.data)
}

//栈是否为空
func (s *arrayStack) isEmpty() bool {
	return s.size() == 0
}

//入栈
func (s *arrayStack) push(v int) {
	//切片自动扩容
	s.data = append(s.data, v)
}

//出栈
func (s *arrayStack) pop() any {
	val := s.peek()
	s.data = s.data[:len(s.data)-1]
	return val
}

//获取栈顶元素
func (s *arrayStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	val := s.data[:len(s.data)-1]
	return val
}

//获取切片用于打印
func (s *arrayStack) toList() []int {
	return s.data
}
