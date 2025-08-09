package arrays

import (
	"errors"
	"math"
)

// TwoPointers_KSum_SortedArray - Takes in a sorted array and finds 2 indices in
// that array whose value would sum up to match the target value.
//
// This method uses the `Two Pointer` strategry to solve the `K Sum` problem.
func TwoPointers_KSum_SortedArray(inputs []int, target int) (int, int, error) {
	left, right := 0, len(inputs)-1
	for right > left {
		if inputs[left]+inputs[right] == target {
			break
		} else if inputs[left]+inputs[right] <= target {
			// since the sum is to be bumpped up, we move the left pointer to bump it
			// up by the smallest amount
			left += 1
		} else if inputs[left]+inputs[right] >= target {
			// since the sum is to be bumpped down, we move the right pointer to bump
			// it down by the largest amount
			right -= 1
		}
	}

	if inputs[left]+inputs[right] == target {
		return left, right, nil
	}
	return -1, -1, errors.New("No such combination found that results in target")
}

// TwoPointers_CycleDetection - Uses 2 pointers (specifically the Floyd's The
// Tortoise and The Hare algorithm) to solve the following problem:
//
// Problem: Find the Duplicate Number
// You are given an array of n + 1 integers where each integer is in the range
// [1, n] inclusive.
//
// There is only one repeated number in the array, but it could be repeated more
// than once.
//
// Return the duplicate number without modifying the array and using only
// constant extra space.
//
// Constraints:
//   - 1 <= n <= 10⁵
//   - The array contains n + 1 integers.
//   - All integers are in the range [1, n].
//   - Only one integer is repeated, but it may appear more than once.
//   - You must not modify the array (e.g., no sorting or replacing values).
//   - You must use only constant extra space.
func TwoPointers_CycleDetection(inp []int) bool {
	/*
		 Notes:
		 	The Floyd's Tortoise and Hare algorithm works as such:
				- We have 2 pointes: `slow` and `hare`
				- `slow` moves 1 step at a time
				- `hare` moves 2 steps at a time
				- if there is a cycle the `fast` meets the `slow`
				- if there is non - then `fast` goes out of bound / end of linked list

			Here the provided array can essentially be looked at as an linked list:

			Eg: [1, 3, 4, 2, 2] - if we were to use the value as the index of the next
			then, this can be converted to:
					(1)->(3)->(2)->(4)
										 ^____|

			As such we can use the formula:
				`s = inp[s] & f = inp[inp[f]]`
			which implies:
				- `s`is moving 1 step in the linked list
				- `f` is moving 2 steps in the linked list
	*/

	s, f := 0, 0
	for f < len(inp) && s < len(inp) {
		s = inp[s]
		f = inp[inp[f]]

		if s == f {
			return true
		}
	}
	return false
}

// TwoPointers_TwoSequenceComparing - Using 2 pointer solve the following:
//
// Problem: Backspace String Compare
// Given two strings s and t, return true if they are equal when both are typed
// into empty text editors. '#' means a backspace character.
//
// Note: After processing, both strings may be empty.
//
// Constraints:
// 1 <= s.length, t.length <= 200
//
// s and t only contain lowercase letters and the '#' character.
//
// Follow-up:
// Can you solve it in O(n) time and O(1) space?
func TwoPointers_TwoSequenceComparing(a, b string) bool {
	/* Note:
	Here we start from back of both the strings. This is because the characters
	encountered will be confirmed parts of the string.

	In each iteration we will move backwards 1 step and do the following:
		- if `#` is seen: increment a `skip` counter
		- if `skip > 0`: decrement `skip` by 1
		- else compare a[i] == b[j]
	*/

	findNext := func(str string, from int) (int, rune) {
		var char rune
		skip := 0

		for from >= 0 {
			for from >= 0 {
				char = []rune(str)[from]
				if char == '#' {
					skip += 1
					from -= 1
					continue
				}
				break
			}

			if skip > 0 {
				skip -= 1
				from -= 1
				continue
			}

			break
		}

		char = rune(0)
		if from >= 0 {
			char = []rune(str)[from]
		}

		return from - 1, char
	}

	i, j := len(a)-1, len(b)-1
	var ax, bx rune
	for i >= 0 || j >= 0 {
		i, ax = findNext(a, i)
		j, bx = findNext(b, j)

		if ax != bx {
			return false
		}
	}

	return true
}

// Given n non-negative integers representing an elevation map where the width
// of each bar is 1, compute how much water it can trap after raining.
//
// Example: 1
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array
// [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section)
// are being trapped.
//
// Constraints:
//
//	n == height.length
//	1 <= n <= 2 * 104
//	0 <= height[i] <= 105
func TwoPointers_trap(height []int) int {
	/* Note:

	This problem can be solved by jumping between prefix and suffix look up.

	1. Trim off the edges of the array to remove all hights that do not lead into
	smaller hight. E.G.: [0,2,1,2,3,1] => [2,1,2,3] since [0,2] are increasing and
	[3,1] are decreasing left and right edges respectively.

	2. If Left height < Right height; we start collecting water till we reach a
	height greater or equal to the left height. We send the remaining height back
	again to be collected

	3. If Left height >= Right height; we start collectin from the RIGHT till we
	reach a hight greater or equal to the right and send remaining heights back
	again to be collected
	*/

	if len(height) <= 1 {
		return 0
	}

	// shrink left corner to suitable size
	i, j := 0, 1
	for j < len(height) && height[i] <= height[j] {
		i = j
		j += 1
	}

	//shrink right corner to suitable size
	l := len(height) - 1
	k := l - 1
	for k >= 0 && height[k] >= height[l] {
		l = k
		k -= 1
	}

	// we have completely removed all possible captures
	if k <= i {
		return 0
	}

	height = height[i : l+1]

	if len(height) < 3 {
		// no water can be saved with just 2 towers
		return 0
	}

	// --------------------

	i, j = 0, len(height)-1

	// if right edge is smaller than left
	// then we collect water from the right side till we reach equal or greater
	// height, and send the rest to be trapped again to the function
	if height[i] > height[j] {
		collect := 0

		for m := j - 1; i < m; m -= 1 {
			if height[m] >= height[j] {
				return ((j-m-1)*height[j] - collect) + TwoPointers_trap(height[:m+1])
			}
			collect += height[m]
		}

		return (j-i-1)*height[j] - collect
	}

	// otherwise - we collect from the left till we reach equal or greater height
	// and send the rest to be trapped again to the function
	collect := 0
	for m := i + 1; m < j; m += 1 {
		if height[m] >= height[i] {
			return ((m-i-1)*height[i] - collect) + TwoPointers_trap(height[m:])
		}
		collect += height[m]
	}

	return (j-i-1)*height[i] - collect
}

func DP_coinChange(coins []int, amount int) int {
	return walk(coins, amount, map[int]int{0:0})
}

func walk(coins[]int, amount int, memo map[int]int) int {
	if v, ok := memo[amount]; ok {
		return v
	}

	minTake := math.MaxInt
	for _, c := range coins {
		if c > amount {
			continue
		}

		take := walk(coins, amount-c, memo)
		if take == -1 {
			continue
		}

		minTake = min(minTake, 1+take)
	}

	if minTake == math.MaxInt {
		memo[amount] = -1
		return -1
	}
	memo[amount] = minTake
	return minTake
}
