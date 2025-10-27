package ch2_computational_complexity

import (
	"fmt"
	"testing"
)

// 最佳、最差时间复杂度测试用例
func TestWorstBestTimeComplexity(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := 100
		nums := randomNumbers(n)
		index := findOne(nums, 5)
		fmt.Println("数组被打乱后: ", nums)
		fmt.Println("数字5的位置: ", index)
	}

}
