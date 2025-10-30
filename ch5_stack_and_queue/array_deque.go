package ch5_stack_and_queue

import "fmt"

// 基于环形数组的双向队列 头尾均可出队入队
type arraryDeque struct {
	nums        []int //存储双向队列元素 用切片
	front       int   //队首指针
	queSize     int   //队列长度
	queCapacity int   //队列容量
}

//初始化队列
func NewArrayDeque(queCapacity int) *arraryDeque {
	return &arraryDeque{
		nums:        make([]int, queCapacity),
		queCapacity: queCapacity,
		front:       0,
		queSize:     0,
	}
}

//双向队列长度
func (q *arraryDeque) size() int {
	return q.queSize
}

//判断双向队列是否为空
func (q *arraryDeque) isEmpty() bool {
	return q.queSize == 0
}

//计算环形数组索引
func (q *arraryDeque) index(i int) int {
	// 通过取余操作实现数组首尾相连
	// 当 i 越过数组尾部后，回到头部
	// 当 i 越过数组头部后，回到尾部
	return (i + q.queCapacity) % q.queCapacity
}

//队首入队
func (q *arraryDeque) pushFirst(num int) {

	if q.queSize == q.queCapacity { //队满
		fmt.Println("双向队列已满")
		return
	}

	//队首指针的位置
	//队首指针向左移一位
	q.front = q.index(q.front - 1)
	q.nums[q.front] = num
	q.queSize++
}

//队尾入队
func (q *arraryDeque) pushLast(num int) {
	if q.queSize == q.queCapacity { //队满
		fmt.Println("双向队列已满")
		return
	}

	//队尾指针 取余运算 rear查过队尾回到头部
	rear := q.index(q.front + q.queSize)

	//元素添加到队尾
	q.nums[rear] = num
	q.queSize++

}

//队首出队
func (q *arraryDeque) popFirst() any {
	//获取队首
	num := q.peekFirst()
	if num == nil {
		return nil
	}

	//队首 后移 越过队尾则返回头部
	q.front = q.index(q.front + 1)
	q.queSize--

	return num
}

//队尾出队
func (q *arraryDeque) popLast() any {
	//获取元素
	num := q.peekLast()
	if num == nil {
		return nil
	}

	//队尾出队 直接出就行
	q.queSize--
	return num
}

//访问队首元素
func (q *arraryDeque) peekFirst() any {
	//队非空
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

//访问队尾元素
func (q *arraryDeque) peekLast() any {
	//队非空
	if q.isEmpty() {
		return nil
	}
	last := q.index(q.front + q.queSize - 1)
	return q.nums[last]
}

//获取Slice用于打印
func (q *arraryDeque) toSlice() []int {
	//仅取有效长度的元素
	res := make([]int, q.queSize)
	for i, j := 0, q.front; i < q.queSize; i++ {
		res[i] = q.nums[q.index(j)]
		j++
	}
	return res
}
