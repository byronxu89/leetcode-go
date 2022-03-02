package algo8

func findRedundantConnection(edges [][]int) []int {
	initset := Construct(len(edges) + 1)
	for _, elem := range edges {
		if !initset.Join(elem[0], elem[1]) {
			return elem
		}
	}
	return nil
}
