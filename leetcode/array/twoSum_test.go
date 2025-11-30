package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15} //切片
	fmt.Println(nums)
	fmt.Println(twoSumII(nums, 26))
}
