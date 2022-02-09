package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string //优先级队列中的数据，可以是任意类型，这里使用string
	priority int    //优先级队列中节点的优先级
	index    int    // index是该节点在堆中的位置
}

type PriorityQueue []*Item

//绑定Len方法
func (pq PriorityQueue) Len() int {
	return len(pq)
}

//绑定Less方法，这里用的是小于号，生成的是小根堆
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

//绑定swap方法
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

//绑定Pop方法，并将index置为-1是为了标识该数据已经出了优先级队列了
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

//绑定push方法
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

//更改修改优先级和值的item在优先级队列中的位置
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	//创建节点并设计他们的优先级
	items := map[string]int{"二毛": 5, "张三": 3, "狗蛋": 9}
	i := 0
	pq := make(PriorityQueue, len(items)) //创建优先级队列，并初始化
	for k, v := range items {
		//将节点放到优先级队列中
		pq[i] = &Item{
			value:    k,
			priority: v,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	item := &Item{
		value:    "李四",
		priority: 4,
	}
	heap.Push(&pq, item)
	//pq.update(item, item.value, 6) //更新item的优先级
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d: %s index: %.2d\n", item.priority, item.value, item.index)
	}
}
