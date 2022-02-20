package algo5

type TopVotedCandidate struct {
	Topvotes []int
	Times    []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	topvotearray := make([]int, len(times))

	highvote := -1
	//the vote in the current time
	personvotes := map[int]int{-1: -1}
	for i, person := range persons {

		personvotes[person]++
		if personvotes[person] >= personvotes[highvote] {
			highvote = person
		}
		topvotearray[i] = highvote
	}
	return TopVotedCandidate{Topvotes: topvotearray, Times: times}
}

func (this *TopVotedCandidate) Q(t int) int {
	//fmt.Println(*this)
	l, r := -1, len(this.Times)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if this.Times[mid] <= t {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return this.Topvotes[r]
	//return this.Topvotes[sort.SearchInts(this.Times, t+1)-1]
}
