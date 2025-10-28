package ch4_array_and_linkedlist

import (
	"fmt"
	"sort"
	"testing"
)

//列表

func TestList(t *testing.T) {

	//初始化列表
	//nums1 :=[]int{}
	nums := []int{1, 3, 2, 5, 4}
	fmt.Println("列表 nums =", nums)

	//访问元素
	num := nums[1]
	fmt.Println("访问索引1处的元素,得到 num = ", num)

	//更新数据
	nums[1] = 0
	fmt.Println("索引1处的元素更新为0,得到 num = ", num)

	//清空列表
	nums = nil
	fmt.Println("清空列表后 nums =", nums)

	//尾部添加元素
	nums = append(nums, 1)
	nums = append(nums, 3)
	nums = append(nums, 2)
	nums = append(nums, 5)
	nums = append(nums, 4)
	fmt.Println("添加元素后 nums =", nums)

	//中间添加元素
	nums = append(nums[:3], append([]int{6}, nums[3:]...)...)
	fmt.Println("在索引3处插入数字6,得到 nums =", nums)

	//删除元素
	nums = append(nums[:3], nums[4:]...)
	fmt.Println("删除索引3处的元素,得到 nums=", nums)

	//通过索引遍历列表
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}

	//直接遍历列表元素
	count = 0
	for _, x := range nums {
		count += x
	}

	//拼接这两个列表
	nums1 := []int{6, 8, 7, 10, 9}
	nums = append(nums, nums1...)
	fmt.Println("将列表 nums1 拼接到 nums 之后，得到 nums =", nums)

	//排序
	sort.Ints(nums) //排序后，列表从小排到大
	fmt.Println("排序列表后 nums =", nums)
}
