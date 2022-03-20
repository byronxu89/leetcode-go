package algo10

import (
	"fmt"
	"sort"
)

type Node struct {
	val  int
	idx  int
	pre  *Node
	next *Node
}

func AddNode(array []int, node *Node, idx int) *Node {
	newNode := Node{
		val: array[idx],
		idx: idx,
	}
	node.next.pre = &newNode
	newNode.next = node.next
	newNode.pre = node
	node.next = &newNode
	return &newNode
}
func DeleteNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func main() {
	length := 0

	fmt.Scanf("%d", &length)
	arr := make([]int, length)
	rk := make([]int, length)
	ans := make([]int, length)
	pos := make([]*Node, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &arr[i])
		rk[i] = i
	}
	//读取数据并建立索引数组
	// 1 5 3 4 2
	//   0 1 2 3 4
	//   0 4 2 3 1
	sort.SliceStable(rk, func(i, j int) bool {
		return arr[rk[i]] < arr[rk[j]]
	})
	//排索引数组 按照数组内部元素大小排
	//
	var head, tail Node
	head.next = &tail
	tail.pre = &head
	head.val = arr[rk[0]] - 1e9
	tail.val = arr[rk[length-1]] + 1e9
	for i := 0; i < length; i++ {
		pos[rk[i]] = AddNode(arr, tail.pre, rk[i])
	}
	for j := length - 1; j > 0; j-- {
		curr := pos[j]
		if arr[j]-curr.pre.val <= curr.next.val-arr[j] {
			ans[j] = curr.pre.idx
		} else {
			ans[j] = curr.next.idx
		}
		DeleteNode(curr)
	}
	for k := 1; k < length; k++ {
		temp := arr[ans[k]] - arr[k]
		if temp < 0 {
			temp = -temp
		}
		fmt.Printf("%d %d\n", temp, ans[k]+1)
	}
}
