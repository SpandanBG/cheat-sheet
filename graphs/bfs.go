package graphs

import (
	"fmt"
	"strings"
)

// On an 2 x 3 board, there are five tiles labeled from 1 to 5, and an empty
// square represented by 0. A move consists of choosing 0 and a 4-directionally
// adjacent number and swapping it.

// The state of the board is solved if and only if the board is [[1,2,3],[4,5,0]].

// Given the puzzle board board, return the least number of moves required so
// that the state of the board is solved. If it is impossible for the state of
// the board to be solved, return -1.

// Example 1:
// Input: board = [[1,2,3],[4,0,5]]
// Output: 1
// Explanation: Swap the 0 and the 5 in one move.

// Example 2:
// Input: board = [[1,2,3],[5,4,0]]
// Output: -1
// Explanation: No number of moves will make the board solved.

// Example 3:
// Input: board = [[4,1,2],[5,0,3]]
// Output: 5
// Explanation: 5 is the smallest number of moves that solves the board.
// An example path:
// After move 0: [[4,1,2],[5,0,3]]
// After move 1: [[4,1,2],[0,5,3]]
// After move 2: [[0,1,2],[4,5,3]]
// After move 3: [[1,0,2],[4,5,3]]
// After move 4: [[1,2,0],[4,5,3]]
// After move 5: [[1,2,3],[4,5,0]]

// Constraints:

// board.length == 2
// board[i].length == 3
// 0 <= board[i][j] <= 5
// Each value board[i][j] is unique.
func BFS_SlidingPuzzle(board [][]int) int {
	/*Note
	Here we will create a BFS tree and save every node's computation to a memory
	(following DP to avoid re-calc and infinite looping).

	At each step we deque a node and create it's left, right, up and down children.
	We enqueue these children if not present on the memory (i.e. already computed
	for).

	If at any point the child is solved we return immediatly.
	*/

	type state struct {
		id       string   // ID of the state
		isEnd    bool     // if true, the board is solved
		depth    int      // BFS search depth
		children []string // IDs of child states
		zPostion int      // position of zero
	}

	memo := map[string]*state{} // memo of visited states for DP

	queue := make([]*state, 0) // queue

	// enqueue
	eque := func(s *state) {
		queue = append(queue, s)
	}

	// dequeue
	dque := func() *state {
		if len(queue) == 0 {
			return nil
		}
		last := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = make([]*state, 0)
		}
		return last
	}

	// move up one step
	up := func(i int) (int, bool) {
		return i - 3, i-3 > -1
	}

	// move down one step
	down := func(i int) (int, bool) {
		return i + 3, i+3 < 6
	}

	// move left one step
	left := func(i int) (int, bool) {
		return i - 1, (i > 2 && i-1 > 2) || (i <= 2 && i-1 > -1)
	}

	// move right one step
	right := func(i int) (int, bool) {
		return i + 1, (i > 2 && i+1 < 6) || (i <= 2 && i+1 < 3)
	}

	// checks if the ID is of solved state
	isSolved := func(id string) bool {
		return id == "123450"
	}

	swap := func(id string, i, j int) string {
		rid := []rune(id)
		rid[i], rid[j] = rid[j], rid[i]
		return string(rid)
	}

	swapNcreate := func(parent *state, move func(int) (int, bool)) bool {
		i, ok := move(parent.zPostion)
		if !ok {
			return false
		}

		childID := swap(parent.id, parent.zPostion, i)
		isEnd := isSolved(childID)
		if isEnd {
			return true
		}

		child := &state{
			id:       childID,
			isEnd:    isEnd,
			depth:    parent.depth + 1,
			zPostion: i,
		}

		// add to parent's child and queue if not visited
		if _, ok := memo[childID]; !ok {
			parent.children = append(parent.children, childID)
			eque(child)
		}

		return child.isEnd
	}

	rootID := fmt.Sprintf("%d%d%d%d%d%d", board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2])
	isEnd := isSolved(rootID)
	if isEnd {
		return 0
	}

	root := &state{
		id:       rootID,
		isEnd:    isEnd,
		depth:    0,
		children: make([]string, 0),
		zPostion: strings.Index(rootID, "0"),
	}

	eque(root)

	for len(queue) != 0 {
		node := dque()
		if node == nil {
			break
		}

		if swapNcreate(node, left) ||
			swapNcreate(node, up) ||
			swapNcreate(node, right) ||
			swapNcreate(node, down) {
			return node.depth + 1
		}

		memo[node.id] = node
	}

	return -1
}

func BiDirBFS_SlidingPuzzle(board [][]int) int {
	/*Note
	Here we are going to follow a bidirectional BFS:
		- We will start a BFS from our source (board)
		- We will start a BFS from our goal (123450)
		- If a node from either already was added to the memo by the other, then we
		have found our connecting state
	*/

	type state struct {
		id       string   // ID of the state
		isTop    bool     // If true this node was added from the source side
		depth    int      // BFS search depth
		children []string // IDs of child states
		zPostion int      // position of zero
	}

	memo := map[string]*state{} // memo of visited states for DP

	newQ := func() (func(*state), func() *state) {
		queue := make([]*state, 0) // queue

		// enqueue
		eque := func(s *state) {
			queue = append(queue, s)
		}

		// dequeue
		dque := func() *state {
			if len(queue) == 0 {
				return nil
			}
			last := queue[0]
			if len(queue) > 1 {
				queue = queue[1:]
			} else {
				queue = make([]*state, 0)
			}
			return last
		}

		return eque, dque
	}

	// source queue for BFS
	sEque, sDque := newQ()

	// goal queue for BFS
	gEque, gDque := newQ()

	// move up one step
	up := func(i int) (int, bool) {
		return i - 3, i-3 > -1
	}

	// move down one step
	down := func(i int) (int, bool) {
		return i + 3, i+3 < 6
	}

	// move left one step
	left := func(i int) (int, bool) {
		return i - 1, (i > 2 && i-1 > 2) || (i <= 2 && i-1 > -1)
	}

	// move right one step
	right := func(i int) (int, bool) {
		return i + 1, (i > 2 && i+1 < 6) || (i <= 2 && i+1 < 3)
	}

	swap := func(id string, i, j int) string {
		rid := []rune(id)
		rid[i], rid[j] = rid[j], rid[i]
		return string(rid)
	}

	swapNcreate := func(parent *state, move func(int) (int, bool)) int {
		i, ok := move(parent.zPostion)
		if !ok {
			return -1
		}

		childID := swap(parent.id, parent.zPostion, i)
		child := &state{
			id:       childID,
			isTop:    parent.isTop,
			depth:    parent.depth + 1,
			zPostion: i,
		}

		// add to parent's child and queue if not visited
		if fNode, ok := memo[childID]; !ok {
			parent.children = append(parent.children, childID)

			if parent.isTop {
				sEque(child)
			} else {
				gEque(child)
			}
		} else if child.isTop != fNode.isTop {
			// returns depth if child's parent and fNode's parents are differnt
			// i.e. they have connected
			return fNode.depth + child.depth
		}

		return -1
	}

	rootID := fmt.Sprintf("%d%d%d%d%d%d", board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2])
	goalID := "123450"
	if rootID == goalID {
		return 0
	}

	root := &state{
		id:       rootID,
		isTop:    true,
		depth:    0,
		children: make([]string, 0),
		zPostion: strings.Index(rootID, "0"),
	}

	goal := &state{
		id:       goalID,
		isTop:    false,
		depth:    0,
		children: make([]string, 0),
		zPostion: 5,
	}

	sEque(root)
	gEque(goal)

	for true {
		sNode := sDque()
		if sNode == nil {
			break
		}

		depth := max(swapNcreate(sNode, left), -1)
		depth = max(swapNcreate(sNode, right), depth)
		depth = max(swapNcreate(sNode, up), depth)
		depth = max(swapNcreate(sNode, down), depth)
		if depth > -1 {
			return depth
		}

		memo[sNode.id] = sNode

		// ----------------- upside umop

		gNode := gDque()
		if gNode == nil {
			continue
		}

		depth = max(swapNcreate(gNode, left), -1)
		depth = max(swapNcreate(gNode, right), depth)
		depth = max(swapNcreate(gNode, up), depth)
		depth = max(swapNcreate(gNode, down), depth)
		if depth > -1 {
			return depth
		}

		memo[gNode.id] = gNode

	}

	return -1
}
