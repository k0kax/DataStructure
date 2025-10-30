package ch5_stack_and_queue

import (
	"container/list"
	. "datastructure/pkg"
	"fmt"
	"testing"
)

// 双向队列
// list 双向队列
func TestDeque(t *testing.T) {
	//初始化队列
	deque := list.New()

	//元素入队
	deque.PushBack(2)
	deque.PushBack(5)
	deque.PushBack(4)
	deque.PushBack(3)
	deque.PushBack(1)
	fmt.Print("双向队列 deque = ")
	PrintList(deque)

	//访问元素
	front := deque.Front()
	fmt.Println("队首元素 front = ", front.Value)
	rear := deque.Back()
	fmt.Println("队尾元素 rear = ", rear.Value)

	//队首元素出队
	deque.Remove(front)
	fmt.Print("队首出队元素 front = ", front.Value, " 队首出队后 deque = ")
	PrintList(deque)

	//队尾元素出队
	deque.Remove(rear)
	fmt.Print("队尾出队元素 front = ", front.Value, " 队尾出队后 deque = ")
	PrintList(deque)

	//双向队列长度
	size := deque.Len()
	fmt.Println("双向队列长度 size = ", size)

	//判断双向队列是否为空
	isEmpty := deque.Len() == 0
	fmt.Println("双向队列是否为空", isEmpty)
}

func TestArrayDeque(t *testing.T) {
	//初始化队列
	deque := NewArrayDeque(16)

	//元素队尾入队
	deque.pushLast(111)
	deque.pushLast(222)
	deque.pushLast(333)
	fmt.Print("双向队列 deque = ")
	PrintSlice(deque.toSlice())

	//访问元素
	front := deque.peekFirst()
	fmt.Println("队首元素 front = ", front)
	rear := deque.peekLast()
	fmt.Println("队尾元素 rear = ", rear)

	//队首入队
	deque.pushFirst(888)
	fmt.Print("元素 888 队尾入队后 deque = ")
	PrintSlice(deque.toSlice())
	deque.pushFirst(777)
	fmt.Print("元素 777 队尾入队后 deque = ")
	PrintSlice(deque.toSlice())

	//元素队首出队
	popFirst := deque.popFirst()
	fmt.Printf("队首元素%d出队后 deque = ", popFirst)
	PrintSlice(deque.toSlice())
	popFirst = deque.popFirst()
	fmt.Printf("队首元素%d出队后 deque = ", popFirst)
	PrintSlice(deque.toSlice())

	//元素队尾出队
	popLast := deque.popLast()
	fmt.Printf("队尾元素%d出队后 deque = ", popLast)
	PrintSlice(deque.toSlice())
	popLast = deque.popLast()
	fmt.Printf("队尾元素%d出队后 deque = ", popLast)
	PrintSlice(deque.toSlice())

	//获取双向队列长度
	size := deque.size()
	fmt.Println("双向队列长度 size = ", size)

	//双向队列是否为空
	isEmpty := deque.isEmpty()
	fmt.Println("双向队列是否为空  = ", isEmpty)
}

// 测试链表实现的双向队列
func TestLinkedListDeque(t *testing.T) {
	//初始化队列
	deque := newLinkedListDeque()

	//元素队尾入队
	deque.pushLast(111)
	deque.pushLast(222)
	deque.pushLast(333)
	deque.pushLast(444)
	deque.pushLast(555)
	fmt.Print("双向队列 deque = ")
	PrintList(deque.toList())

	//访问元素
	front := deque.peekFirst()
	fmt.Println("队首元素 front = ", front)
	rear := deque.peekLast()
	fmt.Println("队尾元素 rear = ", rear)

	//队首入队
	deque.pushFirst(888)
	fmt.Print("元素 888 队尾入队后 deque = ")
	PrintList(deque.toList())
	deque.pushFirst(777)
	fmt.Print("元素 777 队尾入队后 deque = ")
	PrintList(deque.toList())

	//元素队首出队
	popFirst := deque.popFirst()
	fmt.Printf("队首元素%d出队后 deque = ", popFirst)
	PrintList(deque.toList())
	popFirst = deque.popFirst()
	fmt.Printf("队首元素%d出队后 deque = ", popFirst)
	PrintList(deque.toList())
	//元素队尾出队
	popLast := deque.popLast()
	fmt.Printf("队尾元素%d出队后 deque = ", popLast)
	PrintList(deque.toList())
	popLast = deque.popLast()
	fmt.Printf("队尾元素%d出队后 deque = ", popLast)
	PrintList(deque.toList())

	//获取双向队列长度
	size := deque.size()
	fmt.Println("双向队列长度 size = ", size)

	//双向队列是否为空
	isEmpty := deque.isEmpty()
	fmt.Println("双向队列是否为空  = ", isEmpty)
}
