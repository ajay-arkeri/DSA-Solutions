package greedy

import (
	"math"
	"sort"
)

// GFG - chocolate distribution Problem  -->https://www.geeksforgeeks.org/chocolate-distribution-problem/

func MinimumDiff(M int, a []int) int {
	sort.Ints(a)

	i, j := 0, M-1
	min := math.MaxInt
	for j < len(a) {
		diff := a[j] - a[i]
		i++
		j++
		min = int(math.Min(float64(min), float64(diff)))
	}
	return min
}
