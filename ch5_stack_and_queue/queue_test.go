package ch5_stack_and_queue

import (
	"container/list"
	. "datastructure/pkg"
	"fmt"
	"testing"
)

// 原始队列 list实现
func TestQueue(t *testing.T) {

	//初始化队列
	queue := list.New() //用list模拟队列

	//元素入列
	queue.PushBack(1)
	queue.PushBack(22)
	queue.PushBack(333)
	queue.PushBack(4444)
	queue.PushBack(55555)
	fmt.Print("队列 queue = ")
	PrintList(queue)

	//访问队首元素
	peek := queue.Front()
	fmt.Println("队首元素 = ", peek.Value)

	//元素出队
	pop := queue.Front() //队首出队
	queue.Remove(pop)
	fmt.Print("出队元素 pop =", pop.Value, " 新队列 queue = ")
	PrintList(queue)

	//获取队长
	size := queue.Len()
	fmt.Println("队长为 size = ", size)

	//判断队空
	isEmpty := queue.Len() == 0
	fmt.Println("队是否为空 ", isEmpty)
}

// 链表实现的队列
func TestLinkedListQueue(t *testing.T) {
	//初始化队列
	queue := newLinkedListQueue()

	//元素入队
	queue.push(1)
	queue.push(2)
	queue.push(3)
	queue.push(4)
	queue.push(5)
	fmt.Print("链表队列 queue = ")
	PrintList(queue.toList())

	//队首元素
	peek := queue.peek()
	fmt.Println("出队 peek = ", peek)

	//元素出队
	pop := queue.pop()
	fmt.Print("出队元素 pop = :", pop, " 新链表队列 queue = ")
	PrintList(queue.toList())

	//队是否为空
	isEmpty := queue.isEmpty()
	fmt.Println("队是否为空 ", isEmpty)
	//队长
	size := queue.size()
	fmt.Println("队长 size = ", size)

}

// 测试数组队列 环形
func TestArrayQueue(t *testing.T) {
	//初始化队列
	queue := NewArrayQueue(10)
	if queue.pop() != nil {
		t.Errorf("want:%v,got:%v", nil, queue.pop())
	}

	//元素入队
	queue.push(1)
	queue.push(3)
	queue.push(2)
	queue.push(5)
	queue.push(4)
	fmt.Print("数组队列 queue = ")
	PrintSlice(queue.toSlice())

	//队首元素
	peek := queue.peek()
	fmt.Println("出队 peek = ", peek)

	//元素出队
	pop := queue.pop()
	fmt.Print("出队元素 pop = ", pop, " 新数组队列 queue = ")
	PrintSlice(queue.toSlice())

	//队是否为空
	isEmpty := queue.isEmpty()
	fmt.Println("队是否为空 = ", isEmpty)

	//获取队长
	size := queue.size()
	fmt.Println("队长 size = ", size)

	//测试环形队列
	for i := range 10 {
		queue.push(i)
		queue.pop()
		fmt.Print("第", i, "轮入队 + 出队后 queue = ")
		PrintSlice(queue.toSlice())
	}
}
