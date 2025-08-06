package graphs

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and
// '0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent lands
// horizontally or vertically. You may assume all four edges of the grid are all
// surrounded by water.
//
// Example 1:
//
//	 Input: grid = [
//		 ["1","1","1","1","0"],
//		 ["1","1","0","1","0"],
//		 ["1","1","0","0","0"],
//		 ["0","0","0","0","0"]
//	 ]
//	 Output: 1
//
// Example 2:
//
//	 Input: grid = [
//		 ["1","1","0","0","0"],
//		 ["1","1","0","0","0"],
//		 ["0","0","1","0","0"],
//		 ["0","0","0","1","1"]
//	 ]
//	 Output: 3
//
// Constraints:
//
//	m == grid.length
//	n == grid[i].length
//	1 <= m, n <= 300
//	grid[i][j] is '0' or '1'.
func DFS_NumberOfIslands(grid [][]byte) int {
	/* Notes

	Depth-First Search - Here we go to the deepest part of a branch first and then
	back track. That is -> we would keep pushing items into the stack, till no more
	can be found in that direction. Then we pop the latest item and check the next
	direction of that item and procceed to flow the deepest path in the original
	direction.

	In this problem statement - we will have to move in the cardinal directions.
	For we will select [up, left, down, right] in that sequence of order for each
	node we push into the state.

	If we find a valid node in any of those directions, we will push that node and
	procceed with that node in the [up, left, down, right] directions.

	When all the directions are check-for and popped off / empty, we will pop off
	the node and mark it as already visited.

	When the stack is empty - we will count it as 1 island.
	*/

	stack := make([][2]int, 0)

	push := func(i [2]int) {
		stack = append(stack, i)
	}

	pop := func() ([2]int, bool) {
		if len(stack) == 0 {
			return [2]int{}, false
		}

		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return last, true
	}

	islands := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 1 {
				continue
			}

			islands += 1
			push([2]int{i, j})
			grid[i][j] = 0

			for len(stack) > 0 {
				last := stack[len(stack)-1]
				last_i, last_j := last[0], last[1]

				// check up
				if last_i-1 >= 0 && grid[last_i-1][last_j] == 1 {
					grid[last_i-1][last_j] = 0
					push([2]int{last_i - 1, last_j})
					continue
				}

				// check left
				if last_j-1 >= 0 && grid[last_i][last_j-1] == 1 {
					grid[last_i][last_j-1] = 0
					push([2]int{last_i, last_j - 1})
					continue
				}

				// check down
				if last_i+1 < len(grid) && grid[last_i+1][last_j] == 1 {
					grid[last_i+1][last_j] = 0
					push([2]int{last_i + 1, last_j})
					continue
				}

				// check right
				if last_j+1 < len(grid[0]) && grid[last_i][last_j+1] == 1 {
					grid[last_i][last_j+1] = 0
					push([2]int{last_i, last_j + 1})
					continue
				}

				pop()
			}
		}
	}

	return islands
}
