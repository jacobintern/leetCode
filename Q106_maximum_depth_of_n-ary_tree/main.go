package Q106

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root.Children == nil {
		return 0
	}
	max := 1
	for _, n := range root.Children {
		tmp := maxDepth(n)
		if tmp > max {
			max = tmp
		}
	}

	return max + 1
}
