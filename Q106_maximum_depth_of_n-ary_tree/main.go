package Q106

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root.Children == nil {
		return 0
	}

	for _, n := range root.Children {
		tmp := maxDepth(n)
	}

	return 0
}
