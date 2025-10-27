package ch2_computational_complexity

//递归
import (
	"container/list"
)

func recur(n int) int {
	if n == 1 {
		return 1
	}
	res := recur(n - 1)
	return n + res
}

func tailRecur(n int, res int) int {
	if n == 0 {
		return res
	}
	return tailRecur(n-1, res+n)
}

// 斐波那契数列
func fib(n int) int {
	if n == 1 || n == 2 {
		return n - 1
	}
	res := fib(n-1) + fib(n-2)
	return res
}

func forLoopRecur(n int) int {
	stack := list.New()
	res := 0
	//递
	for i := n; i > 0; i-- {
		//入栈操作 模拟递
		stack.PushBack(i)
	}
	//归
	for stack.Len() != 0 {
		//出栈操作 模拟归
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}
	return res
}
