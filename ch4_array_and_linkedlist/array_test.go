package ch4_array_and_linkedlist

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {

	var arr [5]int
	fmt.Println("数组 arr =", arr)

	nums := []int{1, 3, 2, 5, 4}
	fmt.Println("数组 arr =", nums)

	//随机访问
	randomNum := randonAccess(nums)
	fmt.Println("在nums 获取随机元素", randomNum)

	//长度扩展
	nums = extend(nums, 3)
	fmt.Println("将nums 扩容3,得到nums=", nums)

	//插入元素
	insert(nums, 6, 3) //在索引3位置插入6
	fmt.Println("在索引3位置插入6,得到nums=", nums)

	//删除元素
	remove(nums, 2)
	fmt.Println("删除索引2的元素,得到nums=", nums)

	//遍历数组
	traverse(nums)

	//查找元素 按值查找索引
	index := find(nums, 5)
	fmt.Println("元素5的索引位置index=", index)
}
