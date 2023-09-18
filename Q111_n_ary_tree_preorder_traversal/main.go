package Q111

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	traversalTree(root, &res)
	return res
}

func traversalTree(root *Node, arr *[]int) {
	if root == nil {
		return
	}

	*arr = append(*arr, root.Val)

	for _, v := range root.Children {
		traversalTree(v, arr)
	}
}
