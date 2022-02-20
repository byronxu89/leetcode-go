package algo6

func climbStairs(n int) int {
	/* 初始边界设为 爬 0, 1, 2 台阶的情况
	f(x) = f(x-1) + f(x-2)
	*/
	p, q, r := 1, 1, 2
	i := 2
	if n == 0 {
		return p
	}
	if n == 1 {
		return q
	}
	if n == 2 {
		return r
	}
	for i < n {
		tmp := q + r
		p = q
		q = r
		r = tmp
		i++
	}
	return r
}
