package ch5_stack_and_queue

import (
	. "datastructure/pkg"
	"fmt"
	"testing"
)

// 栈测试
func TestStack(t *testing.T) {
	//初始化栈 使用切片模拟
	var stack []int

	//元素入栈
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)
	stack = append(stack, 5)
	stack = append(stack, 4)
	fmt.Print("栈 stack = ")
	PrintSlice(stack)

	//访问栈顶元素
	peek := stack[len(stack)-1]
	fmt.Println("栈顶元素 peek =", peek)

	//元素出栈
	pop := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Print("出栈元素 pop =", pop, "出栈后 栈 stack = ")
	PrintSlice(stack)

	//获取栈长度
	size := len(stack)
	fmt.Println("栈的长度 size =", size)

	//判断是否栈空
	isEmpty := len(stack) == 0
	fmt.Println("栈是否为空  =", isEmpty)

}

// 测试链表栈
func TestLinkedListStack(t *testing.T) {

	//初始化栈
	stack := newLinkedListStack()

	//元素入栈
	stack.push(1)
	stack.push(3)
	stack.push(2)
	stack.push(5)
	stack.push(4)
	fmt.Print("栈 stack = ")
	PrintList(stack.toList())

	//访问栈顶元素
	peek := stack.peek()
	fmt.Println("栈顶元素 peek =", peek)

	//元素出栈
	pop := stack.pop()
	fmt.Print("出栈元素 pop =", pop, "出栈后 栈 stack = ")
	PrintList(stack.toList())

	//获取栈的长度
	size := stack.size()
	fmt.Println("栈的长度为 size =", size)

	//判断是否栈空
	isEmpty := stack.isEmpty()
	fmt.Println("栈是否为空 =", isEmpty)
}

// 测试数组栈
func TestArrayStack(t *testing.T) {

	//初始化栈
	stack := newArrayStack()

	//入栈
	stack.push(11)
	stack.push(22)
	stack.push(33)
	stack.push(44)
	stack.push(55)

	//访问栈顶元素
	peek := stack.peek()
	fmt.Println("栈顶元素 peek =", peek)

	//元素出栈
	pop := stack.pop()
	fmt.Print("出栈元素 pop =", pop, "出栈后 栈 stack = ")
	PrintSlice(stack.toList())

	//获取栈的长度
	size := stack.size()
	fmt.Println("栈的长度为 size =", size)

	//判断是否栈空
	isEmpty := stack.isEmpty()
	fmt.Println("栈是否为空 =", isEmpty)
}
