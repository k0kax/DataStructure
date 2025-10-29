package ch4_array_and_linkedlist

import (
	"fmt"
	"testing"
)

func TestMyList(t *testing.T) {
	//初始化列表
	nums := newMyList()
	//尾部添加元素
	nums.add(1)
	nums.add(3)
	nums.add(2)
	nums.add(5)
	nums.add(4)
	fmt.Printf("列表 nums = %v , 容量 = %v , 长度 = %v\n", nums.toArray(), nums.capacity(), nums.size())

	//在中间插入元素
	nums.insert(666, 3)
	fmt.Println("在索引3插入数字666, nums = ", nums.toArray())

	//删除元素
	nums.remove(3)
	fmt.Println("删除索引3, nums = ", nums.toArray())

	//访问元素
	num := nums.get(1)
	fmt.Println("索引1处的值 =  ", num)

	//更新元素
	nums.set(0, 1)
	fmt.Println("将索引1值修改为0, nums = ", nums.toArray())

	//测试扩容机制
	for i := 0; i < 10; i++ {
		nums.add(i)
	}
	fmt.Printf("扩容后 列表 nums = %v , 容量 = %v , 长度 = %v\n", nums.toArray(), nums.capacity(), nums.size())
}
