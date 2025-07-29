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
}
