package ch6_hashing

import (
	. "datastructure/pkg"
	"fmt"
	"strconv"
	"testing"
)

// 哈希表
func TestHashMap(t *testing.T) {
	//初始化哈希表
	hmap := make(map[int]string)

	//添加操作 键值对（key,value）
	hmap[12836] = "小哈"
	hmap[15937] = "小啰"
	hmap[16750] = "小算"
	hmap[13276] = "小法"
	hmap[10583] = "小鸭"
	fmt.Println("添加数据后 哈希表 hmap为\nKey -> Value")
	PrintMap(hmap)

	//查询操作 通过键key 得到value
	name := hmap[15937]
	fmt.Println("\n查询 key = 15937 value = ", name)

	//查询操作 通过键key 得到value
	name = hmap[20336]
	fmt.Println("\n查询 key = 20336 value = ", name)

	//查询操作 通过键key 得到value
	name = hmap[12836]
	fmt.Println("\n查询 key = 12836 value = ", name)

	//删除操作 删除键值对
	delete(hmap, 10583)
	fmt.Println("\n删除键值对 10583 哈希表 hmap为\nKey -> Value")
	PrintMap(hmap)

	//遍历哈希表 键值对 Key->Value
	fmt.Println("\n遍历键值对 Key->Value")
	for key, value := range hmap {
		fmt.Println(key, "->", value)
	}

	//遍历哈希表 键 Key
	fmt.Println("\n单独遍历键 Key")
	for key := range hmap {
		fmt.Println(key)
	}

	//遍历哈希表 值 Key
	fmt.Println("\n单独遍历值 Value")
	for _, value := range hmap {
		fmt.Println(value)
	}
}

// 测试哈希算法
func TestSimpleHash(t *testing.T) {
	var hash int

	key := "Hello 算法"
	hash = addHash(key)
	fmt.Println("加法哈希值为 " + strconv.Itoa(hash))

	hash = mulHash(key)
	fmt.Println("乘法哈希值为 " + strconv.Itoa(hash))

	hash = xorHash(key)
	fmt.Println("亦或哈希值为 " + strconv.Itoa(hash))

	hash = rotHash(key)
	fmt.Println("旋转哈希值为 " + strconv.Itoa(hash))
}
