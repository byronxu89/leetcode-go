package algo8

type DisjointSet struct {
	fa []int
}

func Construct(n int) DisjointSet {
	s := DisjointSet{fa: make([]int, n)}
	for i, _ := range s.fa {
		s.fa[i] = i
	}
	return s
}
func (s *DisjointSet) Find(x int) int {
	if s.fa[x] != x {
		s.fa[x] = s.Find(s.fa[x])
	}
	return s.fa[x]
}
func (s *DisjointSet) Join(x, y int) bool {
	x, y = s.Find(x), s.Find(y)
	if x != y {
		s.fa[x] = y
		return true
	}
	return false

}
