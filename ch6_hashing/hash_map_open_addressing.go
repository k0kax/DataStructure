package ch6_hashing

import "fmt"

//哈希冲突解决方案2 开放寻址
type hashMapOpenAddressing struct {
	size        int     //键值对数
	capacity    int     //哈希表容量
	loadThres   float64 //触发扩容的负载因子阈值
	extendRatio int     //扩容倍数
	buckets     []*pair //桶数组
	TOMBSTONE   *pair   //删除标记 墓碑
}

//构造方法
func newHashMapOpenAddressing() *hashMapOpenAddressing {
	return &hashMapOpenAddressing{
		size:        0,
		capacity:    4,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     make([]*pair, 4),
		TOMBSTONE:   &pair{-1, "-1"},
	}
}

//哈希函数
func (h *hashMapOpenAddressing) hashFunc(key int) int {
	return key % h.capacity
}

//负载因子
func (h *hashMapOpenAddressing) loadFactor() float64 {
	return float64(h.size) / float64(h.capacity)
}

//搜索key对应的桶索引
func (h *hashMapOpenAddressing) findBucket(key int) int {
	index := h.hashFunc(key)
	firstTombstone := -1          //第一个墓碑 删除标记
	for h.buckets[index] != nil { //桶不为空
		if h.buckets[index].key == key { //找到目标的key
			if firstTombstone != -1 {
				//若之前遇到了删除标记（墓碑），则将键值对移至该索引处
				h.buckets[firstTombstone] = h.buckets[index] //将当前键值对移动到第一个墓碑位置（优化后续查找）
				h.buckets[index] = h.TOMBSTONE               //原位置标记为墓碑
				return firstTombstone                        //返回移动后的桶索引
			}
			return index //没遇到墓碑 直接返回找到的索引
		}
		//没有匹配到
		//发现是墓碑（删除标记） 首次发现墓碑并标记
		if firstTombstone == -1 && h.buckets[index] == h.TOMBSTONE {
			firstTombstone = index //记录首个遇到的删除标记位置 墓碑
		}
		index = (index + 1) % h.capacity // 线性探测：每次向前移动1位（循环）
	}
	//若key不存在，则返回添加点的索引
	if firstTombstone != -1 {
		return firstTombstone
	}
	return index //当前桶为空
}

//查询操作
func (h *hashMapOpenAddressing) get(key int) string {
	index := h.findBucket(key)
	if h.buckets[index] != nil && h.buckets[index] != h.TOMBSTONE {
		return h.buckets[index].val
	}
	return ""
}

//添加操作
func (h *hashMapOpenAddressing) put(key int, val string) {
	if h.loadFactor() > h.loadThres {
		//超过阈值 执行扩容
		h.extend()
	}
	index := h.findBucket(key)
	if h.buckets[index] == nil || h.buckets[index] == h.TOMBSTONE { //桶为空或者为墓碑
		h.buckets[index] = &pair{key, val} //键值对不存在，则添加价值对
		h.size++
	} else {
		h.buckets[index].val = val
	}
}

//删除操作
func (h *hashMapOpenAddressing) remove(key int) {
	index := h.findBucket(key)
	if h.buckets[index] != nil && h.buckets[index] != h.TOMBSTONE {
		h.buckets[index] = h.TOMBSTONE //找到键值对置为墓碑
		h.size--
	}
}

//扩容哈希表
func (h *hashMapOpenAddressing) extend() {
	oldBuckets := h.buckets     //暂存旧的哈希表
	h.capacity *= h.extendRatio //更新容量
	h.buckets = make([]*pair, h.capacity)
	h.size = 0

	//将旧哈希表键值对搬运到新哈希表
	for _, pair := range oldBuckets {
		if pair != nil && pair != h.TOMBSTONE {
			h.put(pair.key, pair.val)
		}
	}
}

//打印哈希表
func (h *hashMapOpenAddressing) print() {
	for _, pair := range h.buckets {
		if pair == nil {
			fmt.Println("nil")
		} else if pair == h.TOMBSTONE {
			fmt.Println("TOMNSTONE")
		} else {
			fmt.Printf("%d -> %s\n", pair.key, pair.val)
		}
	}
}
