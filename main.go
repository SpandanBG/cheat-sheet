package main

import (
	"fmt"

	"sudocoding.xyz/cheat-sheet/arrays"
	"sudocoding.xyz/cheat-sheet/graphs"
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

	fmt.Println(`
		(d) Trapping Rain Water
		Input: height = [ 5,4,1,2 ]`)
	fmt.Println("\t\tOutput:", arrays.TwoPointers_trap([]int{5, 4, 1, 2}))

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

	fmt.Println("\n\t4.Heap | Priority Queue | Quickselect")
	fmt.Println("\t\t(a) Max in Sliding Window with Max Heap")
	fmt.Println(
		"\t\tInputs: nums=[1, 3, -1, -3, 5, 3, 6, 7], k=3",
		"\n\t\tOutput:",
		arrays.MaxHeap_SlidingWindowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3),
	)

	fmt.Println("\n\t\t(b) Max in Sliding Window with Priority Queue")
	fmt.Println(
		"\t\tInputs: nums=[1, 3, -1, -3, 5, 3, 6, 7], k=3",
		"\n\t\tOutput:",
		arrays.PriorityQueue_SlidingWindowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3),
	)

	fmt.Println("\n\t\t(c) Kth Largest (Min Heap)")
	fmt.Println(
		"\t\tInputs: nums=[3,2,1,5,6,4], k=2",
		"\n\t\tOutput:",
		arrays.MinHeap_KthLargest([]int{3, 2, 1, 5, 6, 4}, 2),
	)

	fmt.Println("\n\t\t(d) Kth Largest (Quick Select)")
	fmt.Println(
		"\t\tInputs: nums=[3,2,1,5,6,4], k=2",
		"\n\t\tOutput:",
		arrays.QuickSelect_KthLargest([]int{3, 2, 1, 5, 6, 4}, 2),
	)

	fmt.Println("\n\t\t(e) Longest String Chain")
	fmt.Println(
		"\t\tInputs: nums=['a','b','ba','bca','bda','bdca']",
		"\n\t\tOutput:",
		arrays.MinHeap_longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}),
	)

	fmt.Println("--------------------------------------------------------------")

	fmt.Println("Graphs")
	fmt.Println("----------------")

	fmt.Print("\t1. Breadth-First-Search")
	fmt.Println(`
		(a) Minimum Steps - Sliding Puzzle
		Input:
			board: [[3, 2, 4 , [1, 5, 0]]`)
	fmt.Println("\t\tPlain BFS Output:", graphs.BFS_SlidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}}))
	fmt.Println("\t\tBi-Directional BFS Output:", graphs.BiDirBFS_SlidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}}))

	fmt.Println(`
		(b) Shortest Path Visiting all nodes - Sliding Puzzle
		Input:
			board: [[1,2,3],[0],[0],[0]]`)
	fmt.Println("\t\tOutput:", graphs.BFS_ShortestPathAllNodes([][]int{{1, 2, 3}, {0}, {0}, {0}}))
	fmt.Println("\t\tOutput:", graphs.BFS_ShortestPathAllNodes([][]int{{7}, {3}, {3, 9}, {1, 2, 4, 5, 7, 11}, {3}, {3}, {9}, {3, 10, 8, 0}, {7}, {11, 6, 2}, {7}, {3, 9}}))

	fmt.Print("\n\t2. Dijkstra's")
	fmt.Println(`
		(a) Cheapest Flights Within K Stops 
		Input:
			n = 4, 
			flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], 
			src = 0, dst = 3, k = 1`)
	fmt.Println("\t\tOutput:", graphs.Dijkstra_CheapestFlightWithinKStops(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1))

	fmt.Print("\n\t3. Bellman Ford's")
	fmt.Println(`
		(a) Cheapest Flights Within K Stops 
		Input:
			n = 5, 
			flights = {{1,0,5},{2,1,5},{3,0,2},{1,3,2},{4,1,1},{2,4,1}}, 
			src = 2, dst = 0, k = 2`)
	fmt.Println("\t\tOutput:", graphs.BellmanFord_CheapestFlightWithinKStops(5, [][]int{{1, 0, 5}, {2, 1, 5}, {3, 0, 2}, {1, 3, 2}, {4, 1, 1}, {2, 4, 1}}, 2, 0, 2))

	fmt.Print("\n\t4. Depth-First Search")
	fmt.Println(`
		(a) Number Of Islands
		Input: grid = [
				["1","1","1","1","0"],
				["1","1","0","1","0"],
				["1","1","0","0","0"],
				["0","0","0","0","0"]
			]`)
	fmt.Println("\t\tOutput:", graphs.DFS_NumberOfIslands([][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}))

	fmt.Println(`
		(b) All Paths From Source to Target
		Input: graph = [[4,3,1],[3,2,4],[],[4],[]]`)
	fmt.Println("\t\tOutput:", graphs.DSF_allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {}, {4}, {}}))

	fmt.Println(arrays.DP_coinChange([]int{186,419,83,408}, 6249))
}
