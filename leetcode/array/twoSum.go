package array

//两数之和
func twoSum(nums []int, target int) []int {
	//var arr [2]int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumII(nums []int, target int) []int {
	m := make(map[int]int, len(nums)) // 初始化哈希表

	for i, num := range nums {
		diff := target - num              // 计算需要配对的差值
		if j, exists := m[diff]; exists { // 检查差值是否在哈希表中
			return []int{j, i} // 找到配对，返回索引
		}
		m[num] = i // 将当前数字和索引存入哈希表
	}

	return nil
}
