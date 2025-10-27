package ch2_computational_complexity

import (
	"fmt"
	"testing"
)

// 迭代测试用例
func TestIteration(t *testing.T) {
	n := 5
	res := forLoop(n)
	fmt.Println("\n for 循环结果求和 res = ", res)

	res = whileLoop(n)
	fmt.Println("\n while 循环结果求和 res = ", res)

	res = whileLoopII(n)
	fmt.Println("\n while 循环(两次更新)结果求和 res = ", res)

	resStr := nestedForLoop(n)
	fmt.Println("\n while 循环结果求和 res = ", resStr)
}
