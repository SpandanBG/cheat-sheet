package arrays

import (
	"container/heap"
)

type Item struct {
	value int
	index int
}

type PQue struct {
	maxHeap   []*Item
	entryHeap []*Item
}

func (p PQue) Len() int {
	return len(p.maxHeap)
}

func (p PQue) Less(i, j int) bool {
	return p.maxHeap[i].value > p.maxHeap[j].value
}

func (p PQue) Swap(i, j int) {
	p.maxHeap[i], p.maxHeap[j] = p.maxHeap[j], p.maxHeap[i]
	p.maxHeap[i].index = i
	p.maxHeap[j].index = j
}

func (p *PQue) Push(x any) {
	switch y := x.(type) {
	case int:
		n := len(p.maxHeap)
		item := &Item{
			value: y,
			index: n,
		}

		if len(p.maxHeap) == 0 {
			p.maxHeap = []*Item{item}
			p.entryHeap = []*Item{item}
		} else {
			p.maxHeap = append(p.maxHeap, item)
			p.entryHeap = append(p.entryHeap, item)
		}
	default:
		return
	}
}

func (p *PQue) Pop() any {
	return nil
}

func (p *PQue) update(i *Item, value int) {
	i.value = value
	heap.Fix(p, i.index)
}

// You are given an array of integers nums, there is a sliding window of size k
// which is moving from the very left of the array to the very right. You can only
// see the k numbers in the window. Each time the sliding window moves right by one
// position.
//
// Return the max sliding window.
//
// Example 1:
//
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//
//	1 [3  -1  -3] 5  3  6  7       3
//	1  3 [-1  -3  5] 3  6  7       5
//	1  3  -1 [-3  5  3] 6  7       5
//	1  3  -1  -3 [5  3  6] 7       6
//	1  3  -1  -3  5 [3  6  7]      7
//
// Example 2:
//
// Input: nums = [1], k = 1
// Output: [1]
func MaxHeap_SlidingWindowMax(nums []int, k int) []int {
	ans := make([]int, 0)

	/* Note: */

	p := PQue{
		maxHeap:   make([]*Item, 0),
		entryHeap: make([]*Item, 0),
	}

	heap.Init(&p)

	for _, value := range nums[:k] {
		heap.Push(&p, value)
	}

	for i, value := range nums[k:] {
		ans = append(ans, p.maxHeap[0].value)
		next := p.entryHeap[i%k]
		p.update(next, value)
	}
	ans = append(ans, p.maxHeap[0].value)

	return ans
}
