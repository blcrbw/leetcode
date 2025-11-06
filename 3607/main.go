package main

import "container/heap"

// 3607. Power Grid Maintenance
//
// You are given an integer c representing c power stations, each with a unique identifier id from 1 to c (1‑based indexing).
// These stations are interconnected via n bidirectional cables, represented by a 2D array connections, where each element
// connections[i] = [ui, vi] indicates a connection between station ui and station vi. Stations that are directly or
// indirectly connected form a power grid.
//
// Initially, all stations are online (operational).
// You are also given a 2D array queries, where each query is one of the following two types:
//    [1, x]: A maintenance check is requested for station x. If station x is online, it resolves the check by itself.
//   		  If station x is offline, the check is resolved by the operational station with the smallest id in the same
//  	 	  power grid as x. If no operational station exists in that grid, return -1.
//    [2, x]: Station x goes offline (i.e., it becomes non-operational).
// Return an array of integers representing the results of each query of type [1, x] in the order they appear.
//
// Note: The power grid preserves its structure; an offline (non‑operational) node remains part of its grid and taking it
// offline does not alter connectivity.

type Station struct {
	id       int
	isOnline bool
	heapIdx  int
	grid     *Grid
}

type Grid struct {
	onlineStations Heap
}

type Heap []*Station

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i].isOnline == h[j].isOnline {
		return h[i].id < h[j].id
	}
	return h[i].isOnline == true && h[j].isOnline == false
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}

func (h *Heap) Push(x any) {
	item := x.(*Station)
	item.heapIdx = len(*h)
	*h = append(*h, item)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	var res []int
	stations := make(map[int]*Station)

	for i := 1; i <= c; i++ {
		// Create station objects
		s := Station{
			id:       i,
			isOnline: true,
		}
		stations[i] = &s
	}

	for _, conn := range connections {
		// Connect stations into the grids
		u := stations[conn[0]]
		v := stations[conn[1]]
		if u.grid == nil && v.grid == nil {
			// Create new grid with u and v.
			grid := newGrid(u, v)
			u.grid = grid
			v.grid = grid
		} else if u.grid == nil && v.grid != nil {
			// Add u into the v grid
			heap.Push(&v.grid.onlineStations, u)
			u.grid = v.grid
		} else if u.grid != nil && v.grid == nil {
			heap.Push(&u.grid.onlineStations, v)
			v.grid = u.grid
		} else if u.grid != nil && v.grid != nil && u.grid != v.grid {
			// Merge grids into min grid.
			// Reassign all stations to the merged grid
			mergeGrids(u.grid, v.grid)
		}
	}

	for _, q := range queries {
		station := stations[q[1]]
		if q[0] == 1 {
			if station.isOnline {
				res = append(res, station.id)
			} else {
				if station.grid != nil {
					minGrid := station.grid.onlineStations[0]
					if minGrid.isOnline {
						res = append(res, minGrid.id)
					} else {
						res = append(res, -1)
					}
				} else {
					res = append(res, -1)
				}
			}
		} else if q[0] == 2 {
			station.isOnline = false
			if station.grid != nil {
				heap.Fix(&station.grid.onlineStations, station.heapIdx)
			}
		}
	}

	return res
}

func newGrid(u, v *Station) *Grid {
	h := make(Heap, 0, 2)
	heap.Init(&h)
	heap.Push(&h, u)
	heap.Push(&h, v)
	return &Grid{onlineStations: h}
}

func mergeGrids(grid1, grid2 *Grid) {
	var grid, obsolete *Grid
	if grid1.onlineStations.Len() > grid2.onlineStations.Len() {
		grid, obsolete = grid1, grid2
	} else {
		grid, obsolete = grid2, grid1
	}

	for _, st := range obsolete.onlineStations {
		heap.Push(&grid.onlineStations, st)
		st.grid = grid
	}

}
