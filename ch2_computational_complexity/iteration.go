package ch2_computational_complexity

import "fmt"

//迭代
func forLoop(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

//go语言没有while可用for实现类似的效果
func whileLoop(n int) int {
	res := 0
	i := 1
	for i <= n {
		res += i
		i++
	}
	return res
}

func whileLoopII(n int) int {
	res := 0
	i := 1
	for i <= n {
		res += i
		i++
		i *= 2
	}
	return res
}
func nestedForLoop(n int) string {
	res := ""
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			res += fmt.Sprintf("(%d,%d), ", i, j)
		}
	}
	return res
}
