package graphs

import (
	"math/big"
)

type Item struct {
	id       big.Int
	city     int
	depth    int
	priority int
}

type MinHeap []*Item

func (m *MinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MinHeap) Less(i, j int) bool {
	return (*m)[i].priority < (*m)[j].priority
}

func (m *MinHeap) Push(i *Item) {
	*m = append(*m, i) // push the item to the back
	m.up(len(*m) - 1)  // shift up the item till in correct position
}

func (m *MinHeap) up(i int) {
	for {
		j := (i - 1) / 2 // parent

		// check if parent is out of bound
		// or parent is smaller
		if j == i || !m.Less(i, j) {
			break // exit
		}

		// Swap and update new current node index
		m.Swap(i, j)
		i = j
	}
}

func (m *MinHeap) Pop() *Item {
	// if nothing in heap
	if len(*m) == 0 {
		return nil // return nothing
	}

	// send smallest item to the back
	m.Swap(0, len(*m)-1)
	popped := (*m)[len((*m))-1] // pull it out

	if len((*m)) == 1 {
		// if ther ewas only one item on  the heap
		// then empty the heap
		*m = []*Item{}
	} else {
		// otherwise - resize the heap by excluding the popped item
		(*m) = (*m)[:len((*m))-1]
		// down shift the root to its correct position
		m.down(0)
	}

	return popped
}

func (m *MinHeap) down(i int) {
	j := 2*i + 1 // left

	for j < len(*m) {
		// check if right is smaller than left
		if j2 := j + 1; j2 < len(*m) && m.Less(j2, 1) {
			j = j2 // pick the smallest to move to top
		}

		if m.Less(i, j) { // if current is already smallest
			break // exit
		}

		// swap and move the current node index to the new position
		m.Swap(i, j)
		i = j
		j = 2*i + 1 // update the next left
	}
}

// There are n cities connected by some number of flights. You are given an array
// flights where flights[i] = [fromi, toi, pricei] indicates that there is a
// flight from city fromi to city toi with cost pricei.
//
// You are also given three integers src, dst, and k, return the cheapest price
// from src to dst with at most k stops. If there is no such route, return -1.
//
// Example 1:
//
//	 Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]],
//		src = 0, dst = 3, k = 1
//	 Output: 700
//	 Explanation:
//	 The graph is shown above.
//	 The optimal path with at most 1 stop from city 0 to 3 is marked in red and
//	 has cost 100 + 600 = 700.
//
//	 Note that the path through cities [0,1,2,3] is cheaper but is invalid
//	 because it uses 2 stops.
//
// Constraints:
//
//	1 <= n <= 100
//	0 <= flights.length <= (n * (n - 1) / 2)
//	flights[i].length == 3
//	0 <= fromi, toi < n
//	fromi != toi
//	1 <= pricei <= 104
//	There will not be any multiple flights between two cities.
//	0 <= src, dst, k < n
//	src != dst
func Dijkstra_CheapestFlightWithinKStops(n int, flights [][]int, src, dst, k int) int {
	/* Notes

	Dijkstra's Algorithm is BFS with Priority Queue. In plain BFS we use a plain
	Queue. Here we will be using a max/min heap to get the next nodes to travel
	to.
	*/

	const (
		from_i = 0
		to_i   = 1
		cost_i = 2
	)

	var root_id big.Int
	root_id.SetBit(&root_id, src, 1)
	root := &Item{city: src, id: root_id}
	srcHeap := &MinHeap{}
	srcHeap.Push(root)

	// convert flights array to map of:
	//  { from_i: [[to_i, price_i], ...] }
	dict := map[int][][]int{}
	for _, flight := range flights {
		data, ok := dict[flight[from_i]]
		if !ok {
			data = make([][]int, 0)
		}

		data = append(data, []int{flight[to_i], flight[cost_i]})
		dict[flight[from_i]] = data
	}

	// to store the lowest travel cost
	var marked *Item

	// a memory of previously traveled city
	// { "srcID_dstID" : *Item }
	memo := map[string]*Item{}

	for len(*srcHeap) > 0 {
		sNode := srcHeap.Pop()
		if sNode == nil {
			break
		}

		delete(memo, sNode.id.String())

		// here we check for depth > k + 1.
		// e.g.: k=1, depth=0,1,2 -> k+1 is valid (one stop at 1), depth=3 -> too may stops
		if sNode.depth > k+1 {
			// we ignore in this case and continue to the next
			// we don't break cause the heap my have higher costs
			// at lower depths that might still reach the dst
			continue
		}

		if sNode.city == dst && (marked == nil || sNode.priority < marked.priority) {
			marked = sNode
			continue
		}

		for _, flight := range dict[sNode.city] {
			cost := flight[1] + sNode.priority
			if marked != nil && marked.priority <= cost {
				continue
			}

			var next_id big.Int
			next_id.SetBit(&sNode.id, flight[0], 1)

			next := &Item{
				id:       next_id,
				city:     flight[0],
				depth:    sNode.depth + 1,
				priority: cost,
			}

			if _, ok := memo[next.id.String()]; ok {
				// we have already visited this. visiting again will be a waste of monies
				continue
			}

			srcHeap.Push(next)
			memo[next.id.String()] = next
		}
	}

	if marked == nil {
		return -1
	}
	return marked.priority
}
