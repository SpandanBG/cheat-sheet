package graphs

import "math"

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
func BellmanFord_CheapestFlightWithinKStops(n int, flights [][]int, src, dst, k int) int {
	/* Notes

	Bellman Ford algo is ideal for graphs with negative weights. The reason why
	Dijkstra's don't work well is cause it doesn't revisit nodes - which can cut
	off possible solution with better results with node revisits. That means a
	longer path with smaller cost might get left out.

	In this cheapest flights within K stops is also perfect for Bellman Ford, since
	taking a longer route might lead to cheaper cost.

	Bellman Ford is also suitable cause of the constraints: n < 101 and presence
	of k as the limit to depth.

	Bellman Ford Algo:

	Principle of relaxation of edges:
		- Relaxation means updating the shortest part to a node is found through
		another node. (Here a cheaper path is found through another node).
		- For an edge (u->v) with weight w:
			- If u->v gives a shorted (cheaper) path to v from the src node (dist[v] >
			dist[u] + w); then we update dist[v] = dist[u] + w.
		- This process is repeated `n-1` times (n is the number of nodes) for all
		edges

	Why it works:
	The shortest (cheapest) path between 2 nodes is at most `n-1` edges. More than
	that means there would be a cycle. As such relaxing the edges `n-1` times
	ensures that all possible paths between source and any other node has been
	covered.

	Steps:
	- Create a distance array (dist) from src to every other node.
	- Mark src->src dist as 0 and rest as infinity.
	- For n-1 iterations:
		- Make a copy of dist -> dist2
		- Iterate through each edge (flights -> flight):
			- u = distance to u from src (dist[u])
			- v = distance to v from src (dist[v])
			- wt = weight of u->v given on the edge data
			- If u is infinity -> then we don't know the connection between src to u yet
			and we skip the next steps and move on to the next edge
			- If dist[v] > dist[u] + wt: we update dist2[v] = dist[u] + wt
	- At the end we get the result for src->dst's cost

	Notice: Here, we check the older list and update a copy of the list, this is
	cause in an iteration only one step is taken. Imagine this to be time frozen
	activity and we take every possible step and remember the values. Thus, we
	need to do the steps in the frozen set of values.
		When we are to remember the update values in the frozen calculation step, we
	must account for previously remmebered value in that frozen time; because if
	we make more than 1 step to the same destination (i.e. x->y and a->y) and both
	gives smaller cost, we must always remember the smaller one.

	Here we can imagin that one complete iteration of the edges is one stop.

	Therefore, for our problem statement -> we only need to do up to (k + 1)
	iterations, i.e. if k=1 we can move 2 times since we can stop 1 time (k=1).
	*/

	// Creating the initial distance map
	dist := make([]int, n)
	dist_alt := make([]int, n)
	for i := range n {
		if i == src {
			dist[i] = 0
			dist_alt[i] = 0
			continue
		}
		dist[i] = math.MaxInt
		dist_alt[i] = math.MaxInt
	}

	// k+1 since non-inclusive range - [0,1]
	// each iteration is 1 step.
	// if k=1, then range is [0,1] therefore 1 stop
	for range k + 1 {
		for _, edge := range flights {
			u := dist[edge[0]]
			v := dist[edge[1]]
			w := edge[2]

			if u != math.MaxInt && v > (u+w) {
				dist_alt[edge[1]] = min(u+w, dist_alt[edge[1]])
			}
		}

		copy(dist, dist_alt)
	}

	if dist[dst] == math.MaxInt {
		return -1
	}
	return dist[dst]
}
