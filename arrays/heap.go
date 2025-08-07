package arrays

import (
	"container/heap"
	"sort"
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

func PriorityQueue_SlidingWindowMax(nums []int, k int) []int {
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
		if p.Len() >= k {
			if p[0] < value {
				p[0] = value
				heap.Fix(&p, 0)
			}
			continue
		}
		heap.Push(&p, value)
	}

	return heap.Pop(&p).(int)
}

func QuickSelect_KthLargest(nums []int, k int) int {
	/* Note
	Here we chose the kth element from the end as our pivot.
	With that picked, we create 2 sides:
		- LHS => all elements are smaller than our pivot
		- RHS => all elements are larger than our pivot

	Now if, our RHS's length is == (k-1); then our pivot is exactly Kth largest
	and we must return it.

	If RHS's length is greater than or equal to k, we do a quick select on RHS
	for k'th largest again

	Now, if RHS's length is less than `k`, then it means our required value is in
	LHS at some postion `n`, which can be calculated as
			`n = k - len(rhs) - 1`
	where:
		- We find out how many elements we have already dimmed bigger (len(rhs))
		- We then remove that many number of elements from `k` (k - len(rhs))
		- We then also remove the pivot since it wasn't the element we were looking
		for (k - len(rhs) - 1)
	*/

	x := len(nums) - k
	pivot := nums[x]
	lhs, rhs := []int{}, []int{}

	for i, v := range nums {
		if i == x {
			continue // we ignore our pivot
		}

		if v <= pivot {
			lhs = append(lhs, v)
		} else {
			rhs = append(rhs, v)
		}
	}

	if len(rhs) == k-1 {
		return pivot
	}

	if len(rhs) < k {
		n := k - len(rhs) - 1
		return QuickSelect_KthLargest(lhs, n)
	}

	return QuickSelect_KthLargest(rhs, k)
}

// You are given an array of words where each word consists of lowercase English
// letters.
//
// wordA is a predecessor of wordB if and only if we can insert exactly one letter
// anywhere in wordA without changing the order of the other characters to make
// it equal to wordB.
//
// For example, "abc" is a predecessor of "abac", while "cba" is not a predecessor
// of "bcad".
// A word chain is a sequence of words [word1, word2, ..., wordk] with k >= 1,
// where word1 is a predecessor of word2, word2 is a predecessor of word3, and
// so on. A single word is trivially a word chain with k == 1.
//
// For example, "abc" is a predecessor of "abac", while "cba" is not a predecessor
// of "bcad".
//
// For example, "abc" is a predecessor of "abac", while "cba" is not a predecessor
// of "bcad".
//
// Example 1:
//
//	For example, "abc" is a predecessor of "abac", while "cba" is not a predecessor
//	of "bcad".
//	Input: words = ["a","b","ba","bca","bda","bdca"]
//	Output: 4
//	Explanation: One of the longest word chains is ["a","ba","bda","bdca"].
//
// Constraints:
//
//	1 <= words.length <= 1000
//	1 <= words[i].length <= 16
//	words[i] only consists of lowercase English letters.
func MinHeap_longestStrChain(words []string) int {
	/* Notes

	Here we will use 1DP solution:

	First we will need the words in sorted order.
	Then we create an array where we will store the max depth of the i'th word of
	the words list -> say `dp`

	We initialize the 0th of dp as 1 since the 1st will be the smalled word with
	depth 1.

	Now for every words after that, we will check with all words before that.
	We then pluck the list of words that are its predecessors.

	From them we take the largest depth as our target predecessor's depth and add
	that and +1 as the depth of this word.

	As we do this, we keep track of what was the max depth we found. This is the
	answer we are looking for.
	*/

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	sol := 1
	dp := make([]int, len(words))
	dp[0] = 1

	for i := 1; i < len(words); i++ {
		best := 0

		for j := 0; j <= i; j++ {
			if isPredecessor(words[j], words[i]) && dp[j] > best {
				best = dp[j]
			}
		}

		dp[i] = best + 1
		sol = max(dp[i], sol)
	}

	return sol
}

func isPredecessor(w1, w2 string) bool {
	if len(w1) >= len(w2) || len(w2)-len(w1) != 1 {
		return false
	}

	i, j, n := 0, 0, len(w1)
	for i < n {
		if w1[i] != w2[j] {
			j++
			break
		}
		i++
		j++
	}

	for i < n {
		if w1[i] != w2[j] {
			return false
		}
		i++
		j++
	}

	return true
}
