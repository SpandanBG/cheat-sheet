package arrays

import "errors"

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
