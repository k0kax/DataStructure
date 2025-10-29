package array

import "strconv"

// 回文数
// 数组解决
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	arr := []int{}
	var count, num int //位数
	for x != 0 {
		num = x % 10 //取每位对应的数字
		arr = append(arr, num)
		count++
		x = x / 10 //向前移动
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false //跳出循环
		}
	}
	return true
}

// 转换成字符串
func isPalindromeII(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}

	}
	return true
}

// 一半反转 官方推荐
func isPalindromeIII(x int) bool {
	// 小于0，不是回文数
	// 多位，且最后一位为0的也不是回文数
	// 0是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	new := 0      //反转数字
	for x > new { //剩余数字小于反转数字 反转一半
		new = new*10 + x%10
		x /= 10
	}

	//两数相等 或 剩余数=反转数/10（奇数个位数 12321）去掉中间的位
	return x == new || x == new/10
}

//也可以全反转
