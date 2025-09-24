package main

import (
	"fmt"
	"sort"
)

func main() {
	data := [][]int{{1, 1}, {500000000, 499999999}, {1000000000, 999999998}}
	//data := [][]int{{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1}}
	//data := [][]int{{72, 98}, {62, 27}, {32, 7}, {71, 4}, {25, 19}, {91, 30}, {52, 73}, {10, 9}, {99, 71}, {47, 22}, {19, 30}, {80, 63}, {18, 15}, {48, 17}, {77, 16}, {46, 27}, {66, 87}, {55, 84}, {65, 38}, {30, 9}, {50, 42}, {100, 60}, {75, 73}, {98, 53}, {22, 80}, {41, 61}, {37, 47}, {95, 8}, {51, 81}, {78, 79}, {57, 95}}
	fmt.Println(minimumLines(data))
}

func minimumLines(stockPrices [][]int) int {
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] > stockPrices[j][0]
	})
	prev, count := make([]int, 0), 0
	if len(stockPrices) < 2 {
		return 0
	} else {
		count = 1
	}
	pdx, pdy := 0, 0
	for _, sp := range stockPrices {
		if len(prev) != 2 {
			prev = sp
		} else {
			dx := sp[0] - prev[0]
			dy := sp[1] - prev[1]
			a := dx * pdy
			b := dy * pdx
			if a != b {
				count++
			}
			prev = sp
			pdx, pdy = dx, dy
		}
	}

	return count
}

func minimumLinesBU(stockPrices [][]int) int {
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] > stockPrices[j][0]
	})
	prev, count := make([]int, 0), 0
	if len(stockPrices) < 2 {
		return 0
	} else {
		count = 1
	}
	pdx, pdy := 0, 0
	for _, sp := range stockPrices {
		if len(prev) != 2 {
			prev = sp
		} else {
			dx := sp[0] - prev[0]
			dy := sp[1] - prev[1]
			a := dx * pdy
			b := dy * pdx
			if a != b {
				count++
			}
			prev = sp
			pdx, pdy = dx, dy
		}
	}

	return count
}
