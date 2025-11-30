package heap

//堆，一种特殊的树
// Go 语言中可以通过实现 heap.Interface 来构建整数大顶堆
// 实现 heap.Interface 需要同时实现 sort.Interface
type intHeap []any

//push 入堆
func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

//pop 出堆 弹出堆顶元素
func (h *intHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// len sort.Interface 函数
func (h *intHeap) Len() int {
	return len(*h)
}

// less sort.Interface 方法
func (h *intHeap) Less(i, j int) bool {
	// 如果实现小顶堆，则需要调整为小于号
	return (*h)[i].(int) > (*h)[j].(int)
}

// Swap sort.Interface 交换
func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

//Top 获取堆顶元素
func (h *intHeap) Top() any {
	return (*h)[0]
}
