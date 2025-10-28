package ch4_array_and_linkedlist

//列表类
type myList struct {
	arrCapacity int
	arr         []int
	arrSize     int
	extendRatio int
}

//构造函数
func newMyList() *myList {
	return &myList{
		arrCapacity: 10,              //列表容量
		arr:         make([]int, 10), //数组（存储列表元素）
		arrSize:     0,               //列表长度（当前元素数量）
		extendRatio: 2,               //每次列表扩容的倍数
	}
}

//获取列表长度（当前元素数量）
func (l *myList) size() int {
	return l.arrSize
}

//获取列表容量
func (l *myList) capacity() int {
	return l.arrCapacity
}

//访问元素 通过索引
func (l *myList) get(index int) int {
	//越界检查
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	return l.arr[index]
}

//更新元素
func (l *myList) set(num, index int) {
	//越界检查
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	l.arr[index] = num
}

//在尾部添加元素
func (l *myList) add(num int) {
	//元素数量超出容量，触发扩容机制
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	l.arr[l.arrSize] = num
	//更新元素数量
	l.arrSize++
}

//中间插入元素
func (l *myList) insert(num int, index int) {
	//越界检查
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	//元素数量超出容量，触发扩容机制
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	//index索引后的元素后移一位
	for i := l.arrSize; i >= index; i-- {
		l.arr[i+1] = l.arr[i]
	}

	l.arr[index] = num
	//更新元素数量
	l.arrSize++
}

//删除元素
func (l *myList) remove(index int) int {
	//越界检查
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	num := l.arr[index]
	for i := index; i < l.arrSize-1; i++ {
		l.arr[i] = l.arr[i+1]
	}
	//更新元素数量
	l.arrSize--
	return num
}

//列表扩容
func (l *myList) extendCapacity() {
	//新建一个长度为原数组 extendRatio 倍的新数组，并将原数组复制到新数组
	l.arr = append(l.arr, make([]int, l.arrCapacity*(l.extendRatio-1))...)
	//更新容量
	l.arrCapacity = len(l.arr)
}

//返回有效长度的列表
func (l *myList) toArray() []int {
	// 仅转换有效长度范围内的列表元素
	return l.arr[:l.arrSize]
}
