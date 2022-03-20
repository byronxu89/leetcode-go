package algo10

import (
	"math"
	"math/rand"
)

const (
	maxlevel = 16
	maxRand  = 65535.0
)

type SkipNode struct {
	val   int
	right *SkipNode
	down  *SkipNode
}
type Skiplist struct {
	head *SkipNode
}

func dial() int {
	return maxlevel - int(math.Log2(1.0+maxRand*rand.Float64()))
	// [ 1....16]
	// 相当于 第一层 取 N/2 第二层取N/4 .....
}

func Constructor() Skiplist {
	left := make([]*SkipNode, maxlevel)
	right := make([]*SkipNode, maxlevel)
	for i := 0; i < maxlevel; i++ {
		left[i] = &SkipNode{-1, nil, nil}
		right[i] = &SkipNode{20001, nil, nil}
	}
	for i := maxlevel - 2; i >= 0; i-- {
		left[i].right = right[i]
		left[i].down = left[i+1]
		right[i].down = right[i+1]
	}
	left[maxlevel-1].right = right[maxlevel-1]

	return Skiplist{left[0]}
}

func (this *Skiplist) Search(target int) bool {
	node := this.head
	for node != nil {
		if node.right.val > target {
			node = node.down
		} else if node.right.val < target {
			node = node.right
		} else {
			return true
		}
	}
	return false

}

func (this *Skiplist) Add(num int) {
	prev := make([]*SkipNode, maxlevel)
	i := 0
	node := this.head
	for node != nil {
		if node.right.val >= num {
			//go down
			//insert here
			prev[i] = node
			i++

			node = node.down
		} else {
			node = node.right
		}
	}
	n := dial()
	arr := make([]*SkipNode, n)
	//从上往下插
	t := &SkipNode{-1, nil, nil}
	for i, a := range arr {

		p := prev[maxlevel-n+i]

		a = &SkipNode{num, p.right, nil}
		p.right = a
		t.down = a
		t = a

	}
}

func (this *Skiplist) Erase(num int) bool {
	ans := false
	node := this.head
	for node != nil {
		if node.right.val > num {
			node = node.down
		} else if node.right.val < num {
			node = node.right
		} else {
			ans = true
			node.right = node.right.right
			node = node.down
		}
	}
	return ans

}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
