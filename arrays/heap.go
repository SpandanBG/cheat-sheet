package arrays

import "container/heap"

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
	/*Note
	We are going to create a priority queue where we will only keep the index of
	the largest number at any given point of the sliding window.

	At each iteration, we will first pop out (from the end) any numbers on the
	queue that are smaller than the one we are going to insert iteratively.

	Then we are going to push in the new number -> and calculate the start index
	of the sliding window using (i - k + 1), where:
		- i: the current index that has been entered
		- k: the size of the window

		The formula (i - k + 1) says:
			- For the index `i`
			- We will go back `k` steps
			- And jump forward `+1` step to account for the newly added item

	After finding the starting index of the window, we will check from the start
	of the queue if any index index exists that is outside of the window and
	remove them.

	After our starting index of the window has reached 0, we will then start adding
	the highest value (which will be at the beginning of the queue) to our answer
	list.
	*/

	ans := make([]int, len(nums)-k+1)
	i_ans, w_si := 0, 0

	queue := make([]int, 0)

	for i := 0; i < len(nums); i += 1 {
		// First we remove all the elements in the queue which are smaller than the
		// next item
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}

		// Then we add the new index
		queue = append(queue, i)

		// Then we find out where the window's starting index is
		w_si = i - k + 1

		// Then remove all items from the beginning that has moved away from the
		// start index of the window
		for queue[0] < w_si {
			queue = queue[1:]
		}

		// And finally after the entire window has entered the nums array -> we
		// start collecting the the max values into our answers list
		if w_si >= 0 {
			ans[i_ans] = nums[queue[0]]
			i_ans += 1
		}
	}

	return ans
}

// ----------------------------------------------------------------------------

type pq []int

func (p pq) Len() int {
	return len(p)
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pq) Less(i, j int) bool {
	return p[i] < p[j] // check for creating min heap
}

func (p *pq) Push(x any) {
	switch y := x.(type) {
	case int:
		*p = append(*p, y)
	default:
		return
	}
}

// Takes out the last element which should have the min number
func (p *pq) Pop() any {
	if len(*p) == 0 {
		return nil
	}

	last := (*p)[len(*p)-1]
	if p.Len() > 0 {
		*p = (*p)[:len(*p)-1]
	} else {
		*p = make(pq, 0)
	}

	return last
}

// Given an integer array nums and an integer k, return the kth largest element
// in the array.
//
// Note that it is the kth largest element in the sorted order, not the kth
// distinct element.
//
// Can you solve it without sorting?
//
// Example 1:
//
// Input: nums = [3,2,1,5,6,4], k = 2
// Output: 5
// Example 2:
//
// Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
// Output: 4
func MinHeap_KthLargest(nums []int, k int) int {
	p := pq{}

	for _, value := range nums {
		heap.Push(&p, value)
		if p.Len() > k {
			_ = heap.Pop(&p)
		}
	}

	return heap.Pop(&p).(int)
}
