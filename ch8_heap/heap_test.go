package heap

import (
	"container/heap"
	. "datastructure/pkg"
	"fmt"
	"testing"
)

func testPush(h *intHeap, val int) {
	heap.Push(h, val)
	fmt.Printf("\n元素 %d 入堆后 \n", val)
	PrintHeap(*h)
}

func testPop(h *intHeap) {
	val := heap.Pop(h)
	fmt.Printf("\n堆顶元素 %d 出堆后 \n", val)
	PrintHeap(*h)
}

func TestHeap(t *testing.T) {
	//初始化堆
	//大顶堆
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	//元素入堆
	testPush(maxHeap, 1)
	testPush(maxHeap, 3)
	testPush(maxHeap, 2)
	testPush(maxHeap, 5)
	testPush(maxHeap, 4)

	//获取堆顶元素
	top := maxHeap.Top()
	fmt.Printf("堆顶元素 %d\n", top)

	//堆顶元素出堆
	testPop(maxHeap)
	testPop(maxHeap)
	testPop(maxHeap)
	testPop(maxHeap)
	testPop(maxHeap)

	//获取堆大小
	size := len(*maxHeap)
	fmt.Printf("\n堆元素数量为 %d\n", size)

	//判断是否堆空
	isEmpty := len(*maxHeap) == 0
	fmt.Printf("堆是否为空 %t\n", isEmpty)
}
