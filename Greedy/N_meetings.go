package greedy

import "sort"

type pair struct {
	start int
	end   int
}

type pairlist []pair

func (p pairlist) Len() int           { return len(p) }
func (p pairlist) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pairlist) Less(i, j int) bool { return p[i].end < p[j].end }

func N_Meetings(start_times []int, end_times []int) int {
	// we need to pair start and end time
	//Then sort the pairs by end time
	//so implement sort.Sort interface

	//input the the values into a pair and sort the pair
	p := make(pairlist, len(start_times))

	for i, s := range start_times {
		p[i] = pair{s, end_times[i]}
	}
	sort.Sort(p)

	//now loop and count the meetings by comparing prev_end and next meeting start time
	prev_end := 0
	count := 0

	for _, pair := range p {
		if pair.start > prev_end {
			count++
			prev_end = pair.end
		}
	}
	return count
}

// func main() {
// 	start := []int{1, 3, 0, 5, 8, 5}
// 	end := []int{2, 4, 6, 7, 9, 9}

// 	no_of_meeetins := greedy.N_Meetings(start, end)
// 	fmt.Println(no_of_meeetins)
// }
