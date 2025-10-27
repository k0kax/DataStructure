package ch2_computational_complexity

/*
	最佳时间复杂度
	最差时间复杂度
	平均时间复杂度
*/

import (
	"math/rand"
)

func randomNumbers(n int) []int {
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	//随机打乱数组
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	return nums
}

// 找出数组中数字1的索引
func findOne(nums []int, n int) int {
	for i := 0; i < len(nums); i++ {
		//当元素在头部时 达到最佳时间复杂度O（1）
		//当元素在尾部时 达到最佳时间复杂度O（n）
		if nums[i] == n {
			return i
		}
	}
	return -1
}
