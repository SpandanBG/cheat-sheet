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
