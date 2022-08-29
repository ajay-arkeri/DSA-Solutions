package greedy

import "sort"

// candies problem from GFG --> Find the minimum and maximum amount to buy all N candies

func Amount(arr []int, N int, k int) (int, int) {
	sort.Ints(arr)
	max := 0
	min := 0

	//minimum
	buy := 0
	free := N - 1

	for buy <= free {
		min = min + arr[buy]
		buy++
		free -= k
	}

	//Maximum
	buy = N - 1
	free = 0

	for buy >= free {
		max = max + arr[buy]
		buy--
		free += k
	}

	return min, max

}

//time = O(NLogN)
//space O(1)
