package ch6_hashing

import (
	"fmt"
	"testing"
)

// 测试链式地址哈希表 封闭寻址
func TestHashMapChaining(t *testing.T) {
	//初始化hash表
	hmap := newHashMapChaining()

	//添加键值对
	hmap.put(12836, "小哈")
	hmap.put(15937, "小啰")
	hmap.put(16750, "小算")
	hmap.put(13276, "小法")
	hmap.put(10583, "小鸭")
	fmt.Println("添加完成后，哈希表为\nKey -> Value")
	hmap.print()

	//查询键值对
	name := hmap.get(15937)
	fmt.Println("\n查询key=15937,value=", name)

	//删除操作
	hmap.remove(12336)
	fmt.Println("\n删除 12836 后，哈希表为\nKey -> Value")
	hmap.print()
}

// 测试开放寻址
func TestHashMapOpenAddressing(t *testing.T) {
	//初始化hash表
	hmap := newHashMapOpenAddressing()

	//添加键值对
	hmap.put(12836, "小哈")
	hmap.put(15937, "小啰")
	hmap.put(16750, "小算")
	hmap.put(13276, "小法")
	hmap.put(10583, "小鸭")
	fmt.Println("添加完成后，哈希表为\nKey -> Value")
	hmap.print()

	//查询键值对
	name := hmap.get(15937)
	fmt.Println("\n查询key=15937,value=", name)

	//删除操作
	hmap.remove(16750)
	fmt.Println("\n删除 16750 后，哈希表为\nKey -> Value")
	hmap.print()

}
