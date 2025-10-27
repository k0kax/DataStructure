package ch2_computational_complexity

import (
	"fmt"
	"testing"
)

// 递归测试用例
func TestRecursion(t *testing.T) {

	n := 5
	res := recur(n)
	fmt.Println("\n递归函数的求和结果 res = ", res)

	res = forLoopRecur(n)
	fmt.Println("\n使用迭代模拟递归求和结果 res = ", res)

	res = tailRecur(n, 0)
	fmt.Println("\n尾递归函数的求和结果 res = ", res)

	res = fib(n)
	fmt.Println("\n斐波那契数列的第", n, "项为", res)

}
