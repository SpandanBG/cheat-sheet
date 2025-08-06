package graphs

import (
	"math"
)

type Item struct {
	city     int8
	depth    int8
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
		if j2 := j + 1; j2 < len(*m) && m.Less(j2, j) {
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

	Here we will be using Dijkstra's algorithm (BFS+PQ) to find the cheapest path
	with K stops.

	However, since this problem can have a longer path with cheaper cost, we
	cannot rely on first found destination.

	For this we will have to use The Principle Of Relaxing Edges.
	This principle states that, if dist[v] gives the distance of v form the source
	and a edge u->v exists with weigth wt, such that dist[v] > dist[u] + wt, then
	we update the dist[v] = dist[u] + wt

	In this problem statement, since we are limited to k stops, we will relax the
	edge for node in that step.

	Example: for node n, in x step - it contained a cost A; and we find another
	direction to n in the same step with cost B, such that B < A, then we update
	the dist[n] = B.
	*/

	// min heap for BFS+PQ
	pq := make(MinHeap, 0)

	// memory for edge relaxation
	dist := make([][]int, n)
	dist_alt := make([][]int, n)
	for i := range n {
		// if k stops can be made, then k+1steps can be taken
		dist[i] = make([]int, k+2)
		dist_alt[i] = make([]int, k+2)

		for j := range dist[i] {
			dist[i][j] = math.MaxInt
			dist_alt[i][j] = math.MaxInt
		}
	}

	// Set distance cost from source to source at 0th step as 0
	dist[src][0] = 0
	dist_alt[src][0] = 0

	// push source to pq
	pq.Push(&Item{
		city:     int8(src),
		depth:    0,
		priority: 0,
	})

	// move flights array to map
	fMap := map[int8][][]int{}
	for _, flight := range flights {
		if _, ok := fMap[int8(flight[0])]; !ok {
			fMap[int8(flight[0])] = make([][]int, 0)
		}
		fMap[int8(flight[0])] = append(fMap[int8(flight[0])], []int{flight[1], flight[2]})
	}

	// Perfrom Dijkstra's
	for len(pq) > 0 {
		// Pick node with lowest cost
		sNode := pq.Pop()
		if sNode == nil {
			break
		}

		if sNode.depth == int8(k)+1 {
			// max steps taken -> skip node
			continue
		}

		dist_v := dist[sNode.city][sNode.depth]
		if dist_v != math.MaxInt && dist_v < sNode.priority {
			// if this node was previously reached in the same number of steps at
			// cheaper cost -> then ignore
			continue
		}

		// Push valid next nodes
		for _, flight := range fMap[sNode.city] {
			nCity := int8(flight[0])
			nDepth := sNode.depth + 1
			nPriority := sNode.priority + flight[1]

			if nDepth == int8(k)+1 && nCity != int8(dst) {
				// we have reached max depth and not the destination -> skip
				continue
			}

			// get the distance cost of next city from source city
			dist_v := dist[nCity][nDepth]

			if dist_v != math.MaxInt16 && dist_v < nPriority {
				// this city has been visted from the source previously and the cost
				// was cheaper -> we skip entering here
				continue
			}

			dist_alt[nCity][nDepth] = nPriority

			pq.Push(&Item{
				city:     nCity,
				depth:    nDepth,
				priority: nPriority,
			})
			continue
		}

		for i := range dist_alt {
			copy(dist[i], dist_alt[i])
		}
	}

	minCost := math.MaxInt
	for _, cost := range dist[dst] {
		if cost < minCost {
			minCost = cost
		}
	}

	if minCost == math.MaxInt {
		return -1
	}
	return minCost
}
