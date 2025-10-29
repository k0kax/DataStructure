package ch4_array_and_linkedlist

import (
	"fmt"
	"math/rand"
)

//var arr [5]int

//nums := []int{1, 3, 2, 5, 4}

// 随机访问元素
func randonAccess(nums []int) (randomNum int) {

	randomIndex := rand.Intn(len(nums))

	randomNum = nums[randomIndex]

	return
}

// 通过数组索引插入元素
func insert(nums []int, num int, index int) {

	//将索引index及以后元素后移一位
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}

	//插入元素
	nums[index] = num
}

// 通过索引删除元素
func remove(nums []int, index int) {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

// 遍历数组
func traverse(nums []int) {
	// count := 0

	// //通过索引遍历数组
	// for i := 0; i < len(nums)-1; i++ {
	// 	count += nums[i]
	// }

	// count = 0
	// //直接遍历数组
	// for _, num := range nums {
	// 	count += num
	// }

	//同时遍历数据索引和元素
	for i, num := range nums {
		fmt.Printf("索引%d:%d\t", i, num)
		// count += nums[i]
		// count += num
	}

}

// 查找数组 指定元素
func find(nums []int, target int) int {
	index := 0
	for index = range nums {
		if nums[index] == target {
			break
		}
	}
	return index
}

// 扩容数组长度
func extend(nums []int, enlarge int) []int {
	res := make([]int, len(nums)+enlarge) //扩容？新建一个兼容的数组罢了

	// for i, num := range nums {
	// 	res[i] = num
	// }
	//与下面相同
	copy(res, nums)
	return res
}
