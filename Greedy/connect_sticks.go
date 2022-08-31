package greedy

//1167 leetcode --> minimum cost of connecting sticks

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func ConnectSticksMinimumCost(arr []int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, v := range arr {
		heap.Push(h, v)
	}
	result := 0

	for len(*h) != 1 {
		temp := heap.Pop(h).(int) + heap.Pop(h).(int)
		result += temp
		heap.Push(h, temp)
	}
	return result
}
