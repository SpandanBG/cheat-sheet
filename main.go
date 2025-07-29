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
	i, j, err := arrays.TwoPointers_KSum_SortedArray(inp, 10)
	if err != nil {
		fmt.Println("\t\tError: ", err.Error())
	} else {
		fmt.Println("\t\tOutput: ", inp[i], " and ", inp[j])
	}
}
