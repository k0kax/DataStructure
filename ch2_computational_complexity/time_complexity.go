package ch2_computational_complexity

//时间复杂度测试用例
//常数阶
func constant(n int) int {
	count := 0
	size := 100000
	for i := 0; i < size; i++ {
		count++
	}
	return count
}

//线性阶
func linear(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count++
	}
	return count
}

//线性阶 遍历数组
func arrayTraversal(nums []int) int {
	count := 0
	for range nums {
		count++
	}
	return count
}

//平方阶
func quadratic(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count++
		}
	}
	return count
}

//平方阶 冒泡排序
func bubbleSort(nums []int) int {
	count := 0
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
				count += 3
			}
		}
	}
	return count
}

//--------------------------------------存疑-------------------------------------------------------
//指数阶---细胞分裂---计算总分裂次数
func exponential(n int) int {
	count, base := 0, 1
	for i := 1; i <= n; i++ {
		//每轮分裂次数
		for j := 1; j <= base; j++ {
			count++ //每一轮中 每个细胞分裂一次
		}
		base *= 2 //每轮对应细胞数 base
	}
	return count //总分裂次数 count = 1 + 2 + 4 + 8 + .. + 2^(n-1) = 2^n - 1
	//第一轮 分裂1次
	//第二轮 分裂2次
	//第三轮 分裂4次
	//... ... ...
}

//指数阶--细胞分裂--递归实现
func expRecur(n int) int {
	if n == 1 {
		return 1
	}
	//上一轮分裂（两）+本轮分裂（1）
	return 2*expRecur(n-1) + 1 //上一代 + 本次
}

//--------------------------------------存疑-------------------------------------------------------

//对数阶
func logarithmic(n int) int {
	count := 0
	for n > 1 {
		n = n / 2
		count++
	}
	return count
}

//对数阶 递归实现
func logRecur(n int) int {
	if n <= 1 {
		return 0
	}
	return logRecur(n/2) + 1
}

//线性对数阶 n log n
func linearLogRecur(n int) int {
	if n <= 1 {
		return 1
	}
	count := linearLogRecur(n/2) + linearLogRecur(n/2)
	for i := 0; i < n; i++ {
		count++
	}
	return count
}

//阶乘 n! 循环累加实现类似阶乘的效果
func factorialRecur(n int) int {

	if n == 0 {
		return 1
	}

	count := 0

	for i := 0; i < n; i++ {
		count += factorialRecur(n - 1)
	}
	return count
}

/*
n=3
	n=2
		n=1
			n=0 count=1
		n=1
			n=0 count=1
							count=2
	n=2
		n=1
			n=0 count=1
		n=1
			n=0 count=1
							count=2

	n=2
		n=1
			n=0 count=1
		n=1
			n=0 count=1
							count=2
											count=6

*/
