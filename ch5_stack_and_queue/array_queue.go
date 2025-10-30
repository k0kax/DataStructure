package ch5_stack_and_queue

//环形队列  在一段固定的数组模拟队列
type arraryQueue struct {
	nums        []int //存储队列元素的数组
	front       int   //指向队首元素
	queSize     int   //队列长度
	queCapacity int   //队列容量
}

//初始化队列
func NewArrayQueue(queCapacity int) *arraryQueue {
	return &arraryQueue{
		nums:        make([]int, queCapacity),
		queCapacity: queCapacity,
		front:       0,
		queSize:     0,
	}
}

//获取队列长度
func (q *arraryQueue) size() int {
	return q.queSize
}

//判断是否队空
func (q *arraryQueue) isEmpty() bool {
	return q.queSize == 0
}

//入队
func (q *arraryQueue) push(num int) {
	if q.queSize == q.queCapacity { //队满
		return
	}

	//队尾指针 取余运算 rear查过队尾回到头部
	rear := (q.front + q.queSize) % q.queCapacity

	//元素添加到队尾
	q.nums[rear] = num
	q.queSize++

}

//出队
func (q *arraryQueue) pop() any {
	//获取队首
	num := q.peek()
	if num == nil {
		return nil
	}

	//队首 后移 越过队尾则返回头部
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--

	return num
}

//访问队首
func (q *arraryQueue) peek() any {
	//队非空
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

//获取slice 用于打印
func (q *arraryQueue) toSlice() []int {
	rear := (q.front + q.queSize)
	if rear >= q.queCapacity {
		rear %= q.queCapacity
		return append(q.nums[q.front:], q.nums[:rear]...)
	}
	return q.nums[q.front:rear]
}
