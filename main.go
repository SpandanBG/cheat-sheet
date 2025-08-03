package main

import (
	"fmt"

	"sudocoding.xyz/cheat-sheet/arrays"
)

func main() {
	fmt.Println("Welcome to Cheat - Sheet")
	fmt.Println("---------------------------------------------")

	fmt.Println("Array / Strings")
	fmt.Println("----------------")

	fmt.Print("\t1.Two-Pointers")
	fmt.Println(`
		(a) K-Sum Problem
		Input:
			inp: [2, 3, 4, 5, 6],
			target: 10`)
	inp := []int{2, 3, 4, 5, 6}
	if i, j, err := arrays.TwoPointers_KSum_SortedArray(inp, 10); err != nil {
		fmt.Println("\t\tError: ", err.Error())
	} else {
		fmt.Println("\t\tOutput: ", inp[i], " and ", inp[j])
	}

	fmt.Println(`
		(b) Duplicate Number Detection - Floyd's Tortoise and Hare Algo
		Input:
			inp: [1, 3, 4, 2, 2]`)
	if arrays.TwoPointers_CycleDetection([]int{1, 3, 4, 2, 2}) {
		fmt.Println("\t\tOutput: Is Cyclic")
	} else {
		fmt.Println("\t\tOuput: Is NOT Cyclic")
	}

	fmt.Println(`
		(c) Backspace string compare - 2 sequence comparision
		Input:
			inp: "a#b#c#d##e#f", "##f"`)
	fmt.Println("\t\tOutput: ", arrays.TwoPointers_TwoSequenceComparing("a#b#c#d##e#f", "##f"))

	fmt.Print("\n\t2.Sliding-Window")
	fmt.Println(`
		(a) Size of largest substring of non-repeating characters
		Inputs: "abcabcbb", "bbbbb", "pwwkew"`)
	fmt.Println(
		"\t\tOutputs: ",
		arrays.SlidingWindow_LongestSubstringWithoutRepeatingChars("abcabcbb"), ",",
		arrays.SlidingWindow_LongestSubstringWithoutRepeatingChars("bbbbb"), ",",
		arrays.SlidingWindow_LongestSubstringWithoutRepeatingChars("pwwkew"),
	)

	fmt.Println("\n\t3.Binary Search")
	fmt.Println("\t\t(a) Median of 2 sorted arrays")
	fmt.Println(
		"\t\tInputs:\n\t\t\ta = [1, 3]\n\t\t\tb = [2]",
		"\n\t\tOutput:",
		arrays.BinarySearch_MedianOfTwoArrays(
			[]int{1, 3},
			[]int{2},
		), "== 2",
	)
	fmt.Println(
		"\t\tInputs:\n\t\t\ta = [1, 2]\n\t\t\tb = [3, 4]",
		"\n\t\tOutput:",
		arrays.BinarySearch_MedianOfTwoArrays(
			[]int{1, 2},
			[]int{3, 4},
		), "== 2.5",
	)
	fmt.Println(
		"\t\tInputs:\n\t\t\ta = [1, 2, 3, 4, 5]\n\t\t\tb = [6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17]",
		"\n\t\tOutput:",
		arrays.BinarySearch_MedianOfTwoArrays(
			[]int{1, 2, 3, 4, 5},
			[]int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
		), "== 9",
	)
	fmt.Println(
		"\t\tInputs:\n\t\t\ta = [1, 2, 3]\n\t\t\tb = [4, 5, 6]",
		"\n\t\tOutput:",
		arrays.BinarySearch_MedianOfTwoArrays(
			[]int{1, 2, 3},
			[]int{4, 5, 6},
		), "== 3.5",
	)

	fmt.Println("\n\t4.Heap")
	fmt.Println("\t\t(a) Sliding Window with Max Heap")
	fmt.Println(
		"\t\tInputs: nums=[1, 3, -1, -3, 5, 3, 6, 7], k=3",
		"\n\t\tOutput:",
		arrays.MaxHeap_SlidingWindowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3),
	)

	fmt.Println("\n\t\t(b) Kth Largest (Min Heap)")
	fmt.Println(
		"\t\tInputs: nums=[3,2,1,5,6,4], k=2",
		"\n\t\tOutput:",
		arrays.MinHeap_KthLargest([]int{3, 2, 1, 5, 6, 4}, 2),
	)

}
