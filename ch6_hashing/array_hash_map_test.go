package ch6_hashing

import (
	"fmt"
	"testing"
)

// 测试数组实现的哈希表
func TestArrayHashMap(t *testing.T) {
	//初始化哈希表
	hmap := newArrayHashMap()

	//添加键值对
	hmap.put(12836, "小哈")
	hmap.put(15937, "小啰")
	hmap.put(16750, "小算")
	hmap.put(13276, "小法")
	hmap.put(10583, "小鸭")
	fmt.Println("添加完成后，哈希表为\nKey -> Value")
	hmap.print()

	//查询操作 通过key查询value
	name := hmap.get(15937)
	fmt.Println("\nkey = 15937 value = ", name)

	//删除操作
	hmap.remove(10583)
	fmt.Println("删除完成后，哈希表为\nKey -> Value")
	hmap.print()

	//遍历哈希表
	fmt.Println("\n遍历键值对 Key->Value")
	for _, kv := range hmap.pairSet() {
		fmt.Println(kv.key, " -> ", kv.val)
	}

	fmt.Println("\n遍历键Key")
	for _, key := range hmap.keySet() {
		fmt.Println(key)
	}

	fmt.Println("\n遍历值Value")
	for _, val := range hmap.valueSet() {
		fmt.Println(val)
	}

}
