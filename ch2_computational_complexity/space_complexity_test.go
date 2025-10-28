package ch2_computational_complexity

import (
	. "datastructure/pkg"
	"testing"
)

// 空间复杂度测试用例
func TestSpaceComplexity(t *testing.T) {
	n := 5
	//常数阶
	spaceConstant(n)
	//线性阶
	spaceLinear(n)
	spaceLinearRecur(n)
	//平方阶
	spaceQuadratic(n)
	spaceQuadraticRecur(n)
	//指数阶
	root := buildTree(n)
	PrintTree(root)
}
